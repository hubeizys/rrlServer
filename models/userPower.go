package models

import (
	"github.com/astaxie/beego/orm"
)


type UserPower struct {
	UserID int64 `orm:"pk"`
	PowerLev int
	PowerInfo string `orm:"size(2048)" json:"power_info"`
	Remark	string `orm:"size(64)" json:"remark"`
}


func init()  {
	orm.RegisterModel(new(UserPower))
}


func AddUserPower(upower UserPower) (int64, error) {
	o := orm.NewOrm()
	var l_upower UserPower
	l_upower.PowerInfo = upower.PowerInfo
	l_upower.PowerLev = upower.PowerLev
	l_upower.Remark = upower.Remark
	l_upower.UserID = upower.UserID
	return o.Insert(&upower)
}


func GetPower(user_id int64) (UserPower, error) {
	o := orm.NewOrm()
	var _l_userpower UserPower
	err := o.QueryTable(UserPower{UserID:user_id}).One(&_l_userpower)
	return _l_userpower, err
}

