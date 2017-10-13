package models

import (
	"time"
	"github.com/astaxie/beego/orm"

)

type GoodsRecord struct {
	GoodsRecordID 	int64 	`orm:"column(id);pk;auto" json:"goods_record_id"`

	Name 			string 	`orm:"size(50);column(name)" json:"name"`
	CreateDate 		time.Time          `orm:"auto_now_add;type(datetime);null" json:"create_date"`         //创建时间
	Op				string  				`orm:"size(50)" json:"op"`
	GoodsBus		string			`orm:"size(50)" json:"goods_bus"`
	OpUser			int64			`orm:"null" json:"op_user"`
	OpGoodsID		int 			 `orm:";null" json:"op_goods_id"`
	GoodsType 		string 				`orm:"size(50)" json:"goods_type"`
	GoodsPrice 		int64 				`orm:"" json:"goods_price"`
	OpNum			int64			`orm:"column(opNum)" json:"op_num"`
}

func init(){
	orm.RegisterModel(new(GoodsRecord))
}