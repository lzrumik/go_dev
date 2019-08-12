package routers

import (
	"ch83_beego_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserController{})
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/test_login", &controllers.TestLoginController{}, "get:Login;post:Post")
	beego.Router("/test_input", &controllers.TestInputController{}, "get:Get;post:Post")
	beego.Router("/test_model", &controllers.TestModelController{}, "get:Get;post:Post")
	beego.Router("/test_view", &controllers.TestViewController{}, "get:Get;post:Post")
	beego.Router("/test_httplib", &controllers.TestHttpLibController{}, "get:Get;post:Post")
	beego.Router("/test_context", &controllers.TestContextController{}, "get:Get;post:Post")
}
