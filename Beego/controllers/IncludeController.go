/**
  @author: wangyingjie
  @since: 2023/4/28
  @desc:
**/

package controllers

import "github.com/beego/beego/v2/server/web"

// CMS API
type CMSController struct {
	web.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (this *CMSController) StaticBlock() {
	this.Ctx.Output.Body([]byte("StaticBlock"))
}

// @router /all/:key [get]
func (this *CMSController) AllBlock() {
	this.Ctx.Output.Body([]byte("AllBlock"))
}
