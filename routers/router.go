// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"nepliteApi/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	user_ns := beego.NewNamespace("/user",
		beego.NSNamespace("/object", // todo 就是留着占个位置，不然不好看
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),

		// Todo 注解路由太骚包了。 不建议使用。
		// TODO 自动挡把接口隐藏了， 处女座程序员可以用一用
		beego.NSRouter("/map", &controllers.UserController{}, "get:URLMapping"),
		beego.NSRouter("/update", &controllers.UserController{}, "get:UpdateUserYuee"),
		beego.NSRouter("/updateBase", &controllers.UserController{}, "get:UpdateUserBase"),
		beego.NSRouter("/add", &controllers.UserController{}, "get:Add"),
		beego.NSRouter("/info", &controllers.UserController{}, "get:Info"),
		beego.NSRouter("/infobycard", &controllers.UserController{}, "get:FindByCard"),
		beego.NSNamespace("/neplite",
			beego.NSGet(":id", func(context *context.Context) {
				context.Output.Body([]byte("notAllowed"))
			}),

		),
		beego.NSNamespace("/shouhou",
			beego.NSRouter("all", &controllers.UserSHController{}, "get:GetAll"),
			beego.NSRouter("/add", &controllers.UserSHController{}, "get:Add"),

		),

		beego.NSNamespace("/sales",
			beego.NSRouter("/getalltouser", &controllers.SalesController{}, "get:Getall2User"),
			beego.NSRouter("/yingli", &controllers.SalesController{}, "get:Yingli"),
			beego.NSRouter("/getall", &controllers.SalesController{}, "get:Getall"),
			beego.NSRouter("/add", &controllers.SalesController{}, "get:Add"),
			beego.NSRouter("/reportadd", &controllers.ReportRecordConrtroller{}, "post:AddRecord"),
			beego.NSRouter("/report", &controllers.ReportRecordConrtroller{}, "get:GetRecordByname"),
		),

		beego.NSNamespace("/power",
			beego.NSRouter("/add", &controllers.UserPowerController{}, "post:Add"),
			beego.NSRouter("/login", &controllers.UserPowerController{}, "get:Login"),
			beego.NSRouter("/test/:uid", &controllers.UserPowerController{}, "get:PutPower"),
			beego.NSRouter("/getpower", &controllers.UserPowerController{}, "get:GetPower"),
			beego.NSRouter("/all", &controllers.UserPowerController{}, "get:GetAll"),
			beego.NSRouter("/del", &controllers.UserPowerController{}, "get:DeletePower"),
			beego.NSRouter("/getnpower", &controllers.UserPowerController{}, "get:GetNormalPower"),
			beego.NSRouter("/update", &controllers.UserPowerController{}, "post:UpdatePower"),
		),
	)

	comm_ns := beego.NewNamespace("/comm",
		beego.NSRouter("/get", &controllers.SomeNewsController{}, "get:GetAll"),
		beego.NSRouter("/add", &controllers.SomeNewsController{}, "post:Add"),
	)

	wuliao_ns := beego.NewNamespace("wuliao",
		beego.NSInclude(
			&controllers.NepliteControllers{},
		),
		beego.NSRouter("/pandian", &controllers.PanDianController{}),
		beego.NSRouter("/pandianadd", &controllers.PanDianController{}, "get:Add"),
		beego.NSRouter("/all", &controllers.GoodsController{}),
		beego.NSRouter("/show", &controllers.GoodsController{}, "get:ShowAll"),
		beego.NSRouter("/add", &controllers.GoodsController{}, "get:Add"),
		beego.NSNamespace("/record",
			beego.NSRouter("/ruku", &controllers.GoodsRecordController{}, "get:Ruku"),
			beego.NSRouter("/chuku", &controllers.GoodsRecordController{}, "get:Chuku"),
			beego.NSRouter("/add", &controllers.GoodsRecordController{}, "get:Add"),
			beego.NSRouter("/del", &controllers.GoodsRecordController{}, "get:Del"),
		),
	)
	beego.AddNamespace(user_ns, wuliao_ns, comm_ns)
}
