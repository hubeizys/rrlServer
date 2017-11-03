package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SomeNews struct {
	NewsID      int64  `orm:"column(Id);pk;auto"`
	Title       string `orm:"size(64)" json:"title"`
	NewsINFO    string `orm:"size(512)" json:"news_info"`
	NewsUper    string  `json:"news_uper"`
	NewsRemaker string `orm:"size(128)" json:"news_remaker"`
	CreateDate 	time.Time          `orm:"auto_now_add;type(datetime);null" json:"create_date"`
}

func init() {
	orm.RegisterModel(new(SomeNews))
}

func AddNews(nns SomeNews) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(&nns)
}

func GetNews() ([]SomeNews, int64, error) {
	o := orm.NewOrm()
	var _l_somenews_list []SomeNews
	num, err:=o.QueryTable(SomeNews{}).OrderBy("Id").Limit(10, 0).All(&_l_somenews_list)
	return _l_somenews_list, num, err
}