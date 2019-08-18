package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me1111"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	type u struct {
		Name string
		Age int
		Sex string
	}

	user:= &u{
		Name :"Joe",
		Age:20,
		Sex:"男",
	}
	c.Data["User"] = user

	nums := []int{1,2,3,4,5,6,7,8}
	c.Data["nums"] = nums

}
