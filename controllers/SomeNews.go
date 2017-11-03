package controllers

import (
	"github.com/astaxie/beego"
	"nepliteApi/models"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"nepliteApi/comm"
)

type SomeNewsController struct {
	beego.Controller
}

func (someNewObj *SomeNewsController) GetAll() {
	result := comm.Result{Ret: map[string]interface{}{"err": "aaa", "num": 0, "result": ""}}
	logs.Info("someNewObj : %s", result.Ret)
	if _l_snlist, num, err := models.GetNews(); err != nil {
		logs.Info("_l_snlist : %s, err %s", _l_snlist, err.Error())
		result.SetValue(err.Error(), num, "获得最新的公告数据出现了错误")
	} else {
		result.SetResult(_l_snlist)
	}
	someNewObj.Data["json"] = result.Get()
	someNewObj.ServeJSON()
}

func (someNewObj *SomeNewsController) Add() {
	result := comm.Result{Ret: map[string]interface{}{"err": "aaa", "num": 0, "result": ""}}
	var _l_temp_somenews models.SomeNews
	logs.Info("请求数据是： %s", someNewObj.Ctx.Input.RequestBody)
	err := json.Unmarshal(someNewObj.Ctx.Input.RequestBody, &_l_temp_somenews)
	if err != nil {
		result.SetValue("-1", 0, err.Error())
		someNewObj.Data["json"] = result.Get()
		someNewObj.ServeJSON()
	}
	if id, err := models.AddNews(_l_temp_somenews); err != nil {
		result.SetValue("-2", 0, err.Error())
		logs.Info("反馈的id是: %d", id)
	} else {
		result.SetResult("")
	}
	someNewObj.Data["json"] = result.Get()
	someNewObj.ServeJSON()
}
