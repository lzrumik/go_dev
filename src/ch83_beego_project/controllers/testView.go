package controllers

import (
	"github.com/astaxie/beego"
	"ch83_beego_project/models"
)

type TestViewController struct {
	beego.Controller
}

func (c *TestViewController) Get() {
	var users []models.UserInfo
	models.ReadUserInfo(&users)


	c.Data["Users"] = users
	c.Data["len"]   = len(users)
	c.TplName = "test_view.tpl"
}
