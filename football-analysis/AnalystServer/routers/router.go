package routers

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func Init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/Get", &controllers.MainController{}, "*:Get")
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.FApiController{})

	beego.SetStaticPath("/assets", "static/assets")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"token", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		AllowCredentials: true,
	}))
	beego.Run()
}
