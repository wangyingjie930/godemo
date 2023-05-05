/**
  @author: wangyingjie
  @since: 2023/4/29
  @desc:
**/

package Models

import "github.com/beego/beego/v2/client/orm"

var ormer orm.Ormer

func Builder() orm.Ormer {
	if ormer == nil {
		ormer = orm.NewOrm()
	}
	return ormer
}
