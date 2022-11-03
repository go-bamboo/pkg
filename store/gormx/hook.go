package gormx

import (
	"context"
	"fmt"

	"github.com/go-bamboo/pkg/log"
	"github.com/go-bamboo/pkg/tracing"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

const (
	gormSpanKey        = "__gorm_span"
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"
	gormPluginName     = "opentracingPlugin"
)

func setDbSpan(ctx context.Context, span trace.Span, m *gorm.Statement) {
	attrs := []attribute.KeyValue{
		attribute.String("db.system", "gorm"),
	}
	if len(m.Selects) > 0 {
		for i, s := range m.Selects {
			attrs = append(attrs, attribute.String(fmt.Sprintf("db.select.%d", i), s))
		}
	}
	span.SetAttributes(attrs...)
}

func before(db *gorm.DB) {
	p := db.Plugins[gormPluginName].(*GormTracingHook)
	tracer := p.tracer

	//operation := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	operation := "gorm: " + db.Statement.Table

	ctx, span := tracer.Start(db.Statement.Context, operation, &DBTextMapCarrier{db: db})
	setDbSpan(ctx, span, db.Statement)
	//db = db.WithContext(ctx)
	db.InstanceSet(gormSpanKey, span)
	return
}

func after(db *gorm.DB) {
	p := db.Plugins[gormPluginName].(*GormTracingHook)
	tracer := p.tracer
	v, ok := db.InstanceGet(gormSpanKey)
	if ok {
		span := v.(trace.Span)
		tracer.End(db.Statement.Context, span, nil, nil)
	}
	if db.Error != nil {
		log.Error(db.Error)
		db.Error = WrapGormError(db.Error)
	}
	return
}

type GormTracingHook struct {
	tracer *tracing.Tracer
}

func NewGormTracingHook(opts ...tracing.Option) gorm.Plugin {
	tracer := tracing.NewTracer(trace.SpanKindClient, tracing.WithPropagator(
		propagation.NewCompositeTextMapPropagator(tracing.Metadata{}, propagation.Baggage{}, tracing.TraceContext{}),
	))
	return &GormTracingHook{
		tracer: tracer,
	}
}

func (p *GormTracingHook) Name() string {
	return gormPluginName
}

func (p *GormTracingHook) Initialize(db *gorm.DB) error {

	// 开始前
	db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// 结束后
	db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return nil
}

var _ gorm.Plugin = &GormTracingHook{}
