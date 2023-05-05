/**
  @author: wangyingjie
  @since: 2023/4/23
  @desc:
**/

package Models

import "time"

type User struct {
	BaseModel
	Number    string
	Name      string
	Password  string
	Gender    string
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"column:create_time;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:update_time;type:timestamp"`
}

func (u User) TableName() string {
	return "user"
}
