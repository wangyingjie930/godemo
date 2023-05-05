/**
  @author: wangyingjie
  @since: 2023/4/29
  @desc:
**/

package Providers

import (
	_ "Beego/models"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		"homestead:secret@tcp(127.0.0.1:33060)/testdb?charset=utf8mb4&parseTime=True&loc=Local")
	orm.RunCommand()
}
