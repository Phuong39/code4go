package main

import (
	"github.com/gin-gonic/gin"
	"jaegers/common"
)

func main() {
	_, closer := common.InitJaeger("hello-consumer")
	defer closer.Close()

	g := gin.Default()
	g.Use(common.TraceGin())
	g.GET("/hello", Hello)
	err := g.Run(":8080")
	if err != nil {
		return
	}
}
