package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"nepliteApi/models"
)

type PanDianController struct {
	beego.Controller
}

type pandian struct {
	Goods_p_d_i_d    string
	Goods_p_d_name   string
	Goods_p_d_result string
	Goods_p_d_time   string
	Username         string
}

func (pandianObj *PanDianController) Add() {

	GoodsPDName := pandianObj.GetString("GoodsPDName")
	GoodsPDResult := pandianObj.GetString("GoodsPDResult")
	result := make(map[string]interface{})
	result["err"] = nil
	result["num"] = 0
	result["result"] = ""
	o := orm.NewOrm()
	var gpd models.GoodsPD
	gpd.GoodsPDName = GoodsPDName
	gpd.GoodsPDResult = GoodsPDResult
	gpd.ID = 1
	num, err := o.Insert(&gpd)
	beego.Info("num, err := o.Insert(&gpd)", num, err)
	result["err"] = err
	result["num"] = num
	pandianObj.Data["json"] = result
	pandianObj.ServeJSON()
}

func (pandianObj *PanDianController) Get() {
	result := make(map[string]interface{})
	//var pandian  []models.GoodsPD

	//num, err:= o.QueryTable("goods_p_d").All(&pandian)
	o := orm.NewOrm()
	var goodsPD []pandian
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("goods_p_d.goods_p_d_i_d,goods_p_d.goods_p_d_name,goods_p_d.goods_p_d_result,goods_p_d.goods_p_d_time,user.username ").
		From("goods_p_d").
		LeftJoin("user").
		On("goods_p_d.i_d = user.id")
	sql := qb.String()
	num, err2 := o.Raw(sql, 20).QueryRows(&goodsPD)

	result["result"] = goodsPD
	result["num"] = num
	result["err"] = err2

	pandianObj.Data["json"] = result
	pandianObj.ServeJSON()
}
