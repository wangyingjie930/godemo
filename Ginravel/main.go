package main

import (
	"Gin/app/Models"
	"Gin/app/Providers"
	"Gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	provider := Providers.DatabaseServiceProvider{}
	provider.Boot()
	defer Models.Query().Close()

	// 1.创建路由
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	routes.Web(r)

	// 3.监听端口，默认在8080
	r.Run(":8001")
}
