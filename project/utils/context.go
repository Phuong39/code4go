package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"log"
)

const (
	TracerKey            = "Tracer"
	ParentSpanContextKey = "ParentSpanContext"
)

func StartSpanWithHttp(ctx *gin.Context, operation string, opts ...opentracing.StartSpanOption) (opentracing.Span, map[string][]string, error) {
	header := make(map[string][]string)
	tracer, _ := ctx.Get(TracerKey)
	parentSpanContext, _ := ctx.Get(ParentSpanContextKey)
	opts = append(opts,
		opentracing.ChildOf(parentSpanContext.(opentracing.SpanContext)),
		opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
		ext.SpanKindRPCClient)
	span := opentracing.StartSpan(operation, opts...)
	err := tracer.(opentracing.Tracer).Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(header))
	if err != nil {
		log.Fatalf("%s: Couldn't inject headers", err)
	}
	return span, header, err
}

func GetSpanWithContext(ctx *gin.Context, operation string, opts ...opentracing.StartSpanOption) (opentracing.Span, error) {
	tracer := opentracing.GlobalTracer()
	if tracer == nil {
		return nil, errors.New("GlobalTracer is nil")
	}
	spCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
	if err != nil {
		return tracer.StartSpan(operation, opts...), nil
	} else {
		opts = append(opts, opentracing.ChildOf(spCtx), opentracing.Tag{Key: string(ext.Component), Value: "HTTP"}, ext.SpanKindRPCServer)
		return opentracing.StartSpan(operation, opts...), nil
	}
}
