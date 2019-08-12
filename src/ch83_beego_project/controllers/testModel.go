package controllers

import (
	"ch83_beego_project/models"
	"github.com/astaxie/beego"
)

type TestModelController struct {
	beego.Controller
}

func (c *TestModelController) Get() {
	user := models.UserInfo{Username:"liusi", Password:"7654321"}
	models.AddUser(&user)

	c.Ctx.WriteString("call model success!")
}