/**
  @author: wangyingjie
  @since: 2023/4/29
  @desc:
**/

package controllers

import (
	Models "Beego/models"
	"Beego/utils"
	Response "Beego/utils/response"
	"context"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type ModelController struct {
	MainController
}

func (m *ModelController) FindUser() {
	user := new(Models.User)

	//查询
	user.Find(m.Ctx.Input.GetData("user_id").(int))
	fmt.Println("根据id:", user)
	user.FindByName(user.Name)
	fmt.Println("根据name", user)

	//插入
	//newUser := user.CreateRandom()
	//fmt.Println(newUser)

	//更新
	//newUser.Name = "updateName:" + newUser.Number
	//Models.Builder().Update(newUser, "Name")

	o1 := orm.NewOrm()
	o2 := orm.NewOrm()
	fmt.Printf("%+v %+v", o1, o2)

	m.Ctx.Output.JSON(user, false, false)
}

// BatchInsert
//
//	@Description: 批量插入测试
//	@receiver m
func (m *ModelController) BatchInsert() {
	user := &Models.User{}
	ret := user.BatchCreate()
	m.Ctx.Output.JSON(ret, false, false)
}

// TransAction
//
//	@Description: 测试事务
//	@receiver m
func (m *ModelController) TransAction() {
	userData := new(Models.User).GetRandData()
	err := Models.Builder().DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		_, err := txOrm.Insert(&userData)
		return err
	})
	fmt.Println(err)
	Response.Success(m.Ctx, userData)
}

func (m *ModelController) Validate() {
	type userData struct {
		Id     int    `form:"id"`
		Number string `form:"number" valid:"Required"`
	}

	//解析字段
	u := new(userData)
	if err := m.ParseForm(u); err != nil {
		Response.Error(m.Ctx, err)
	}
	//验证字段
	utils.FormValidate(m.Ctx, u)
	Response.Success(m.Ctx, u)
}
