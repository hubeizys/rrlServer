package controllers

import (
	"nepliteApi/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
	"strconv"
	"github.com/astaxie/beego/logs"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// uid := models.AddUser(user)
	u.Data["json"] = map[string]int{"uid": 1}
	u.ServeJSON()
}

func (c *UserController) URLMapping() {
	logs.Info("nihaohahahah")
	c.Mapping("Info", c.Info)
}


func (u *UserController) UpdateUserBase() {
	result :=make(map[string]interface{})
	username := u.GetString("username")
	cardnum:= u.GetString("cardnum")
	phonenum := u.GetString("phonenum")
	JianKang := u.GetString("JianKang")
	xingbie := u.GetString("xingbie")
	shengri := u.GetString("shengri")
	shengao := u.GetString("shengao")
	tizhong := u.GetString("tizhong")

	o := orm.NewOrm()
	//user := User{Id: 1}
	var user models.User
	o.QueryTable("user").Filter("card_num", cardnum).One(&user)
	result["err"] =  -1
	result["result"] =  ""
	result["num"] =  0
	beego.Info("cardsadasdasdasdasdasdasd",cardnum, user, o.Read(&user))
	if o.Read(&user) == nil {
		user.Username = username
		user.Phone = phonenum
		user.Jiankangzhishu = JianKang
		user.Gender = xingbie
		user.Shengao = shengao
		user.Tizhong = tizhong
		timeLayout := "2006-01-02 15:04:05"
		loc, _ := time.LoadLocation("Local")
		theTime, eee := time.ParseInLocation(timeLayout, shengri, loc)
		beego.Info("the time @@@@@", theTime, shengri, eee)
		user.CreateDate = theTime
		beego.Info("asdasd", user)
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
			result["num"] = num
			result["err"] = err
		}
	}
	u.Data["json"]  = result
	u.ServeJSON()
}



func (u *UserController) UpdateUserYuee() {
	result :=make(map[string]interface{})
	card_num := u.GetString("card_num")
	yuee :=u.GetString("yuee")
	card_type := u.GetString("cardtype")
	o := orm.NewOrm()
	//user := User{Id: 1}
	var user models.User
	o.QueryTable("user").Filter("card_num", card_num).One(&user)
	result["err"] =  -1
	result["result"] =  ""
	result["num"] =  0
	beego.Info("cardsadasdasdasdasdasdasd",card_num, user, o.Read(&user))
	if o.Read(&user) == nil {
		real_int,_:=strconv.Atoi(yuee)
		user.Yueee = real_int
		user.UserType = card_type
		beego.Info("asdasd", user)
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
			result["num"] = num
			result["err"] = err
		}
	}
	u.Data["json"]  = result
	u.ServeJSON()
}

func (u *UserController) FindByCard() {
	result :=make(map[string]interface{})
	if  u.GetString("cardnum") != ""{

		var users []models.User
		o := orm.NewOrm()
		num, err:= o.QueryTable("user").Filter("card_num", u.GetString("cardnum")).All(&users)

		result["num"] = num
		result["result"] = users
		result["err"] = err
		u.Data["json"]  = result
		u.ServeJSON()
	}else {
		result["err"] = -1
		u.Data["json"]  = result
		u.ServeJSON()
	}
}



func (u *UserController) Add() {
	result :=make(map[string]interface{})
	o := orm.NewOrm()
	var user models.User
	user.Username = u.GetString("username")
	user.CardNum = u.GetString("cardnum")
	user.Phone = u.GetString("phonenum")
	user.Jiankangzhishu = u.GetString("JianKang")
	user.Gender = u.GetString("xingbie")


	toBeCharge := u.GetString("shengri")                         //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	user.CreateDate =  theTime
	user.Shengao = u.GetString("shengao")
	user.Tizhong = u.GetString("tizhong")


	num , err :=  o.Insert(&user)
	result["num"] = num
	result["result"] = num
	result["err"] = err
	u.Data["json"]  = result
	u.ServeJSON()
}
// @router /info [get]
func (u *UserController) Info() {
	result :=make(map[string]interface{})
	var users []models.User
	o := orm.NewOrm()
	num, err:= o.QueryTable("user").All(&users)

	logs.Warn("aasdasdasda")
	result["num"] = num
	result["result"] = users
	result["err"] = err
	u.Data["json"]  = result
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	// users := models.GetAllUsers()
	u.Data["json"] = "{}"
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {

	/*
	uid,_ := u.GetInt(":uid")
	if uid != 0 {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}*/

	/*
	o := orm.NewOrm()
	o.Using("default")
	profile := new(models.Profile)
	profile.Age = 24
	profile.Address = "zhongguo"
	profile.Gender = "male"
	profile.Email = "124296140@qq.com"

	user := new(models.User)
	user.Username = "zhuyunsong"
	user.Password = "123213"
	user.Profile = profile

	o.Insert(profile)
	o.Insert(user)
	u.ServeJSON()
	*/
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {

	//uid,_ := u.GetInt(":uid")

	/*
	if uid != 0 {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}*/

	u.Data["json"] = "{}"
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	// uid,_ := u.GetInt(":uid")
	// models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	/*
	username := u.GetString("username")
	password := u.GetString("password")

	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}*/
	u.Data["json"] = "login success"
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

