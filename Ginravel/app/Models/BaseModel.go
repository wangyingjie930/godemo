/**
  @author: wangyingjie
  @since: 2023/4/22
  @desc:
**/

package Models

import (
	"Gin/app/Providers"
	"Gin/src/Container"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
}

func Query() *gorm.DB {
	db := Container.App.App("db").(Providers.ConnectionResolverInterface)
	return db.GetConnection()
}
