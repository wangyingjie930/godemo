package routers

import (
	"Beego/controllers"
	Response "Beego/utils/response"
	"errors"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/test/?:id", &controllers.TestController{})

	beego.Router("/test/custom", &controllers.TestController{},
		"get:CustomGet;post:CustomPost")

	beego.Include(&controllers.CMSController{})

	// test namespace
	ns := beego.NewNamespace("v1", beego.NSRouter("/test/custom", &controllers.TestController{},
		"get:CustomGet;post:CustomPost"))
	beego.AddNamespace(ns)

	//test stopRun
	beego.Router("/test/stopRun", &controllers.TestController{}, "get:TestStopRun")

	//orm
	beego.Router("/orm/find", &controllers.ModelController{}, "get:FindUser")
	beego.Router("/orm/insert", &controllers.ModelController{}, "get:BatchInsert")
	beego.Router("/orm/transaction", &controllers.ModelController{}, "get:TransAction")

	//validate
	beego.Router("/validate", &controllers.ModelController{}, "get:Validate")

	//middleware
	testMiddleware()

	//login
	beego.Router("/login", &controllers.MainController{}, "get:Login")
}

func testMiddleware() {
	beego.InsertFilterChain("/*", func(next beego.FilterFunc) beego.FilterFunc {
		return func(ctx *context.Context) {
			// do something
			fmt.Println("hello")
			// don't forget this
			next(ctx)
		}
	})

	var FilterUser = func(ctx *context.Context) {
		userId := ctx.Input.Session("user_id")
		uri := ctx.Input.URL()
		if uri == "/login" {
			return
		}

		if userId == nil {
			Response.Error(ctx, errors.New("未进行登录"))
		}
		ctx.Input.SetData("user_id", userId)
	}
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

}
