/**
  @author: wangyingjie
  @since: 2023/4/30
  @desc:
**/

package utils

import (
	Response "Beego/utils/response"
	"github.com/beego/beego/v2/adapter/validation"
	Context "github.com/beego/beego/v2/server/web/context"
)

func FormValidate(c *Context.Context, data interface{}) {
	validator := new(validation.Validation)
	b, err := validator.Valid(data)
	if err != nil {
		Response.Error(c, err)
	}
	if !b {
		c.Output.JSON(Response.NewJson(4000, "字段验证失败", validator.Errors), false, false)
	}
}
