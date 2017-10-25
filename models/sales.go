package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 销售记录

type Sales struct {
	SalesID    int64 `orm:"pk;auto"`
	CardNum    string
	GoodsID    int64
	GoodsName  string
	JiaGE      int
	TiJian     string
	Ewai       string
	Tiyan      string
	CreateDate time.Time          `orm:"auto_now_add;type(datetime);null" json:"create_date"`
}

func init() {
	orm.RegisterModel(new(Sales))
}
