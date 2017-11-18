package controllers

import (
	"github.com/astaxie/beego"
	"nepliteApi/models"
	"github.com/astaxie/beego/orm"
	"nepliteApi/comm"
	"github.com/astaxie/beego/logs"
)

type GoodsController struct {
	beego.Controller
}

func (goodsObj *GoodsController) UpdateYuzhi()  {
	result := comm.Result{Ret: map[string]interface{}{"err": "", "num": 0, "result": ""}}
	logs.Info("result : == ", result.Ret)

	goods_id, _ := goodsObj.GetInt64("goods_id")
	Yuzhi, _ := goodsObj.GetInt64("yuzhi")
	o := orm.NewOrm()
	var goods models.Goods
	if err := o.QueryTable(models.Goods{}).Filter("id", goods_id).One(&goods); err != nil{
		result.SetValue("-1", 0, err)
	}else {
		goods.Yuzhi = Yuzhi
		logs.Warn("gods == %s", goods)
		if num, er2 := o.Update(&goods); er2 != nil{
			result.SetValue("-2", 0 , er2)
		}else {
			result.SetValue("0", num, "")
		}
	}
	goodsObj.Data["json"] = result.Get()
	goodsObj.ServeJSON()

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