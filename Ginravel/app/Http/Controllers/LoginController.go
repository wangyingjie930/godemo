/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package Controllers

import (
	"Gin/app/Http/Components/Response"
	"github.com/gin-gonic/gin"
)

type LoginStruct struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

type LoginController struct {
}

func (l LoginController) LoginJson(ctx *gin.Context) {
	// 声明接收的变量
	var json LoginStruct
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := ctx.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		Response.Error(ctx, err)
		return
	}

	if json.User == "root" && json.Password == "root" {
		Response.Success(ctx, "success")
		return
	}
	Response.Error(ctx, Response.PasswordError)
}

func (l LoginController) Login(ctx *gin.Context) {
	Response.Success(ctx, "login")
}

func (l LoginController) Submit(ctx *gin.Context) {
	Response.Success(ctx, "submit")
}
