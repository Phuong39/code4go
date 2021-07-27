package common

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/pborman/uuid"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
)

const (
	RequestId          = "Wps-Docer-Request-Id"
	SpanContext        = "spanContext"
	LocalAgentHostPort = "localhost:6831"
)

func InitJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "192.168.140.128:6831",
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("Error: connot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

func StartSpanWithContext(ctx context.Context, operationName string, opts ...opentracing.StartSpanOption) (opentracing.Span, context.Context) {
	span := opentracing.GlobalTracer().StartSpan(operationName, opts...)
	spanContext := opentracing.ContextWithSpan(ctx, span)
	return span, spanContext
}

func StartSpanWithHttpContext(ctx context.Context, operation string, opts ...opentracing.StartSpanOption) (opentracing.Span, error) {
	defaultTracer := opentracing.GlobalTracer()
	var span opentracing.Span
	spanContext, ok := ctx.Value(SpanContext).(opentracing.SpanContext)
	if ok {
		opts = append(opts, opentracing.ChildOf(spanContext))
		span = defaultTracer.StartSpan(operation, opts...)
	}
	if span == nil {
		span = defaultTracer.StartSpan(operation, opts...)
	}
	return span, nil
}

func TraceGin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqId string
		defaultTracer := opentracing.GlobalTracer()
		if defaultTracer != nil {
			var span opentracing.Span
			spanContext, err := defaultTracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
			if err != nil {
				span = defaultTracer.StartSpan(ctx.FullPath())
			} else {
				span = defaultTracer.StartSpan(ctx.FullPath(), opentracing.ChildOf(spanContext))
			}
			sp, ok := span.Context().(jaeger.SpanContext)
			if ok && reqId == "" {
				reqId = sp.TraceID().String()
			}
			defer span.Finish()
			ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), SpanContext, span.Context()))
			ctx.Set(SpanContext, span.Context())
			if ctx.Request.Header.Get(RequestId) == "" {
				if reqId == "" {
					reqId = uuid.New()
				}
				ctx.Request.Header.Set(RequestId, reqId)
			} else {
				reqId = ctx.Request.Header.Get(RequestId)
			}
			span.SetTag(RequestId, reqId)
		}

		ctx.Next()
	}
}
