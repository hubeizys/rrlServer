package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type ReportRecord struct {
	RecordId   int64 	`orm:"pk;auto"`
	UserID     string
	RecordInfo string
	RecordType int
	RecordTime time.Time
}

func init()  {
	orm.RegisterModel(new(ReportRecord))
}

func GetRecordByName(p_ID string, start int, end int) ([]*ReportRecord, error, int64) {
	o := orm.NewOrm()
	var _l_record  []*ReportRecord
	num, err := o.QueryTable(ReportRecord{UserID:p_ID}).Limit(end, start).All(&_l_record)
	return _l_record, err, num
}

func ADDRcord(record ReportRecord)(int64 ,error) {
	o:= orm.NewOrm()
	//var _l_record ReportRecord
	return o.Insert(&record)
}