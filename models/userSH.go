package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)


type UserSH struct {
	ID  			int
	UserID			int
	Status			string
	CreateDate 		time.Time          `orm:"auto_now_add;type(datetime);null" json:"create_date"`
	TJxiangmu 		string
	XFxiangmu 		string
	XiaoHao			string
}


func init(){
	orm.RegisterModel(new(UserSH))
}