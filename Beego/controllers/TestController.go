/**
  @author: wangyingjie
  @since: 2023/4/28
  @desc:
**/

package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type TestController struct {
	beego.Controller
}

func (t *TestController) Get() {
	fmt.Println("hello")
	t.Ctx.WriteString("hello")
}

func (t *TestController) CustomGet() {
	///test/custom?test=1
	fmt.Println("当前的url:", beego.URLFor("TestController.CustomGet", "test", 1))
	t.Ctx.Output.Body([]byte("custom get1111"))
}

func (t *TestController) CustomPost() {
	t.Ctx.Output.Body([]byte("custom post"))
}

func (t *TestController) TestStopRun() {
	t.StopRun()
}
