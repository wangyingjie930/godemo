/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package Response

import (
	"github.com/gin-gonic/gin"
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

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewJson(http.StatusOK, "success", data))
}

func Error(c *gin.Context, err error) {
	errData, e := err.(*ErrorData)
	if e {
		c.JSON(http.StatusOK, NewJson(errData.code, err.Error(), nil))
	} else {
		c.JSON(http.StatusOK, NewJson(http.StatusInternalServerError, err.Error(), nil))
	}
	c.Abort()
}
