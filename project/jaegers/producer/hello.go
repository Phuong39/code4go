package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Hello(ctx *gin.Context) {
	//span, err := utils.GetSpanWithContext(ctx, "producer Hello")
	//if err == nil {
	//	defer span.Finish()
	//}
	time.Sleep(50 * time.Millisecond)
	ctx.JSON(200, "hello!")
}

func Say(ctx *gin.Context) {
	//span, err := utils.GetSpanWithContext(ctx, "producer Say")
	//if err == nil {
	//	defer span.Finish()
	//}
	time.Sleep(50 * time.Millisecond)
	ctx.JSON(200, "say!")
}
