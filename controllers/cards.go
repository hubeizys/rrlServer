package controllers

import (
	"github.com/astaxie/beego"
	"nepliteApi/comm"
	"nepliteApi/models"
)

type CardsController struct {
	beego.Controller
}

func (cardsobj *CardsController)Get()  {
	result := comm.Result{Ret: map[string]interface{}{"err": "aaa", "num": 0, "result": ""}}
	id := cardsobj.GetString("id")
	if len(id) != 32 {
		result.SetValue("-1", 0, "卡号不正确")
		cardsobj.Data["json"] = result.Get()
		cardsobj.ServeJSON()
	}
	if card_obj , err := models.GetCardsInfo(id); err== nil {
		if card_obj.Id == ""{
			result.SetValue("-3", 1, "没有发现数据")
		}else {
			result.SetValue("0", 1, card_obj)
		}

		cardsobj.Data["json"] = result.Get()
		cardsobj.ServeJSON()
	}else {
		result.SetValue("-2", 0, err)
		cardsobj.Data["json"] = result.Get()
		cardsobj.ServeJSON()
		}
}

func (cardsobj * CardsController)AddCard(){
	result := comm.Result{Ret: map[string]interface{}{"err": "aaa", "num": 0, "result": ""}}
	id := cardsobj.GetString("id")
	if len(id) != 32 {
		result.SetValue("-1", 0, "卡号不正确")
		cardsobj.Data["json"] = result.Get()
		cardsobj.ServeJSON()
	}
	var card =  models.Cards{Id:id, Remark:"nothing", Cardtype:"1"}
	if num, err := models.AddCard(card); err ==nil{
		result.SetValue("0", num, err)
	}else {
		result.SetValue("-2", num, err)
	}
	cardsobj.Data["json"] = result.Get()
	cardsobj.ServeJSON()
}


