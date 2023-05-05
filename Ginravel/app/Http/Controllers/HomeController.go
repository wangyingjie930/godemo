/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package Controllers

import (
	"Gin/app/Http/Components/Response"
	"Gin/app/Models"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func (h HomeController) Index(context *gin.Context) {
	Response.Success(context, "success")
}

func (h HomeController) Search(context *gin.Context) {
	userId, exist := context.GetQuery("user_id")
	if !exist {
		Response.Error(context, Response.ValidateError)
		return
	}

	var user Models.User
	Models.Query().First(&user, userId)
	Response.Success(context, user)
}

func (h HomeController) Insert(context *gin.Context) {

}
