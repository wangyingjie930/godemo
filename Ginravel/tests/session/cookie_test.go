/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package session

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestCookie(t *testing.T) {
	r := gin.Default()
	r.GET("/cookie", func(context *gin.Context) {
		cookie, err := context.Cookie("key1")
		if err != nil {
			context.SetCookie("key1", "value1", 60,
				"/", "127.0.0.1", false, true)
		}
		context.String(http.StatusOK, cookie)
	})
	r.Run(":8081")
}

func LoginMiddleware(context *gin.Context) {
	cookie, err := context.Cookie("user_id")
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "StatusUnauthorized"})
		context.Abort()
	}
	context.Set("user_id", cookie)
}

func TestLoginMiddleware(t *testing.T) {
	r := gin.Default()
	r.GET("/home", LoginMiddleware, func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": context.Request.URL})
	})

	r.GET("login", func(context *gin.Context) {
		_, err := context.Cookie("user_id")
		if err != nil {
			context.SetCookie("user_id", "1", 60, "/", "127.0.0.1", false, true)
			context.String(http.StatusOK, "登陆成功")
		}
		context.String(http.StatusOK, "已登陆")
	})

	r.Run(":8081")
}
