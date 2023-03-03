package gormx

import (
	"fmt"

	otelext "github.com/go-bamboo/pkg/otel"
	"go.opentelemetry.io/otel"
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

type GormTracer struct {
	tracer     trace.Tracer
	propagator propagation.TextMapPropagator
}

func NewGormTracer() gorm.Plugin {
	return &GormTracer{
		tracer:     otel.Tracer("gorm"),
		propagator: propagation.NewCompositeTextMapPropagator(otelext.Metadata{}, propagation.Baggage{}, otelext.TraceContext{}),
	}
}

func (p *GormTracer) Name() string {
	return gormPluginName
}

func (p *GormTracer) Initialize(db *gorm.DB) error {

	// 开始前
	db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, p.before)
	db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, p.before)
	db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, p.before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, p.before)
	db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, p.before)
	db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, p.before)

	// 结束后
	db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, p.after)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, p.after)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, p.after)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, p.after)
	db.Callback().Row().After("gorm:row").Register(callBackAfterName, p.after)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, p.after)
	return nil
}

func (p *GormTracer) before(db *gorm.DB) {
	tracer := p.tracer
	//operation := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	operation := "gorm: " + db.Statement.Table
	_, span := tracer.Start(db.Statement.Context, operation)
	attrs := []attribute.KeyValue{
		attribute.String("db.system", "gorm"),
	}
	if len(db.Statement.Selects) > 0 {
		for i, s := range db.Statement.Selects {
			attrs = append(attrs, attribute.String(fmt.Sprintf("db.select.%d", i), s))
		}
	}
	span.SetAttributes(attrs...)
	//db = db.WithContext(ctx)
	db.InstanceSet(gormSpanKey, span)
	return
}

func (p *GormTracer) after(db *gorm.DB) {
	v, ok := db.InstanceGet(gormSpanKey)
	if !ok {
		return
	}
	span, ok := v.(trace.Span)
	if !ok {
		return
	}
	if db.Error != nil {
		span.RecordError(db.Error)
	}
	span.End()
	return
}

var _ gorm.Plugin = &GormTracer{}
