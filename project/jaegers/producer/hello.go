package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"jaegers/common"
	"time"
)

func TestDemo(req string, ctx context.Context) (reply string) {
	span, _ := common.StartSpanWithHttpContext(ctx, "span_test_demo")
	defer func() {
		span.Finish()
	}()
	//2. 模拟耗时
	time.Sleep(time.Second / 2)
	//3. 返回reply
	reply = "TestDemoReply"
	return
}

func TestDemo2(req string, ctx context.Context) (reply string) {
	span, _ := common.StartSpanWithHttpContext(ctx, "span_test_demo")
	defer func() {
		span.Finish()
	}()
	time.Sleep(time.Second / 2)
	reply = "TestDemo2Reply"
	TestDemo3(req, ctx)
	return
}

func TestDemo3(req string, ctx context.Context) (reply string) {
	span, _ := common.StartSpanWithHttpContext(ctx, "span_test_demo")
	defer func() {
		span.Finish()
	}()
	time.Sleep(time.Second / 2)
	reply = "TestDemo2Reply"
	return
}

func Hello(ctx *gin.Context) {
	// spanContext, ok := ctx.Value(SpanContext).(opentracing.SpanContext)
	span, err := common.StartSpanWithHttpContext(ctx, "span_root")
	if err != nil {
		defer func() {
			span.Finish()
		}()
	}
	// span, tracerCtx := StartSpanWithContext(ctx.Request.Context(), "span_root", &opentracing.Tag{Key: "request-id", Value: uid})
	r1 := TestDemo("Hello TestDemo", ctx)
	r2 := TestDemo2("Hello TestDemo2", ctx)
	fmt.Println(r1, r2)
	ctx.JSON(200, ctx.Request.Header.Get(common.RequestId))
}
