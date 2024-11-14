package analyst

import (
	"fmt"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"testing"
	"time"
)

func init1() {
	ttLog.InitLogs()
	models.InitRegisterModel()

	dbAlias := "default"
	dbUser := "root"
	dbPwd := "root"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "test_football"
	dbCharset := "utf8"
	dbType := "mysql"
	dataS := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset + "&loc=Local"
	e := orm.RegisterDataBase(dbAlias, dbType, dataS, 30)
	db, e := orm.GetDB(dbAlias)
	if e == nil {
		db.SetConnMaxLifetime(0)
	}
	orm.RunSyncdb("default", false, true)
}
func Test_001GetAndAddRaceTmpList(t *testing.T) {
	init1()
	aWork := NewAnalystWork(2, "./conf/A.xlsx", basePath1, basePath2, basePath3)
	aWork.GetAndAddRaceTmpList()
}

func Test_002GetAndAddHcRaceTmpList(t *testing.T) {
	init1()
	aWork := NewAnalystWork(2, "./conf/A.xlsx", basePath1, basePath2, basePath3)
	aWork.GetAndAddHcRaceTmpList("2022-11-18")
}

func Test_repairTeamName(t *testing.T) {
	str1 := `天津津门虎<font color=#880000>(中)</font>`
	str := repairTeamName(str1)
	fmt.Println(str)

	str1 = `梅塔利斯特U21`
	str = repairTeamName(str1)

	fmt.Println(str)
}

func Test_003GetAndSaveRaceData(t *testing.T) {
	init1()
	aWork := NewAnalystWork(2, "./conf/A.xlsx", basePath1, basePath2, basePath3)
	aWork.GetAndSaveRaceData(30)
}

func Test_004ExcelFormulaData(t *testing.T) {
	init1()
	aWork := NewAnalystWork(2, "./conf/A.xlsx", basePath1, basePath2, basePath3)
	//for i := 0; i < 10; i++ {
	data := aWork.mExcelFormulaData.GetExcelFormulaData()
	ttLog.LogDebug(stringKit.GetJsonStr(data))
	//}

}

func Test_004GetEuropeOddsInfoData(t *testing.T) {
	init1()
	aWork := NewAnalystWork(2, "./conf/A.xlsx", basePath1, basePath2, basePath3)
	//for i := 0; i < 10; i++ {
	aRaceInfo := models.FaRaceInfo{}
	aRaceInfo.RaceInfoId = 2310427

	data, _ := aWork.GetEuropeOddsInfoData(aRaceInfo)
	ttLog.LogDebug(stringKit.GetJsonStr(data))
	//}

}

func Test_004Time(t *testing.T) {
	strT := `11-20 02:07`
	tt, e := time.Parse(`01-02 15:04`, strT)
	fmt.Println(tt, e)
}

func Test_004Time2(t *testing.T) {
	StrBeginOddsTime := "2022,11-1,19,18,50,00"
	trBeginOddsTimeL := "2006,1-2,15,04,05,00"

	BeginOddsTime, e := time.Parse(trBeginOddsTimeL, StrBeginOddsTime)

	fmt.Println(BeginOddsTime, e)
}
