/**
  @author: wangyingjie
  @since: 2023/4/23
  @desc:
**/

package Models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	Number    string
	Name      string
	Password  string
	Gender    string
	ID        uint      `orm:"pk;column(id);auto"`
	CreatedAt time.Time `orm:"auto_now_add;column(create_time);type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;column(update_time)"`
}

type Profile struct {
	Id  int
	Age int16
}

func init() {
	orm.RegisterModel(new(User), new(Profile))
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Find(id int) {
	u.ID = uint(id)
	err := Builder().Read(u)
	fmt.Println(err)
}

func (u *User) FindByName(name string) {
	u.Name = name
	err := Builder().Read(u, "Name")
	fmt.Println(err)
}

func (u *User) CreateRandom() *User {
	randUser := u.GetRandData()
	_, err := Builder().Insert(&randUser)
	fmt.Println(err)
	return &randUser
}

func (u *User) BatchCreate() []*User {
	var users []*User
	for i := 0; i < 10; i++ {
		user := u.GetRandData()
		users = append(users, &user)
	}
	Builder().InsertMulti(10, &users)
	return users
}

func (u *User) GetRandData() User {
	randNum := strconv.Itoa(rand.Intn(1000))
	return User{
		Number:    randNum,
		Name:      "test:" + randNum,
		Password:  randNum,
		Gender:    "女",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
