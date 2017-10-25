package controllers

import (
	"nepliteApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserSHController struct {
	beego.Controller
}

func (usersh *UserSHController) Add() {
	result := make(map[string]interface{})
	cardnum, _ := usersh.GetInt("cardnum")
	Status := usersh.GetString("Status")
	// create_date := usersh.GetString("create_date")
	TJxiangmu := usersh.GetString("TJxiangmu")
	XFxiangmu := usersh.GetString("XFxiangmu")
	result["result"] = ""
	result["num"] = 0
	result["err"] = nil

	o := orm.NewOrm()
	var usersssss models.UserSH
	usersssss.UserID = cardnum
	usersssss.Status = Status
	usersssss.TJxiangmu = TJxiangmu
	usersssss.XFxiangmu = XFxiangmu
	num, err := o.Insert(&usersssss)

	result["num"] = num
	result["err"] = err
	usersh.Data["json"] = result
	usersh.ServeJSON()

}

func (usersh *UserSHController) GetAll() {

	result := make(map[string]interface{})
	var users []models.UserSH
	o := orm.NewOrm()
	num, err := o.QueryTable("user_s_h").All(&users)
	//var gpd models.GoodsPD

	result["num"] = num
	result["result"] = users
	result["err"] = err
	usersh.Data["json"] = result
	usersh.ServeJSON()
}
