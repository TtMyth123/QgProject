package Bll

import (
	"fmt"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/GInstance"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models"
	"github.com/TtMyth123/kit/lotteryKit"
	"github.com/astaxie/beego/orm"
	"strings"
)

func UpdateHcTmpRaceList(SysId int64, Days string) (map[string]interface{}, error) {
	arrDay := strings.Split(Days, ",")
	GInstance.GetWorkContainer().GetWork(SysId).GetAndAddHcRaceTmpList(arrDay...)
	return nil, nil
}

func AddHcRaceIds2RaceTable(SysId int64, Ids string) (map[string]interface{}, error) {
	arrIds := lotteryKit.GetStrNum2Arr(Ids)
	raceIds := ArrInt2Int64(arrIds)
	o := orm.NewOrm()
	for _, Id := range raceIds {
		aFaRaceInfo := models.FaRaceInfo{}
		aFaRaceInfo.Id = Id
		aFaRaceInfo.IsHcTmp = 0
		aFaRaceInfo.IsHc = 1
		aFaRaceInfo.Update(o, "IsHcTmp", "IsHc")
	}
	//GInstance.GetWorkContainer().GetWork(SysId).GetAndSaveRaceData(raceIds...)

	return nil, nil
}

func DelHcRaceTable(SysId int64, RaceInfoTmpIds string) (map[string]interface{}, error) {
	arrRaceInfoTmpId := lotteryKit.GetStrNum2Arr(RaceInfoTmpIds)
	for _, id := range arrRaceInfoTmpId {
		aFaRaceInfo := models.FaRaceInfo{}
		aFaRaceInfo.Id = int64(id)
		aFaRaceInfo.IsHcTmp = 1
		aFaRaceInfo.IsHc = 1
		e := aFaRaceInfo.Update(nil, "IsHcTmp", "IsHc")
		if e != nil {
			fmt.Println()
		}
	}
	return nil, nil
}
