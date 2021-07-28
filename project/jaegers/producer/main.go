package main

import (
	"github.com/gin-gonic/gin"
	"utils"
)

func main() {
	_, closer, err := utils.NewJaegerTracer("consumer", utils.JaegerHostPort)
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	g := gin.Default()
	g.Use(utils.TraceSetUp())
	g.GET("/hello", Hello)
	g.GET("/say", Say)
	err = g.Run(":9000")
	if err != nil {
		return
	}
}
