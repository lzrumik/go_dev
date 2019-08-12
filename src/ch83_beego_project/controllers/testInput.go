package controllers

type TestInputController struct {
	BaseController
}

type User struct{
	Username string
	Password string
}

func (c *TestInputController) Get(){
	//id := c.GetString("id")
	//c.Ctx.WriteString("<html>" + id + "<br/>")

	//name := c.Input().Get("name")
	//c.Ctx.WriteString(name + "</html>")

	c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_input" method="post"> 
							<input type="text" name="Username"/>
							<input type="password" name="Password"/>
							<input type="submit" value="提交"/>
					   </form></html>`)
}


func (c *TestInputController) Post(){
	u := User{}
	if err:=c.ParseForm(&u) ; err != nil{
		//process error
	}

	c.Ctx.WriteString("Username:" + u.Username + " Password:" + u.Password)
}
