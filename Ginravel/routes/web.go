/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package routes

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Handle(e *gin.Engine)
}

func Web(e *gin.Engine) {
	mapper := []Router{
		HomeRouter{},
		LoginRouter{},
	}

	for _, router := range mapper {
		router.Handle(e)
	}
}
