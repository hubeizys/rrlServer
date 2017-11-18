package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Goods struct {
	GoodsId int 	`orm:"column(id);pk;auto" json:"goods_id"`
	Name 	string 	`orm:"size(50)" json:"name"`
	Count 	int64	`orm:""  json:"count"`
	CreateDate time.Time          `orm:"auto_now_add;type(datetime);null" json:"create_date"`         //创建时间
	UpdateDate time.Time          `orm:"auto_now;type(datetime);null" json:"update_date"`
	CreateUser int64              `orm:"null" json:"userid"`                        //创建者
	UpdateUser int64              `orm:"null" json:"-"`                        //最后更新者
	Remark 	string  `orm:"size(256)" json:"remark"`
	Yuzhi	int64    `orm:"null"  json:"yuzhi"`
	Danjia   int64
}

func init(){
	orm.RegisterModel(new(Goods))
}
