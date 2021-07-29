package main

import (
	"github.com/gin-gonic/gin"
	"utils"
)

func main() {
	utils.NewTracer("producer", "192.168.140.128")
	//_, closer, err := utils.NewJaegerTracer("producer", utils.JaegerHostPort)
	//if err != nil {
	//	panic(err)
	//}
	//defer closer.Close()

	g := gin.Default()
	// g.Use(utils.TraceSetUp())
	g.GET("/hello", Hello)
	g.GET("/say", Say)
	err := g.Run(":9000")
	if err != nil {
		return
	}
}
