package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"nepliteApi/models"
	"time"
	"fmt"
	"nepliteApi/comm"
	"github.com/astaxie/beego/logs"
)

type GoodsRecordController struct {
	beego.Controller
}

func (goodRcd *GoodsRecordController) Get() {
	result := make(map[string]interface{})
	var goods_records []models.GoodsRecord
	o := orm.NewOrm()
	num, error := o.QueryTable("goods_record").All(&goods_records)
	result["num"] = num
	result["err"] = error
	result["result"] = goods_records
	goodRcd.Data["json"] = result
	goodRcd.ServeJSON()
}

type GoodsRDINFO struct {
	Op_goods_i_d string
	GoodId       int
	Name         string
	Create_date  string
	Username     string
	Goods_bus    string
	Goods_type   string
	OpNum        int
	Goods_price  int
}

func (goodRcd *GoodsRecordController) Chuku() {
	result := comm.Result{Ret: map[string]interface{}{"err": "", "num": 0, "result": ""}}
	logs.Info("result : == ", result.Ret)

	//	result :=make(map[string]interface{})
	o := orm.NewOrm()
	var gds []GoodsRDINFO
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("goods_record.op_goods_i_d,goods_record.name,goods_record.Goods_price,goods_record.OpNum, goods_record.goods_bus,goods_record.goods_type,goods_record.goods_price, goods_record.create_date, user.username").
		From("goods_record").
		LeftJoin("user").
		On("goods_record.op_user = user.id").
		Where(`op = "出库"`)

	sql := qb.String()
	if num, err2 := o.Raw(sql, 20).QueryRows(&gds); err2 != nil {
		beego.Info("qb=============", num, "123123123123", err2, gds)
		result.SetValue("-1", num, err2)
		goodRcd.Data["json"] = result.Get()
		goodRcd.ServeJSON()
	}

	result.SetValue("0", 0, gds)
	/*result["num"] = num
	result["err"] = err2
	result["result"] = gds
	*/
	goodRcd.Data["json"] = result
	goodRcd.ServeJSON()
}

func (goodRcd *GoodsRecordController) Ruku() {

	result := comm.Result{Ret: map[string]interface{}{"err": "", "num": 0, "result": ""}}
	logs.Info("result : == ", result.Ret)
	/*
		result :=make(map[string]interface{})
	*/
	// var goods_records  []models.GoodsRecord
	o := orm.NewOrm()

	var gds []GoodsRDINFO
	/*
	var maps []orm.Params
	num, error := o.QueryTable("goods_record").RelatedSel().Values(&maps,"username")
	beego.Info("err", error, maps)
	qb, err :=  orm.NewQueryBuilder("mysql")
	beego.Info("qb=============", qb)

	qb.Select("goods_record.name").
		From("goods_record").
		LeftJoin("user").
		On("goods_record.op_user_id = user.id")
	*/
	qb, _ := orm.NewQueryBuilder("mysql")
	beego.Info("qb=============", qb)

	qb.Select("goods_record.name,goods_record.Goods_price,goods_record.OpNum, goods_record.goods_bus,goods_record.goods_type,goods_record.goods_price, goods_record.create_date, user.username").
		From("goods_record").LeftJoin("user").On("goods_record.op_user = user.id").
		Where(`op = "入库"`)
	sql := qb.String()
	beego.Info("qb=============", sql)
	if num, err2 := o.Raw(sql).QueryRows(&gds); err2 != nil {
		result.SetValue("-1", num, err2)
		logs.Warn("qb=============%s|%s|%s|%s", num, "123123123123", err2, gds)
		goodRcd.Data["json"] = result.Get()
		goodRcd.ServeJSON()
	} else {
		result.SetValue("0", num, gds)
	}
	/*
	result["num"] = num
	result["err"] = err
	result["result"] = gds*/
	goodRcd.Data["json"] = result
	goodRcd.ServeJSON()
}

// chuku
func (goodRcd *GoodsRecordController) Del() {
	/*    "Op_goods_i_d": "123",
		"GoodId": 0,
		"Name": "sdas",
		"Create_date": "2017-09-16 00:00:00 +0000 UTC",
		"Username": "",
		"Goods_bus": "123123",
		"Goods_type": "41212",
		"OpNum": 0,
		"Goods_price": 123*/

	GoodId, _ := goodRcd.GetInt("GoodId")
	Name := goodRcd.GetString("Name")
	Create_date := goodRcd.GetString("Create_date")

	Goods_bus := goodRcd.GetString("Goods_bus")
	Goods_type := goodRcd.GetString("Goods_type")
	OpNum, _ := goodRcd.GetInt64("OpNum")
	Goods_price, _ := goodRcd.GetInt64("Goods_price")

	result := make(map[string]interface{})

	o := orm.NewOrm()
	var goods_record models.GoodsRecord
	goods_record.Name = Name
	goods_record.OpGoodsID = GoodId

	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, Create_date, loc)

	goods_record.CreateDate = theTime
	goods_record.GoodsBus = Goods_bus
	goods_record.GoodsType = Goods_type
	goods_record.Op = "出库"
	goods_record.OpNum = OpNum
	goods_record.GoodsPrice = Goods_price

	// jianshao
	// zhenduo
	temp_good := models.Goods{GoodsId: GoodId}
	err := o.Read(&temp_good)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		// fmt.Println(user.Id, user.Name)
		temp_good.Count -= OpNum

		o.Update(&temp_good)
	}
	id, errr := o.Insert(&goods_record)
	beego.Info("id", id, "err", errr)
	result["err"] = errr
	result["result"] = id
	result["num"] = 1

	goodRcd.Data["json"] = result
	goodRcd.ServeJSON()

}

//ruku
func (goodRcd *GoodsRecordController) Add() {
	GoodId, _ := goodRcd.GetInt("GoodId")
	Name := goodRcd.GetString("Name")
	Create_date := goodRcd.GetString("Create_date")

	Goods_bus := goodRcd.GetString("Goods_bus")
	Goods_type := goodRcd.GetString("Goods_type")
	OpNum, _ := goodRcd.GetInt64("OpNum")
	Goods_price, _ := goodRcd.GetInt64("Goods_price")

	result := make(map[string]interface{})

	o := orm.NewOrm()
	var goods_record models.GoodsRecord
	goods_record.Name = Name
	goods_record.OpGoodsID = GoodId

	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, Create_date, loc)

	goods_record.CreateDate = theTime
	goods_record.GoodsBus = Goods_bus
	goods_record.GoodsType = Goods_type
	goods_record.Op = "入库"
	goods_record.OpNum = OpNum
	goods_record.GoodsPrice = Goods_price

	// zhenduo
	temp_good := models.Goods{GoodsId: GoodId}
	err := o.Read(&temp_good)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		o := orm.NewOrm()
		var goods models.Goods
		//goods.Count = add_num
		//goods.Name = add_name

		goods.Count = OpNum
		goods.GoodsId = GoodId
		goods.Name = Name
		goods.Danjia = Goods_price

		num, errr := o.Insert(&goods)
		beego.Info("num ===== err", num, err)
		if errr != nil {
			result["err"] = errr
			result["result"] = num
			result["num"] = 1

			goodRcd.Data["json"] = result
			goodRcd.ServeJSON()
			return
		}
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		// fmt.Println(user.Id, user.Name)
		temp_good.Count += OpNum

		o.Update(&temp_good)
	}

	id, errr := o.Insert(&goods_record)
	beego.Info("id", id, "err", errr)
	result["err"] = errr
	result["result"] = id
	result["num"] = 1

	goodRcd.Data["json"] = result
	goodRcd.ServeJSON()
}
