package controllers

import (
	"github.com/astaxie/beego"
	"nepliteApi/models"
	"github.com/astaxie/beego/orm"

)

type GoodsController struct {
	beego.Controller
}

func (goodsObj *GoodsController) Get()  {
	result :=make(map[string]interface{})

	var goods []models.Goods

	o := orm.NewOrm()
	num, err:= o.QueryTable("goods").All(&goods)
	result["result"] = goods
	result["num"] = num
	result["err"] = err
	goodsObj.Data["json"] = result
	goodsObj.ServeJSON()
}

func (goodsObj *GoodsController) ShowAll()  {
	result :=make(map[string]interface{})
	userid,_ := goodsObj.GetInt64("userid")
	var goods []models.Goods
	o := orm.NewOrm()
	num, err:=o.QueryTable("goods").Filter("create_user", userid).Limit(100, 0).All(&goods)
	result["result"] = goods
	result["num"] = num
	result["err"] = err
	goodsObj.Data["json"] = result
	goodsObj.ServeJSON()
}

func (goodsObj *GoodsController) Add(){
	result := make(map[string]interface{})

	add_num, _ := goodsObj.GetInt64("count")
	add_name := goodsObj.GetString("name")

	if add_name == ""{
		result["err"] = "00001"
	}else {
		o := orm.NewOrm()
		var goods  models.Goods
		goods.Count = add_num
		goods.Name = add_name
		id, errr := o.Insert(&goods)
		result["err"] = errr
		result["num"] = 1
		result["result"] = id
	}
	goodsObj.Data["json"] = result
	goodsObj.ServeJSON()
}