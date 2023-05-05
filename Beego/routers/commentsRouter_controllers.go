package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["Beego/controllers:CMSController"] = append(beego.GlobalControllerRouter["Beego/controllers:CMSController"],
		beego.ControllerComments{
			Method:           "AllBlock",
			Router:           "/all/:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Beego/controllers:CMSController"] = append(beego.GlobalControllerRouter["Beego/controllers:CMSController"],
		beego.ControllerComments{
			Method:           "StaticBlock",
			Router:           "/staticblock/:key",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
