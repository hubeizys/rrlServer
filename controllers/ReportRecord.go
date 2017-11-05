package controllers

import (
	"github.com/astaxie/beego"
	"nepliteApi/comm"
	"nepliteApi/models"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

type ReportRecordConrtroller struct {
	beego.Controller
}

func (record *ReportRecordConrtroller) AddRecord() {
	result := comm.Result{Ret: map[string]interface{}{"err": "aaa", "num": 0, "result": ""}}
	var rp_cprd  models.ReportRecord
	if err := json.Unmarshal(record.Ctx.Input.RequestBody, &rp_cprd); err!= nil{
		result.SetValue("-1", 0, "解析上传结果出现错误")
		logs.Info(err)
		record.Data["json"] = result.Get()
		record.ServeJSON()
	}
	if id, err:= models.ADDRcord(rp_cprd); err != nil{
		result.SetValue("-2", 0, err)
		record.Data["json"] = result.Get()
		record.ServeJSON()
	}else {
		result.SetValue("", 0, id)
	}
	record.Data["json"] = result.Get()
	record.ServeJSON()
}



func (record *ReportRecordConrtroller) GetRecordByname() {
	result := comm.Result{Ret: map[string]interface{}{"err": "aaa", "num": 0, "result": ""}}
	var query_start int
	var query_limit int
	var err error

	id := record.GetString("id")
	/*
	if id = record.GetString("id"); err != nil {
		result.SetValue("-10", 0, "ID异常")
		record.Data["json"] = result.Get()
		record.ServeJSON()
	}*/

	if query_start, err = record.GetInt("start"); err != nil {
		result.SetValue("-1", 0, "参数开始位置异常")
		record.Data["json"] = result.Get()
		record.ServeJSON()
	}

	if query_limit, err = record.GetInt("limit"); err != nil {
		result.SetValue("-2", 0, "参数极限位置异常")
		record.Data["json"] = result.Get()
		record.ServeJSON()
	}

	if _lr_list, err, num := models.GetRecordByName(id, query_start, query_limit); err != nil {
		result.SetValue("-3", num, _lr_list)
		record.Data["json"] = result.Get()
		record.ServeJSON()
	} else {
		result.SetValue("0", num, _lr_list)
	}

	record.Data["json"] = result.Get()
	record.ServeJSON()

}
