package controllers

type TestController struct {
	BaseController
}

func (c *TestController) Get() {
	c.Ctx.WriteString("<font style=\"color:red\">welcome test controller!</font>") //相当于php里面的echo 函数
}
