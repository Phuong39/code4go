package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"utils"
)

func Hello(ctx *gin.Context) {
	resp1, err := utils.HttpGet(ctx, "http://127.0.0.1:9000/hello", "HTTP -> Hello")
	if err != nil {
		panic(err)
	}
	resp2, err := utils.HttpGet(ctx, "http://127.0.0.1:9000/say", "HTTP -> Say")
	if err != nil {
		panic(err)
	}
	time.Sleep(50 * time.Millisecond)
	ctx.JSON(200, string(resp1)+"-"+string(resp2))
}
