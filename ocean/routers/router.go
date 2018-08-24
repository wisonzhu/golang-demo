package routers

import (
	"awesomeProject/ocean/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test/",&controllers.TestController{})
    beego.Router("/", &controllers.MainController{})
	//beego.SetStaticPath("/static/","/static/")

}
