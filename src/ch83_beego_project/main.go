package main

import (
	_ "ch83_beego_project/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/*
open in terminal
bee run
 */
func main() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	beego.SetStaticPath("/down1", "download1")
	beego.SetStaticPath("/down2", "download2")

	fmt.Println(beego.AppConfig.String("mysqlpass"))
	beego.Run()
}

