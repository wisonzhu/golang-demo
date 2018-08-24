package controllers

import (
	"github.com/astaxie/beego"
	"awesomeProject/ocean/models"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "test"
	c.Data["Email"] = "test@soundai.com"
	c.TplName = "index.tpl"
}


type TestController struct {
	beego.Controller
}


func (c *TestController) Post() {
	//c.Ctx.WriteString("this is post method")

	name := c.Input().Get("name")
	fmt.Println(name)
	pwd := c.Input().Get("pwd")
	o := orm.NewOrm()

	login := models.BlogLogin{Name: name, Pwd: pwd}
	o.Insert(&login)
	err := o.Read(&login, "Name", "Pwd")

	if err != nil {
		fmt.Printf("ERR: %v\n", err)
		c.Redirect("http://www.163.com", 301)
		return
	}
	fmt.Printf("Data: %v\n", login)

	c.Redirect("/static/admin.html", 301)
	return

}


func (This *TestController) Get() {
	Id, _ := This.GetInt(":Id")
	This.Ctx.WriteString("this is get method")
	beego.Debug("Page",Id)



}









