package models

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

type BlogLogin struct {
	Id            int64
	Name          string
	Pwd           string
	WechatId      string
	WechatInfo    string
	CreateTime    string
	LastLoginIp   string
	LastLoginTime string
}


type User struct {
	Id          int
	Name        string
	Profile     *Profile   `orm:"rel(one)"`
	Post        []*Post `orm:"reverse(many)"`
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"`
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}


func RegisterDB() {
	orm.RegisterModel(new(BlogLogin))
	orm.RegisterDataBase("default", "mysql", "root:123456789@/test?charset=utf8")
	//orm.RunSyncdb("default", false, true)

}

