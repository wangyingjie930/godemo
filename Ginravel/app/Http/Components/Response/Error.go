/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package Response

var (
	ValidateError = NewError(4000, "字段验证失败")
	PasswordError = NewError(4001, "密码错误")
)

type ErrorData struct {
	code    int
	message string
}

func NewError(code int, message string) *ErrorData {
	return &ErrorData{code: code, message: message}
}

func (e ErrorData) Error() string {
	return e.message
}
