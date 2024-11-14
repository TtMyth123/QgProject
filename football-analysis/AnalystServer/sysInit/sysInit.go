package sysInit

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/routers"
	"github.com/astaxie/beego"
)

func Init() {
	models.Init()
	routers.Init()
	beego.Run()
}
