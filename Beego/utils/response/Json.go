/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package Response

import (
	Context "github.com/beego/beego/v2/server/web/context"
	"net/http"
)

type Json struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewJson(code int, message string, data interface{}) *Json {
	return &Json{Code: code, Message: message, Data: data}
}

func Success(c *Context.Context, data interface{}) {
	c.Output.JSON(NewJson(http.StatusOK, "success", data), false, false)
}

func Error(c *Context.Context, err error) {
	errData, e := err.(*ErrorData)
	if e {
		c.Output.JSON(NewJson(errData.code, err.Error(), nil), false, false)
	} else {
		c.Output.JSON(NewJson(http.StatusInternalServerError, err.Error(), nil), false, false)
	}
}
