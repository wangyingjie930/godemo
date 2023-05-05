package controllers

import (
	Models "Beego/models"
	"Beego/utils"
	Response "Beego/utils/response"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (m *MainController) Login() {
	type LoginData struct {
		Number   string `form:"number" valid:"Required"`
		Password string `form:"password" valid:"Required"`
	}

	loginData := new(LoginData)
	if err := m.ParseForm(loginData); err != nil {
		Response.Error(m.Ctx, Response.ValidateError)
	}

	//验证字段
	utils.FormValidate(m.Ctx, loginData)

	user := &Models.User{Number: loginData.Number}
	if err := orm.NewOrm().Read(user, "Number"); err != nil {
		Response.Error(m.Ctx, err)
	}

	if user.Password != loginData.Password {
		Response.Error(m.Ctx, Response.PasswordError)
	}

	m.SetSession("user_id", int(user.ID))
	Response.Success(m.Ctx, "登陆成功")
}
