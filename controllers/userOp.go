package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"nepliteApi/models"
)

//   管理用户操作记录

type UserOpController struct {
	beego.Controller
}


func (uop * UserOpController) Getall(){
	result :=make(map[string]interface{})
	result["num"] = 0
	result["err"] = ""
	result["result"] = ""

	// var _l_user_op_list  []models.UserOp
	if  fl_l_user_op_list, num, err := models.GetAllUserOp(); err !=nil{
		result["num"] = num
		result["err"] = err
		result["result"] = ""
		uop.Data["json"] = result
		uop.ServeJSON()
	} else {
		result["result"] = fl_l_user_op_list
	}

	uop.Data["json"] = result
	uop.ServeJSON()
}


func (uop *UserOpController) Add() {
	result := make(map[string]interface{})
	result["num"] = 0
	result["err"] = ""
	result["result"] = ""

	var _l_user_op models.UserOp
	err := json.Unmarshal(uop.Ctx.Input.RequestBody, &_l_user_op)
	if err != nil {
		result["err"] = -1
		result["result"] = err
		uop.Data["json"] = result
		uop.ServeJSON()
	}
	id, op_err := models.AddUserOp(_l_user_op)

	if op_err != nil {
		result["err"] = -2
		result["result"] = op_err
		uop.Data["json"] = result
		uop.ServeJSON()
	}

	result["result"] = id
	uop.Data["json"] = result
	uop.ServeJSON()
}

func (uop * UserOpController) Delelte()  {
	result := make(map[string]interface{})
	result["num"] = 0
	result["err"] = ""
	result["result"] = ""

	op_id, err := uop.GetInt64("oprecord_id")
	if err != nil {
		result["num"] = 0
		result["err"] = -1
		result["result"] = "解析id出错"
		uop.Data["json"] = result
		uop.ServeJSON()
	}

	if err := models.DelUserOp(op_id); err != nil{
		result["num"] = 0
		result["err"] = -2
		result["result"] = "删除高级用户记录出现了错误"
		uop.Data["json"] = result
		uop.ServeJSON()
	}
	uop.Data["json"] = result
	uop.ServeJSON()
}