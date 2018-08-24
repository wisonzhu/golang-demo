package main

import (
	_ "awesomeProject/ocean/routers"
	"github.com/astaxie/beego"
	"awesomeProject/ocean/models"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()

}

func main() {
	orm.Debug = true
	beego.Run()
}

