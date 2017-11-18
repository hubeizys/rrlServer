package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"nepliteApi/models"
)

type SalesController struct {
	beego.Controller
}


func (sales * SalesController) Getall(){
	result :=make(map[string]interface{})
	var sssal []models.Sales

	MasterID,_ := sales.GetInt64("MasterID")
	o:=orm.NewOrm()
	num, error := o.QueryTable("sales").Filter("master_i_d", MasterID).All(&sssal)
	result["err"]  =error
	result["num"]  =num
	result["result"] = sssal
	sales.Data["json"] = result
	sales.ServeJSON()
}

func (sales * SalesController) Getall2User(){
	result :=make(map[string]interface{})
	CardNum := sales.GetString("CardNum")
	var sssal []models.Sales
	o:=orm.NewOrm()
	num, error := o.QueryTable("sales").Filter("card_num",CardNum).All(&sssal)
	result["err"]  =error
	result["num"]  =num
	result["result"] = sssal
	sales.Data["json"] = result
	sales.ServeJSON()
}

type Yingli struct {
	Riqi 	string
	Yingli  string
}
func (sales * SalesController) Yingli() {
	result :=make(map[string]interface{})
	var ylarr []Yingli
	o:=orm.NewOrm()
	qb, _ :=  orm.NewQueryBuilder("mysql")
	qb.Select("DATE(create_date) as  riqi, sum(jia_g_e) as yingli").From("sales").GroupBy("DATE(create_date)")
	sql := qb.String()
	beego.Info("sqlasdasd", sql)
	num, err2 := o.Raw(sql).QueryRows(&ylarr)
	result["err"] = err2
	result["num"] = num
	result["result"] = ylarr
	sales.Data["json"] = result
	sales.ServeJSON()
}



func (sales * SalesController) Add(){
	result :=make(map[string]interface{})

	var sale models.Sales
	sale.CardNum = sales.GetString("CardNum")
	sale.GoodsName = sales.GetString("GoodsName")
	sale.JiaGE,_ = sales.GetInt("JiaGE")
	sale.Tiyan = sales.GetString("Tiyan")
	sale.Ewai = sales.GetString("Ewai")
	sale.TiJian = sales.GetString("TiJian")
	var err error
	sale.MasterID, err = sales.GetInt64("MasterID")
	if err != nil{
		result["err"] = "-1"
		result["num"] = 0
		result["result"] = err
	}
	o:=orm.NewOrm()

	if err == nil{
		//  jianshao xianji
		var user models.User
		o.QueryTable("user").Filter("CardNum", sale.CardNum).One(&user)
		result["result"] = user
		if o.Read(&user) == nil{
			// int,err:=strconv.Atoi(user.Yueee)
			var num  = user.Yueee - sale.JiaGE
			if num < 0 {
				result["err"]  = "钱不够"

			}else{
				user.Yueee = num
				err_num, up_err := o.Update(&user)
				result["err"]  = up_err
				result["num"]  = err_num
				if up_err != nil{
					num, err := o.Insert(&sale)
					result["err"]  = err
					result["num"]  = num
					result["result"] =  ""
				}
			}
		}
	}
	sales.Data["json"] = result
	sales.ServeJSON()
}