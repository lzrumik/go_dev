package main

import (
	_ "ch84_beego_api_project/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/**

https://blog.csdn.net/Libra412/article/details/40401791

bee api hello -conn=root:root@tcp(127.0.0.1:3306)/test

bee run -downdoc=true -gendoc=true


http://localhost:8081/swagger/#!/fortune_user
 */
func main() {
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run(":8081")
}

