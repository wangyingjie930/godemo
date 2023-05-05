/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package routes

import (
	"Gin/app/Http/Controllers"
	"Gin/app/Http/Middleware"
	"github.com/gin-gonic/gin"
)

type HomeRouter struct {
}

func (r HomeRouter) Handle(e *gin.Engine) {
	controller := Controllers.HomeController{}
	group := e.Group("home")
	{
		group.GET("/index", Middleware.AuthenticateMiddleware{}.Handle, controller.Index)
		group.GET("/search", controller.Search)
	}
}
