package Bll

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/GInstance"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models"
	"github.com/TtMyth123/kit/lotteryKit"
	"github.com/astaxie/beego/orm"
)

func GetAndAddRaceTmpList(SysId int64) (map[string]interface{}, error) {
	GInstance.GetWorkContainer().GetWork(SysId).GetAndAddRaceTmpList()
	return nil, nil
}
func AddRaceIds2RaceTable(SysId int64, Ids string) (map[string]interface{}, error) {
	arrId := lotteryKit.GetStrNum2Arr(Ids)
	intIds := ArrInt2Int64(arrId)
	o := orm.NewOrm()
	for _, Id := range intIds {
		aFaRaceInfo := models.FaRaceInfo{}
		aFaRaceInfo.Id = Id
		aFaRaceInfo.IsJsTmp = 0
		aFaRaceInfo.IsJs = 1
		aFaRaceInfo.SysId = SysId
		aFaRaceInfo.Update(o, "IsJsTmp", "IsJs")
	}
	return nil, nil
}

func DelJsRaceTable(SysId int64, strIds string) (map[string]interface{}, error) {
	arrId := lotteryKit.GetStrNum2Arr(strIds)
	intIds := ArrInt2Int64(arrId)
	o := orm.NewOrm()
	for _, id := range intIds {
		aFaRaceInfo := models.FaRaceInfo{}
		aFaRaceInfo.Id = id
		aFaRaceInfo.IsJsTmp = 1
		aFaRaceInfo.IsJs = 1
		aFaRaceInfo.Update(o, "IsJsTmp", "IsJs")
	}
	return nil, nil
}

func GetAndSaveRaceData(SysId int64, strIds string) (map[string]interface{}, error) {
	arrId := lotteryKit.GetStrNum2Arr(strIds)
	Ids := ArrInt2Int64(arrId)
	GInstance.GetWorkContainer().GetWork(SysId).GetAndSaveRaceData(Ids...)

	return nil, nil
}

func ArrInt2Int64(arr []int) []int64 {
	iLen := len(arr)
	arr64 := make([]int64, iLen)
	for i := 0; i < iLen; i++ {
		arr64[i] = int64(arr[i])
	}

	return arr64
}

func GetRaceDataList(SysId int64, Ids string, isJs bool) (interface{}, error) {
	arrIds := lotteryKit.GetStrNum2Arr(Ids)
	raceIds := ArrInt2Int64(arrIds)
	data, e := GInstance.GetWorkContainer().GetWork(SysId).GetRaceDataList(raceIds, isJs)
	return data, e
}
