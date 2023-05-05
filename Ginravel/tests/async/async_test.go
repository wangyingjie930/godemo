/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package async

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestGinAsync(t *testing.T) {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.GET("/login_async", func(c *gin.Context) {
		c1 := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行：" + c1.Request.URL.Path)
		}()
	})

	r.Run(":8001")
}
