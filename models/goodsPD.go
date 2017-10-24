package models

import (
	"time"
	"github.com/astaxie/beego/orm"

)


// 用户盘点


type GoodsPD struct {
	GoodsPDID 			int64 			`orm:"pk;auto" json:"goods_pdid"`
	GoodsPDName 		string			`orm:"size(64)" json:"goods_pd_name"`
	GoodsPDResult		string			`orm:"size(128)" json:"goods_pd_result"`
	GoodsPDTime			time.Time		`orm:"auto_now;type(datetime)" json:"goods_pd_time"`
	ID             		int64			`orm:"size(64)" json:"ID"`
}

func init(){
	orm.RegisterModel(new(GoodsPD))
}