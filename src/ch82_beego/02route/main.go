package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RESTfullController struct {
	beego.Controller
}

func (this *RESTfullController) Get(){
	this.Ctx.WriteString("Hello World in GET method")
}

// curl -d "" "http://127.0.0.1:8080/RESTful"
func (this *RESTfullController) Post(){
	this.Ctx.WriteString("Hello World in POST method")
}

type RegExpController struct {
	beego.Controller
}


//http://127.0.0.1:8080/RegExp1/123
func (this *RegExpController) Get(){
	this.Ctx.WriteString(fmt.Sprintln("In RegExp Mode!"))

	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString(fmt.Sprintln("id is %s",id))

	splat := this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString(fmt.Sprintln("splat is %s",splat))

	path := this.Ctx.Input.Param(":path")
	this.Ctx.WriteString(fmt.Sprintln("path is %s",path))

	ext := this.Ctx.Input.Param(":ext")
	this.Ctx.WriteString(fmt.Sprintln("ext is %s",ext))
}


func main(){
	beego.Router("/RESTful",&RESTfullController{})

	//http://127.0.0.1:8080/RegExp1/afdsa23
	beego.Router("/RegExp1/?:id",&RegExpController{})
	//http://127.0.0.1:8080/RegExp2/aaa
	//http://127.0.0.1:8080/RegExp2/123
	beego.Router("/RegExp2/?:id([0-9]+)",&RegExpController{})
	beego.Router("/RegExp3/?:id([\\w]+)",&RegExpController{})
	beego.Router("/RegExp4/abc?:id([1-9]+)de",&RegExpController{})
	beego.Router("/RegExp5/*",&RegExpController{})
	//http://127.0.0.1:8080/RegExp6/123.txt
	beego.Router("/RegExp6/*.*",&RegExpController{})

	beego.Run("127.0.0.1:8080")
}
