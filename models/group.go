package models

import (
	"github.com/astaxie/beego/orm"
)

type Group struct {
	GroupID       int64 `orm:"pk;auto" json:"group_id"`
	GroupName     string
	GroupType     int64
	GroupMasterID string
	Remark        string
}

func init() {
	orm.RegisterModel(new(Group))
}

func AddGroup(group Group)(int64 ,error)  {
	o:= orm.NewOrm()
	return o.Insert(&group)
}

func GetGroupByUserID(user_id string)(Group ,error)  {
	o:= orm.NewOrm()
	var group  Group
	err:= o.QueryTable(Group{}).Filter("GroupMasterID", user_id).One(&group)
	return group , err
}