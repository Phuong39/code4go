package main

func init() {
	Setup()
	CasbinSetup()
}

func main() {
	r := InitRouter()
	r.Run(":9000") //参数为空 默认监听8080端口
}
