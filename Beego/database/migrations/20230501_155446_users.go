package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Users_20230501_155446 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20230501_155446{}
	m.Created = "20230501_155446"

	migration.Register("Users_20230501_155446", m)
}

// Run the migrations
func (m *Users_20230501_155446) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable("users", "innodb", "utf8mb4")
	m.PriCol("id").SetAuto(true).SetDataType("int").SetUnsigned(true)
	//####中间部分省略
	m.NewCol("deleted_at").SetDataType("timestamp").SetNullable(true)
	m.NewCol("created_at").SetDataType("timestamp").SetDefault("NOW()")
	m.NewCol("updated_at").SetDataType("timestamp").SetNullable(false).SetDefault("NOW()")
	m.SQL(m.GetSQL())

}

// Reverse the migrations
func (m *Users_20230501_155446) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS users")
}
