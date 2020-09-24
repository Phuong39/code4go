package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//拦截器
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var e *casbin.Enforcer
		e := Enforcer
		//从DB加载策略
		e.LoadPolicy()

		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := "admin"
		//判断策略中是否存在
		if ok, _ := e.Enforce(sub, obj, act); ok {
			fmt.Println("权限验证通过")
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"message": "权限验证失败",
			})
			c.Abort()
		}
	}
}

func Hello(c *gin.Context) {
	fmt.Println("Hello 接收到GET请求..")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": "Hello 接收到GET请求..",
	})
}

func InitRouter() *gin.Engine {

	//获取router路由对象
	r := gin.New()

	r.GET("/add", Add)
	r.GET("/get", Get)
	r.GET("/delete", Delete)

	apiv1 := r.Group("/api/v1")
	//使用自定义拦截器中间件
	apiv1.Use(Authorize())
	{
		//创建请求
		apiv1.GET("/hello", Hello)

		userController := new(UserInfo)
		apiv1.GET("/users", userController.GetUsers)
		apiv1.GET("/users/:id", userController.GetUser)
		apiv1.POST("/users", userController.Add)
		apiv1.PATCH("/users/:id", userController.Edit)
		apiv1.DELETE("/users/:id", userController.Delete)
	}

	return r
}
