/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package request

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"testing"
)

type People struct {
	Username string `form:"username" json:"username" binding:"required,startswith=hi" validate:"checkName"`
	Password string `form:"password" json:"password" binding:"required"`
}

func TestStruct(t *testing.T) {
	r := gin.Default()
	r.GET("validate", func(context *gin.Context) {
		var people People
		err := context.ShouldBind(&people)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, people)
	})

	r.Run(":8081")
}

// 自定义验证函数
func checkName(fl validator.FieldLevel) bool {
	if fl.Field().String() != "root" {
		return false
	}
	return true
}

func TestCheckNameBind(t *testing.T) {
	r := gin.Default()
	validate := validator.New()
	r.GET("/validate", func(context *gin.Context) {
		var people People
		if err := context.ShouldBind(&people); err != nil {
			context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		validate.RegisterValidation("checkName", checkName)
		if err := validate.Struct(people); err != nil {
			context.JSON(http.StatusOK, gin.H{"errors": err.Error()})
			return
		}
	})

	r.Run(":8081")
}
