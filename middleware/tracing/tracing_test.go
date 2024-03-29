package tracing

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/propagation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var _ transport.Transporter = &mockTransport{}

type headerCarrier http.Header

// Get returns the value associated with the passed key.
func (hc headerCarrier) Get(key string) string {
	return http.Header(hc).Get(key)
}

// Set stores the key-value pair.
func (hc headerCarrier) Set(key string, value string) {
	http.Header(hc).Set(key, value)
}

func (hc headerCarrier) Add(key string, value string) { http.Header(hc).Add(key, value) }

// Keys lists the keys stored in this carrier.
func (hc headerCarrier) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range http.Header(hc) {
		keys = append(keys, k)
	}
	return keys
}

func (hc headerCarrier) Values(key string) []string {
	return http.Header(hc).Values(key)
}

type mockTransport struct {
	kind      transport.Kind
	endpoint  string
	operation string
	header    headerCarrier
}

func (tr *mockTransport) Kind() transport.Kind            { return tr.kind }
func (tr *mockTransport) Endpoint() string                { return tr.endpoint }
func (tr *mockTransport) Operation() string               { return tr.operation }
func (tr *mockTransport) RequestHeader() transport.Header { return tr.header }
func (tr *mockTransport) ReplyHeader() transport.Header   { return tr.header }

func TestTracer(t *testing.T) {
	carrier := headerCarrier{}
	tp := tracesdk.NewTracerProvider(tracesdk.WithSampler(tracesdk.TraceIDRatioBased(0)))

	// caller use Inject
	cliTracer := NewTracer(
		trace.SpanKindClient,
		WithTracerProvider(tp),
		WithPropagator(
			propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}),
		),
	)

	ts := &mockTransport{kind: transport.KindHTTP, header: carrier}

	ctx, aboveSpan := cliTracer.Start(transport.NewClientContext(context.Background(), ts), ts.Operation(), ts.RequestHeader())
	defer cliTracer.End(ctx, aboveSpan, nil, nil)

	// server use Extract fetch traceInfo from carrier
	svrTracer := NewTracer(trace.SpanKindServer, WithPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})))
	ts = &mockTransport{kind: transport.KindHTTP, header: carrier}

	ctx, span := svrTracer.Start(transport.NewServerContext(ctx, ts), ts.Operation(), ts.RequestHeader())
	defer svrTracer.End(ctx, span, nil, nil)

	if aboveSpan.SpanContext().TraceID() != span.SpanContext().TraceID() {
		t.Fatalf("TraceID failed to deliver")
	}

	if v, ok := transport.FromClientContext(ctx); !ok || len(v.RequestHeader().Keys()) == 0 {
		t.Fatalf("traceHeader failed to deliver")
	}
}

func TestServer(t *testing.T) {
	tr := &mockTransport{
		kind:      transport.KindHTTP,
		endpoint:  "server:2233",
		operation: "/test.server/hello",
		header:    headerCarrier{},
	}

	tracer := NewTracer(
		trace.SpanKindClient,
		WithTracerProvider(tracesdk.NewTracerProvider()),
	)

	logger := log.NewStdLogger(os.Stdout)
	logger = log.With(logger, "span_id", SpanID())
	logger = log.With(logger, "trace_id", TraceID())

	var (
		childSpanID  string
		childTraceID string
	)
	next := func(ctx context.Context, req interface{}) (interface{}, error) {
		_ = log.WithContext(ctx, logger).Log(log.LevelInfo,
			"kind", "server",
		)
		childSpanID = SpanID()(ctx).(string)
		childTraceID = TraceID()(ctx).(string)
		return req.(string) + "https://go-kratos.dev", nil
	}

	var ctx context.Context
	ctx, span := tracer.Start(
		transport.NewServerContext(context.Background(), tr),
		tr.Operation(),
		tr.RequestHeader(),
	)

	_, err := Server(
		WithTracerProvider(tracesdk.NewTracerProvider()),
		WithPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})),
	)(next)(ctx, "test server: ")

	span.End()
	assert.NoError(t, err)
	assert.NotEmpty(t, childSpanID)
	assert.NotEqual(t, span.SpanContext().SpanID().String(), childSpanID)
	assert.Equal(t, span.SpanContext().TraceID().String(), childTraceID)

	_, err = Server(
		WithTracerProvider(tracesdk.NewTracerProvider()),
		WithPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})),
	)(next)(context.Background(), "test server: ")

	assert.NoError(t, err)
	assert.Empty(t, childSpanID)
	assert.Empty(t, childTraceID)
}

func TestClient(t *testing.T) {
	tr := &mockTransport{
		kind:      transport.KindHTTP,
		endpoint:  "server:2233",
		operation: "/test.server/hello",
		header:    headerCarrier{},
	}

	tracer := NewTracer(
		trace.SpanKindClient,
		WithTracerProvider(tracesdk.NewTracerProvider()),
	)

	logger := log.NewStdLogger(os.Stdout)
	logger = log.With(logger, "span_id", SpanID())
	logger = log.With(logger, "trace_id", TraceID())

	var (
		childSpanID  string
		childTraceID string
	)
	next := func(ctx context.Context, req interface{}) (interface{}, error) {
		_ = log.WithContext(ctx, logger).Log(log.LevelInfo,
			"kind", "client",
		)
		childSpanID = SpanID()(ctx).(string)
		childTraceID = TraceID()(ctx).(string)
		return req.(string) + "https://go-kratos.dev", nil
	}

	var ctx context.Context
	ctx, span := tracer.Start(
		transport.NewClientContext(context.Background(), tr),
		tr.Operation(),
		tr.RequestHeader(),
	)

	_, err := Client(
		WithTracerProvider(tracesdk.NewTracerProvider()),
		WithPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})),
	)(next)(ctx, "test client: ")

	span.End()
	assert.NoError(t, err)
	assert.NotEmpty(t, childSpanID)
	assert.NotEqual(t, span.SpanContext().SpanID().String(), childSpanID)
	assert.Equal(t, span.SpanContext().TraceID().String(), childTraceID)
}
