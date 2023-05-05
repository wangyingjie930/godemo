/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthenticateMiddleware struct {
}

func (m AuthenticateMiddleware) Handle(c *gin.Context) {
	fmt.Println("auth middleware")
}
