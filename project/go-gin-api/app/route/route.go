package route

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/controller"
	"go-gin-api/app/route/middleware/exception"
	"go-gin-api/app/route/middleware/jaeger"
	"go-gin-api/app/route/middleware/logger"
	"go-gin-api/app/util/response"
)

func SetupRouter(engine *gin.Engine) {
	//设置路由中间件
	engine.Use(logger.SetUp(), exception.SetUp(), jaeger.SetUp())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1, "pong", nil)
	})

	// 测试链路追踪
	engine.GET("/jaeger_test", controller.JaegerTest)
}
