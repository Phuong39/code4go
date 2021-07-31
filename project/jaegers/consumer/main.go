package main

import (
	"github.com/gin-gonic/gin"
	"utils"
)

func main() {
	utils.NewTracer("consumer", "192.168.140.128")
	//_, closer, err := utils.NewJaegerTracer("consumer", utils.JaegerHostPort)
	//if err != nil {
	//	panic(err)
	//}
	//defer closer.Close()

	g := gin.Default()
	g.Use(utils.TraceSetUp())
	g.GET("/hello", Hello)
	err := g.Run(":8080")
	if err != nil {
		return
	}
}
