/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
	"time"
)

func TestGlobal(t *testing.T) {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	})
	{
		r.GET("/middleware", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})
	}

	r.Run(":8081")
}

func Middleware(c *gin.Context) {
	t := time.Now()
	fmt.Println("中间件开始执行了")
	// 设置变量到Context的key中，可以通过Get()取
	c.Set("request", "中间件")
	c.Next()
	status := c.Writer.Status()
	fmt.Println("中间件执行完毕", status)
	t2 := time.Since(t)
	fmt.Println("time:", t2)
}

// TestNextMiddleware
//
//	@Description:
//
// 1. 先执行中间件next之前的
// 2. 执行handle方法
// 3. 执行中间件next之后的
// 输出:
// 中间件开始执行了
// request: 中间件
// 中间件执行完毕
//
//	@param t
func TestNextMiddleware(t *testing.T) {
	r := gin.Default()
	r.Use(Middleware)
	{
		r.GET("/middleware", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})
	}

	r.Run(":8081")
}

func TestPartialMiddleware(t *testing.T) {
	r := gin.Default()
	r.GET("/middleware", Middleware, func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})

	r.Run(":8081")
}
