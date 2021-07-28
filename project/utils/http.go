package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(ctx *gin.Context, url, operation string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 5, //默认5秒超时时间
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	span, header, err := StartSpanWithHttp(ctx, operation)
	if err == nil {
		defer span.Finish()
	}
	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return content, err
}

func TraceSetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var parentSpan opentracing.Span
		tracer := opentracing.GlobalTracer()
		if tracer == nil {
			return
		}
		spCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
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
		c.Set(TracerKey, tracer)
		c.Set(ParentSpanContextKey, parentSpan.Context())
		c.Next()
	}
}
