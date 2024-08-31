package plugins

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
	gormSpanKey       = "__gorm_span"
	tracingBeforeName = "opentracing:before"
	tracingAfterName  = "opentracing:after"
	gormPluginName    = "opentracingPlugin"
)

var _ gorm.Plugin = (*GormTracer)(nil)

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
	db.Callback().Create().Before("gorm:tracing_create").Register(tracingBeforeName, p.before)
	db.Callback().Delete().Before("gorm:tracing_delete").Register(tracingBeforeName, p.before)
	db.Callback().Update().Before("gorm:tracing_update").Register(tracingBeforeName, p.before)
	db.Callback().Query().Before("gorm:tracing_query").Register(tracingBeforeName, p.before)
	db.Callback().Row().Before("gorm:tracing_row").Register(tracingBeforeName, p.before)
	db.Callback().Raw().Before("gorm:tracing_raw").Register(tracingBeforeName, p.before)

	// 结束后
	db.Callback().Create().After("gorm:tracing_create").Register(tracingAfterName, p.after)
	db.Callback().Delete().After("gorm:tracing_delete").Register(tracingAfterName, p.after)
	db.Callback().Update().After("gorm:tracing_update").Register(tracingAfterName, p.after)
	db.Callback().Query().After("gorm:tracing_query").Register(tracingAfterName, p.after)
	db.Callback().Row().After("gorm:tracing_row").Register(tracingAfterName, p.after)
	db.Callback().Raw().After("gorm:tracing_raw").Register(tracingAfterName, p.after)
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
