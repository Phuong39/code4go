package main

import (
	"github.com/gin-gonic/gin"
	"mall_api/controller"
	"utils"
)

// https://segmentfault.com/a/1190000016677230
func main() {
	_, closer, err := utils.NewJaegerTracer("mall_api", utils.JaegerHostPort)
	if err != nil {
		panic(err)
	}
	defer closer.Close()

	g := gin.Default()
	g.Use(utils.TraceSetUp())
	// g.GET("/user/myOrder/:id", controller.MyOrder)
	g.GET("/user/info/:id", controller.Info)
	err = g.Run(":8080")
	if err != nil {
		return
	}
}
