package gormx

import (
	"context"
	"time"

	"edu/pkg/log"
	"edu/pkg/tracing"

	"go.opentelemetry.io/otel/attribute"
)

//const gormPluginName = "opentracingPlugin"

type GormTraceRecorder struct {
	log.ZapLogger
	tracer *tracing.Tracer
}

func (l *GormTraceRecorder) New() *GormTraceRecorder {
	return l
}

func (l *GormTraceRecorder) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	ctx, span := l.tracer.Start(ctx, "gorm:trace", &DBTextMapCarrier{})
	sql, rowsAffected := fc()

	span.RecordError(err)
	attrs := []attribute.KeyValue{}
	attrs = append(attrs, attribute.String("sql", sql))
	attrs = append(attrs, attribute.Int64("rowsAffected", rowsAffected))
	span.SetAttributes(attrs...)

	l.tracer.End(ctx, span, nil, nil)
}
