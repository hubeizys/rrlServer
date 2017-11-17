package controllers

import (
	"github.com/astaxie/beego"
	"nepliteApi/comm"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"nepliteApi/models"
)

type GroupController struct {
	beego.Controller
}

func (group *GroupController) Add() {
	result := comm.Result{Ret: map[string]interface{}{"err": "", "num": 0, "result": ""}}
	logs.Info("result : == ", result.Ret)

	var _l_group models.Group

	logs.Info("请求数据是 ：%s ", group.Ctx.Input.RequestBody)
	if err := json.Unmarshal(group.Ctx.Input.RequestBody, &_l_group); err != nil {
		result.SetValue("-1", 0, err)
		group.Data["json"] = result.Get()
		group.ServeJSON()
	}
	logs.Warn("asdasd  %s", _l_group)

	if _l_group, err := models.GetGroupByUserID(_l_group.GroupMasterID); err == nil {
		result.SetValue("-2", 0, "已经添加过了， 请不要重复添加店名")
		group.Data["json"] = result.Get()
		group.ServeJSON()
		logs.Warn("_l_group ==  %s; %s", _l_group, err)
	}
	logs.Warn("dasdasdasdasdsadsad")
	if id, err := models.AddGroup(_l_group); err != nil {
		result.SetValue("-2", 0, err)
		group.Data["json"] = result.Get()
		group.ServeJSON()
	} else {
		result.SetValue("0", 0 , id)
	}
	group.Data["json"] = result.Get()
	group.ServeJSON()
}

func (group *GroupController) GetByUserId() {
	result := comm.Result{Ret: map[string]interface{}{"err": "", "num": 0, "result": ""}}
	logs.Info("result : == ", result.Ret)

	//var _l_group models.Group
	var user_id string
	if user_id = group.GetString("user_id"); user_id == "" {
		result.SetValue("-1", 0, user_id)
		group.Data["json"] = result.Get()
		group.ServeJSON()
	}

	if _l_group, err := models.GetGroupByUserID(user_id); err != nil {
		result.SetValue("-2", 0, _l_group)
		group.Data["json"] = result.Get()
		group.ServeJSON()
	} else {
		result.SetValue("0", 0, _l_group)
	}
	group.Data["json"] = result.Get()
	group.ServeJSON()

}
