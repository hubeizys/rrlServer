package controllers

import (
	"nepliteApi/models"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

type UserPowerController struct {
	beego.Controller
}

// 新建一个power
func (userPower *UserPowerController) Add() {
	result := make(map[string]interface{})
	result["num"] = 0
	result["err"] = ""
	result["result"] = ""

	var _l_userpower models.UserPower
	err := json.Unmarshal(userPower.Ctx.Input.RequestBody, &_l_userpower)
	if err != nil {
		result["err"] = -1
		result["result"] = err
		userPower.Data["json"] = result
		userPower.ServeJSON()
	}

	if id, err := models.AddUserPower(_l_userpower); err != nil {
		result["err"] = -2
		result["result"] = id
		userPower.Data["json"] = result
		userPower.ServeJSON()
	} else {
		result["result"] = id
	}

	userPower.Data["json"] = result
	userPower.ServeJSON()
}


func (userPower *UserPowerController) GetAll(){
	result := make(map[string]interface{})
	result["err"] = 0
	result["num"] = 0

	power_list , num, err  := models.GetAllPower()
	result["result"] = power_list
	result["err"] = err
	result["num"] = num

	userPower.Data["json"] = result
	userPower.ServeJSON()
}

// 获得权限信息
func (userPower *UserPowerController) GetPower() {
	result := make(map[string]interface{})
	result["err"] = 0
	result["num"] = 0

	var l_user_id int64
	if user_id, err := userPower.GetInt64("user_id"); err != nil {
		result["err"] = -1
		result["result"] = err
		userPower.Data["json"] = result
		userPower.ServeJSON()
	} else {
		l_user_id = user_id
		result["result"] = user_id
	}

	if r_userpower, err := models.GetPower(l_user_id); err != nil {
		result["err"] = -2
		result["result"] = err
	} else {
		result["result"] = r_userpower
	}
	userPower.Data["json"] = result
	userPower.ServeJSON()
}

// 推送权限信息  // 测试用的已经放弃了
func (userPower *UserPowerController) PutPower() {
	result := make(map[string]interface{})
	uid, err := userPower.GetInt64(":uid")

	logs.Warn("uid : %d   err : %d", uid, err)
	userPower.Data["json"] = result
	userPower.ServeJSON()
}

func (userPower *UserPowerController) UpdatePower() {
	result := make(map[string]interface{})
	result["err"] = 0
	result["num"] = 0
	result["result"] = ""

	var l_uid int64
	var l_err error
	if l_uid, l_err = userPower.GetInt64("uid"); l_err != nil {
		result["err"] = -1
		result["result"] = "没有找到用户信息"
		logs.Info("uid === %s", l_uid)
		userPower.Data["json"] = result
		userPower.ServeJSON()
	}

	var l_userpower models.UserPower
	if l_userpower, l_err = models.GetPower(l_uid); l_err != nil {
		result["err"] = -2
		result["result"] = "获取power信息失败"
		userPower.Data["json"] = result
		userPower.ServeJSON()
	}
	result["result"] = l_userpower
	userPower.Data["json"] = result
	userPower.ServeJSON()
}

func (userPower *UserPowerController) DeletePower() {
	result := make(map[string]interface{})
	result["err"] = 0
	result["num"] = 0
	result["result"] = ""

	var _l_user_id int64
	var _l_err  error
	var _l_num  int64

	if _l_user_id , _l_err = userPower.GetInt64("user_id"); _l_err != nil{
		result["err"] = -1
		result["result"] = "没有得到正确的参数"
		logs.Info("user id 的结果是 %d", _l_user_id)
		userPower.Data["json"] = result
		userPower.ServeJSON()
	}

	if _l_num, _l_err = models.DelPowerById(_l_user_id); _l_err != nil{
		logs.Info("被操作的条目 %d", _l_num)
		result["err"] = -2
		result["result"] = "条目不正确"
		userPower.Data["json"] = result
		userPower.ServeJSON()
	}

	userPower.Data["json"] = result
	userPower.ServeJSON()
}

func (userPower *UserPowerController) Login() {
	result := make(map[string]interface{})
	result["err"] = 0
	result["num"] = 0
	username := userPower.GetString("username")
	password := userPower.GetString("password")
	utime := userPower.GetString("utime")

	o := orm.NewOrm()
	var user models.User

	o.QueryTable("user").Filter("username", username).One(&user)
	logs.Error("err: %s", o.Read(&user))
	if o.Read(&user) != nil {
		result["result"] = "用户名或者密码错误"
		result["err"] = -1
		userPower.Data["json"] = result
		userPower.ServeJSON()
	}

	logs.Info("user name : %s", user)

	in_password := user.Password

	//------------------------------------------------------------------------------//
	// 判断 加密 是不是对的
	// 用户名 + “neplite” + 密码 + "iampassword" + unix 时间戳
	password_key := username + "neplite" + in_password + "iampassword" + utime
	// 判断是不是 加密正常

	beego.Warn("明文密码是： - %s -", in_password)

	data := []byte(password_key)
	has := md5.Sum(data)
	md5string := fmt.Sprintf("%x", has)

	if md5string != password {
		result["err"] = -2
		result["result"] = "密码不正确  正确的密码应该是" + md5string
		userPower.Data["json"] = result
		userPower.ServeJSON()
	}
	//------------------------------------------------------------------------------//

	//------------------------------------------------------------------------------//
	// 反馈密码正确 ！  并且 反馈一个密码 交给客户端去判断
	fmt.Println(password)
	// 内密码 + username + utime + “iampassword”
	// 加密
	r_password_key := in_password + username + utime + "iampassword"

	data2 := []byte(r_password_key)
	has2 := md5.Sum(data2)
	r_passsowrd := fmt.Sprintf("%x", has2)

	logs.Info("我的密码是% s", r_passsowrd)
	//beego.Warning("我的密码是%s", r_passsowrd)
	result["result"] = r_passsowrd
	result["err"] = 0
	userPower.Data["json"] = result
	userPower.ServeJSON()
}
