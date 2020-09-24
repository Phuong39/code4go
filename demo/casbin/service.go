package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/casbin/casbin"
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

var Enforcer *casbin.Enforcer

// 初始化casbin
func CasbinSetup() {
	a, err := xormadapter.NewAdapter("mysql", "root:fuck123@tcp(119.29.117.244:3306)/rbac_db?charset=utf8", true)
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
		return
	}
	e, err := casbin.NewEnforcer("casbin/rbac_models.conf", a)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		return
	}

	// 从DB懒加载策略
	if err = e.LoadPolicy(); err != nil {
		panic(err)
	}

	Enforcer = e
}

func Get(c *gin.Context) {
	c.JSON(200, Enforcer.GetPolicy())
}

func Delete(c *gin.Context) {
	if ok, _ := Enforcer.RemovePolicy("admin", c.Query("url"), "GET"); !ok {
		c.JSON(200, "Policy不存在")
	} else {
		c.JSON(200, "删除成功")
	}
}

func Add(c *gin.Context) {
	if ok, _ := Enforcer.AddPolicy("admin", c.Query("url"), "GET"); !ok {
		c.JSON(200, "Policy已经存在")
	} else {
		c.JSON(200, "增加成功")
	}
}
