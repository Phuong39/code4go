package utils

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

const (
	UserHost  = ":10010"
	PayHost   = ":10020"
	OrderHost = ":10030"
)

func InjectMd(ctx context.Context, md metadata.MD) error {
	if opentracing.GlobalTracer() == nil || ctx == nil || md == nil {
		return nil
	}
	spanContext, ok := ctx.Value(ParentSpanContextKey).(opentracing.SpanContext)
	if ok {
		carrier := opentracing.HTTPHeadersCarrier{}
		err := ctx.Value(TracerKey).(opentracing.Tracer).Inject(spanContext, opentracing.TextMap, carrier)
		if err != nil {
			return err
		}
		for k, v := range carrier {
			md.Set(k, v...)
		}
	}
	return nil
}

// ServerInterceptor grpc server
func ServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		defaultTracer := opentracing.GlobalTracer()
		if defaultTracer != nil {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				md = metadata.New(nil)
			}
			carrier := opentracing.HTTPHeadersCarrier{}
			for k, v := range md {
				carrier[k] = v
			}
			tracer := opentracing.GlobalTracer()
			spanContext, err := tracer.Extract(opentracing.HTTPHeaders, carrier)
			if err != nil && err != opentracing.ErrSpanContextNotFound {
				grpclog.Errorf("extract from metadata err: %v", err)
			} else {
				span := tracer.StartSpan(
					info.FullMethod,
					opentracing.ChildOf(spanContext),
					ext.SpanKindRPCServer,
				)
				defer span.Finish()
				ctx = opentracing.ContextWithSpan(ctx, span)
			}
		}
		return handler(ctx, req)
	}
}
