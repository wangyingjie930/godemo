/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package routes

import (
	"Gin/app/Http/Controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRouter struct {
}

func (r LoginRouter) Handle(e *gin.Engine) {
	//路由组
	v1 := e.Group("/v1")
	{ //语句块
		v1.GET("/submit", Controllers.LoginController{}.Submit)
		v1.POST("/login", Controllers.LoginController{}.LoginJson)
	}

	extra(e)
}

func extra(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	//url参数
	e.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "枯藤1")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	//表单参数
	e.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
}
