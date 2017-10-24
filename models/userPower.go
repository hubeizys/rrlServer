package models

import "github.com/astaxie/beego/orm"

type UserPower struct {
	UserID int64 `orm:"pk"`
	PowerLev int
	PowerInfo string `orm:"size(2048)" json:"power_info"`
}

func init()  {
	orm.RegisterModel(new(UserPower))
}