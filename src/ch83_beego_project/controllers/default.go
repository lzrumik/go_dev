package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.Data["Name"]  = "zhangsan"
	c.TplName = "index.tpl"
}
