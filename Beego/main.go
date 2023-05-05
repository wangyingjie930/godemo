package main

import (
	_ "Beego/Providers"
	_ "Beego/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
