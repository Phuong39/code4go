package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"jaegers/common"
	"net/http"
)

func Hello(ctx *gin.Context) {
	span, err := common.StartSpanWithHttpContext(ctx, "span_root")
	if err != nil {
		defer func() {
			span.Finish()
		}()
	}
	resp, err := http.Get("http://127.0.0.1:9000/hello")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	ctx.JSON(200, string(bytes))
}
