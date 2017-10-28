package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type UserPower struct {
	UserID    string `orm:"pk"`
	PassWord  string 	//todo  现在没有内加密。明文密码是不被赞许的，现在先偷个懒
	PowerLev  int
	PowerInfo string `orm:"size(2048)" json:"power_info"`
	Remark    string `orm:"size(64)" json:"remark"`
}

func init() {
	orm.RegisterModel(new(UserPower))
}

func AddUserPower(u_power UserPower) (int64, error) {
	o := orm.NewOrm()
	var _l_upower UserPower
	_l_upower.PowerInfo = u_power.PowerInfo
	_l_upower.PowerLev = u_power.PowerLev
	_l_upower.Remark = u_power.Remark
	_l_upower.UserID = u_power.UserID
	return o.Insert(&u_power)
}

func GetPower(user_id string) (UserPower, error) {
	o := orm.NewOrm()
	var _l_userpower UserPower
	err := o.QueryTable(UserPower{UserID: user_id}).One(&_l_userpower)
	return _l_userpower, err
}

func GetAllPower() ([]*UserPower, int64, error) {
	o := orm.NewOrm()
	var _l_power_list []*UserPower
	num, err := o.QueryTable(UserPower{}).Offset(0).All(&_l_power_list)
	//fmt.Printf("ret %s, err %s", num, err)
	logs.Info("ret %d, err %s", num, err)
	logs.Info("ret %s", _l_power_list)
	return _l_power_list, num, err
}

// 非常严肃情况下的使用
func DelPowerById(p_user_id string) (int64, error) {
	o := orm.NewOrm()
	return o.Delete(&UserPower{UserID: p_user_id})
}

func UpdateUserPower(uid string, power UserPower) (string, error) {
	o := orm.NewOrm()
	_l_power := UserPower{UserID: uid}
	if read_err := o.Read(&_l_power); read_err == nil {
		_l_power.Remark = power.Remark
		_l_power.PowerLev = power.PowerLev
		_l_power.PowerInfo = power.PowerInfo
		num, err := o.Update(&_l_power)
		return string(num), err
	} else {
		return uid, read_err
	}
}
