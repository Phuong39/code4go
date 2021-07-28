package jaeger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go-gin-api/app/config"
	"utils"
)

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var parentSpan opentracing.Span

		tracer, closer, err := utils.NewJaegerTracer(config.AppName, utils.JaegerHostPort)
		if err != nil {
			fmt.Printf("new tracer err: %+v\n", err)
		}
		defer closer.Close()

		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(c.Request.URL.Path)
		} else {
			parentSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
		}
		defer parentSpan.Finish()
		c.Set("Tracer", tracer)
		c.Set("ParentSpanContext", parentSpan.Context())
		c.Next()
	}
}
