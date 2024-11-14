package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	InitRegisterModel()
	initDatabase()
	InitData()
}
func InitRegisterModel() {
	orm.RegisterModel(new(SysInfo))

	//orm.RegisterModel(new(FaRaceInfo))
	orm.RegisterModel(new(FaRaceInfo))
	orm.RegisterModel(new(FaTeamInfo))
	orm.RegisterModel(new(FaCustomRankingInfo))
	//orm.RegisterModel(new(FaRaceInfo))
	//orm.RegisterModel(new(FaRaceInfoHc))
	orm.RegisterModel(new(FaRaceSetup))
	orm.RegisterModel(new(FaHistoryRaceInfoExt))
	orm.RegisterModel(new(FaIntegralRankingInfo))
	orm.RegisterModel(new(FaFutureRaceInfo))

	orm.RegisterModel(new(FaHistoryAsiaOdds))
	orm.RegisterModel(new(FaHistoryEuropeOdds))
	orm.RegisterModel(new(FaInjury))

	orm.RegisterModel(new(FaAsiaOddsInfo))
	orm.RegisterModel(new(FaEuropeOddsInfo))
	orm.RegisterModel(new(FaGSOddsInfo))
	orm.RegisterModel(new(FaExcelFData))

	orm.RegisterModel(new(FaAuthorityCompany))
	orm.RegisterModel(new(FaAuthorityCompanyAlias))

}
func InitData() {
	InitSysInfo(nil)
	InitAuthorityCompany(nil)
	InitAuthorityCompanyAlias(nil)
}

// 初始化数据连接
func initDatabase() {
	//读取配置文件，设置数据库参数
	//数据库类别
	dbType := "mysql"
	//连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	//数据库端口
	dbPort := beego.AppConfig.String(dbType + "::db_port")

	switch dbType {
	case "sqlite3":
		orm.RegisterDataBase(dbAlias, dbType, dbName)
	case "mysql":
		dbCharset := beego.AppConfig.String(dbType + "::db_charset")
		//dataS := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&loc=Asia%2FShanghai"
		dataS := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&loc=Local"
		//orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset+"&loc=Asia%2FShanghai", 30)
		e := orm.RegisterDataBase(dbAlias, dbType, dataS, 30)
		if e != nil {
			panic(e)
		}
		db, e := orm.GetDB(dbAlias)
		if e == nil {
			db.SetConnMaxLifetime(0)
		}
	}

	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	//自动建表
	orm.RunSyncdb("default", false, true)
	if isDev {
		orm.Debug = isDev
	}

}
