package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

/*
记录用户 添加 还有权限的变更
*/
type UserOp struct {
	OpRecordId int64        `orm:"column(id);pk;auto" json:"record_id"`
	UserId     int64        `orm:"null" json:"be_op_user"` // 被操作的人
	OpUserId   int64        `orm:"null" json:"op_user"`    // 操作者
	Op         string       `orm:"size(64)" json:"op"`     // 操作项目
	OpDate     time.Time    `orm:"auto_now_add;type(datetime);null" json:"op_date"`
	OpRemark   string       `orm:"size(128)"  json:"op_remark"`
}

func init() {
	orm.RegisterModel(new(UserOp))
}

func AddUserOp(u UserOp) (int64, error) {
	o := orm.NewOrm()
	var userop UserOp
	userop.Op = u.Op
	userop.OpDate = u.OpDate
	userop.OpRemark = u.OpRemark
	userop.OpUserId = u.OpUserId
	userop.UserId = u.UserId
	return o.Insert(&userop)
}

func DelUserOp(OpRecord_Id int64) error {
	o := orm.NewOrm()
	if num, err := o.Delete(&UserOp{OpRecordId: OpRecord_Id}); err != nil {
		logs.Info("ret num %s", num)
		return err
	}
	return nil
}

func GetAllUserOp() ([]UserOp, int64, error) {
	o := orm.NewOrm()
	var _l_user_op_list []UserOp
	num, err := o.QueryTable(UserOp{}).All(&_l_user_op_list)
	return _l_user_op_list, num, err
}
