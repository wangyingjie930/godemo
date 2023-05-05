/**
  @author: wangyingjie
  @since: 2023/4/23
  @desc:
**/

package Providers

import (
	"Gin/src/Container"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

type DatabaseServiceProvider struct {
}

func (d *DatabaseServiceProvider) Boot() {
	Container.App.Single("db", new(DbManager))
}

type ConnectionResolverInterface interface {
	GetConnection() *gorm.DB
}

type DbManager struct {
	Connection *gorm.DB
}

func (d *DbManager) GetConnection() *gorm.DB {
	if d.Connection == nil {
		db, err := gorm.Open("mysql", "homestead:secret@tcp(127.0.0.1:33060)/testdb?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		d.Connection = db
	}
	return d.Connection
}
