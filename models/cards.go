package models

import (
	"github.com/astaxie/beego/orm"

)

type Cards struct {
	Id string `orm:"column(id);size(64);pk" json:"cards_id"`
	Cardtype string `orm:"size(16)" json:"cardtype"`
	Remark string
}

func init()  {
	orm.RegisterModel(new(Cards))
}

func GetCardsInfo(card_id string)(Cards ,error)  {
	o := orm.NewOrm()
	var _card_info Cards
	err:= o.QueryTable(Cards{Id:card_id}).Filter("id", card_id).One(&_card_info)
	return _card_info, err
}

func AddCard(u_card Cards)(int64, error){
	o:= orm.NewOrm()
	return o.Insert(&u_card)
}