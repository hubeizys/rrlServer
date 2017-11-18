package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"nepliteApi/models"
	"nepliteApi/comm"
	"github.com/astaxie/beego/logs"
)

type PanDianController struct {
	beego.Controller
}

type pandian struct {
	Goods_p_d_i_d    string
	Goods_p_d_name   string
	Goods_p_d_result string
	Goods_p_d_time   string
	I_d         string
}

func (pandianObj *PanDianController) Add() {
	GoodsPDName := pandianObj.GetString("GoodsPDName")
	GoodsPDResult := pandianObj.GetString("GoodsPDResult")
	UserID,_ := pandianObj.GetInt64("ID")
	result := make(map[string]interface{})
	result["err"] = nil
	result["num"] = 0
	result["result"] = ""
	o := orm.NewOrm()
	var gpd models.GoodsPD
	gpd.GoodsPDName = GoodsPDName
	gpd.GoodsPDResult = GoodsPDResult
	gpd.ID = UserID
	num, err := o.Insert(&gpd)
	beego.Info("num, err := o.Insert(&gpd)", num, err)
	result["err"] = err
	result["num"] = num
	pandianObj.Data["json"] = result
	pandianObj.ServeJSON()
}

func (pandianObj *PanDianController) Get() {
	//result := make(map[string]interface{})
	//var pandian  []models.GoodsPD
	result := comm.Result{Ret: map[string]interface{}{"err": "", "num": 0, "result": ""}}
	logs.Info("result : == ", result.Ret)

	//num, err:= o.QueryTable("goods_p_d").All(&pandian)
	o := orm.NewOrm()
	var goodsPD []pandian
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("goods_p_d.goods_p_d_i_d,goods_p_d.goods_p_d_name,goods_p_d.goods_p_d_result,goods_p_d.goods_p_d_time,goods_p_d.i_d ").
		From("goods_p_d").
		LeftJoin("user_power").
		On("goods_p_d.i_d = user_power.user_i_d")
	sql := qb.String()
	if num, err2 := o.Raw(sql).QueryRows(&goodsPD); err2 != nil {
		result.SetValue("-1", num, err2)
		pandianObj.Data["json"] = result.Get()
		pandianObj.ServeJSON()
	}else {
		result.SetValue("0", num, goodsPD)
	}

	/*
	result["result"] = goodsPD
	result["num"] = num
	result["err"] = err2
	*/
	pandianObj.Data["json"] = result.Get()
	pandianObj.ServeJSON()
}
