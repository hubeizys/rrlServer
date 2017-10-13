package controllers

import "github.com/astaxie/beego"

type NepliteControllers struct {
	beego.Controller
}

//@router / [get]
func (o *NepliteControllers) Get() {
	o.Data["json"] = map[string]int{"zhuyunsong": 11231}
	o.ServeJSON()
}

//@router /send [get]
func (o *NepliteControllers) Sent() {
	o.Data["json"] = map[string]int{"neplite": 111}
	o.ServeJSON()
}