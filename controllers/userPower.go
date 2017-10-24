package controllers

import (
	"nepliteApi/models"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"github.com/astaxie/beego/logs"
)

type UserPowerController struct {
	beego.Controller
}

func (userPower *UserPowerController) Add() {
	result := make(map[string]interface{})
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
	logs.Error("errr : %s", o.Read(&user))
	if o.Read(&user) != nil {
		result["result"] = "没有发现用户"
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
