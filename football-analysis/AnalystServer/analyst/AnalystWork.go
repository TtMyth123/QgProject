package analyst

import (
	"fmt"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/footballKit"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models"
	"github.com/TtMyth123/kit/goqueryKit"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"github.com/opesun/goquery"
	"math"
	"strings"
	"sync"
	"time"
)

type AnalystWork struct {
	SysId              int64
	mAnalystHttpClient *AnalystHttpClient
	mExcelFormulaData  *ExcelFormulaDataEx
}

func NewAnalystWork(SysId int64, excelFormulaFile string, basePath1, basePath2, basePath3 string) *AnalystWork {
	work := new(AnalystWork)
	work.SysId = SysId
	work.mAnalystHttpClient = NewAnalystHttpClient(basePath1, basePath2, basePath3)
	work.mExcelFormulaData = NewExcelFormulaDataEx(excelFormulaFile)
	if work.mExcelFormulaData == nil {
		panic("NewExcelFormulaData 出错")
	}
	return work
}

func repairTeamNameOld(oldTeamName string) string {

	i := strings.Index(oldTeamName, `(`)
	teamName := oldTeamName
	if i > 0 {
		teamName = oldTeamName[i:]
	}
	return teamName
}
func repairTeamName(oldTeamName string) string {
	strName := ""
	ii := 0
	for {
		i := strings.Index(oldTeamName, `<`)
		if i > ii {
			strName += oldTeamName[:i]
			oldTeamName = oldTeamName[i:]

			i1 := strings.Index(oldTeamName, `>`)
			if i1 > 0 {
				oldTeamName = oldTeamName[i1+1:]
			}
		} else {
			strName += oldTeamName
			break
		}
	}

	return strName
}

func (this *AnalystWork) GetAndAddRaceTmpList() {
	html, e := this.mAnalystHttpClient.GetBFData()
	if e != nil {
		ttLog.LogDebug(e)
		return
	}
	arrFaRaceInfo, mpFaTeamInfo, e := this.GetHtml2RaceTmpList(html)
	if e != nil {
		ttLog.LogDebug(e)
		return
	}

	o := orm.NewOrm()
	for _, data := range mpFaTeamInfo {
		data.AddUpdate(o)
	}
	for _, data := range arrFaRaceInfo {
		data.AddUpdate(o, "EndCOdds2", "EndEOdds2", "LeagueName", "ATeamCName", "BTeamCName",
			"RaceTime", "AScore", "BScore", "AFirstHalfScore", "BFirstHalfScore", "ATeamId", "BTeamId",
			"RaceInfoId", "SysId", "IsJsTmp", "IsJs")
	}
}

func (this *AnalystWork) GetHtml2RaceTmpList(html string) ([]models.FaRaceInfo, map[int64]models.FaTeamInfo, error) {
	const (
		raceIndexTmp_29 = 29
		raceIndexTmp_02 = 2
		raceIndexTmp_05 = 5
		raceIndexTmp_08 = 8
		raceIndexTmp_00 = 0
		raceIndexTmp_12 = 12
		raceIndexTmp_14 = 14
		raceIndexTmp_15 = 15
		raceIndexTmp_16 = 16
		raceIndexTmp_17 = 17
		raceIndexTmp_37 = 37
		raceIndexTmp_38 = 38
	)
	mapFaTeamInfo := make(map[int64]models.FaTeamInfo)
	arrRace := make([]models.FaRaceInfo, 0)
	arrHtml := strings.Split(html, "\r")
	iLen := len(arrHtml)

	for i := 0; i < iLen; i++ {
		s := strings.Trim(arrHtml[i], " ")
		sLen := len(s)
		A := s[:2]
		if sLen > 3 && A == "A[" {
			strArr := stringKit.GetBetweenStr(s, `"`, `".`)
			arrItems := strings.Split(strArr, `^`)
			EndEOdds2 := strconvEx.StrTry2Float64(arrItems[raceIndexTmp_29], 0)

			EndCOdds2 := footballKit.Goal2GoalCn(EndEOdds2)
			LeagueName := arrItems[raceIndexTmp_02]
			ATeamCName := repairTeamName(arrItems[raceIndexTmp_05])
			BTeamCName := repairTeamName(arrItems[raceIndexTmp_08])

			RaceInfoId := strconvEx.StrTry2Int64(arrItems[raceIndexTmp_00], 0)

			dateItems := strings.Split(arrItems[raceIndexTmp_12], `,`)

			RaceTime := time.Date(
				strconvEx.StrTry2Int(dateItems[0], 0),
				time.Month(strconvEx.StrTry2Int(dateItems[1], 0)+1),
				strconvEx.StrTry2Int(dateItems[2], 0),
				strconvEx.StrTry2Int(dateItems[3], 0),
				strconvEx.StrTry2Int(dateItems[4], 0),
				strconvEx.StrTry2Int(dateItems[5], 0), 0,
				time.Local)
			AScore := strconvEx.StrTry2Int(arrItems[raceIndexTmp_14], 0)
			BScore := strconvEx.StrTry2Int(arrItems[raceIndexTmp_15], 0)
			AFirstHalfScore := strconvEx.StrTry2Int(arrItems[raceIndexTmp_16], 0)
			BFirstHalfScore := strconvEx.StrTry2Int(arrItems[raceIndexTmp_17], 0)

			ATeamId := strconvEx.StrTry2Int64(arrItems[raceIndexTmp_37], 0)
			BTeamId := strconvEx.StrTry2Int64(arrItems[raceIndexTmp_38], 0)

			aFaTeamInfoA := models.FaTeamInfo{
				BaseInfo:  models.BaseInfo{Id: ATeamId},
				TeamEName: ATeamCName,
				TeamCName: ATeamCName,
			}
			mapFaTeamInfo[aFaTeamInfoA.Id] = aFaTeamInfoA

			aFaTeamInfoB := models.FaTeamInfo{
				BaseInfo:  models.BaseInfo{Id: BTeamId},
				TeamEName: BTeamCName,
				TeamCName: BTeamCName,
			}
			mapFaTeamInfo[aFaTeamInfoB.Id] = aFaTeamInfoB

			aFaRaceTable := models.FaRaceInfo{
				IsJsTmp:         1,
				IsJs:            1,
				EndCOdds2:       EndCOdds2,
				EndEOdds2:       EndEOdds2,
				LeagueName:      LeagueName,
				ATeamCName:      ATeamCName,
				BTeamCName:      BTeamCName,
				RaceTime:        RaceTime,
				AScore:          AScore,
				BScore:          BScore,
				AFirstHalfScore: AFirstHalfScore,
				BFirstHalfScore: BFirstHalfScore,
				ATeamId:         ATeamId,
				BTeamId:         BTeamId,
			}
			aFaRaceTable.RaceInfoId = RaceInfoId
			aFaRaceTable.SysId = this.SysId

			arrRace = append(arrRace, aFaRaceTable)
		}
	}

	return arrRace, mapFaTeamInfo, nil
}

func (this *AnalystWork) GetAndAddHcRaceTmpList(strDays ...string) {
	o := orm.NewOrm()
	for _, day := range strDays {
		tTime, e := timeKit.GetDateForTime(day)
		if e != nil {
			continue
		}
		//day1 := tTime.Format("2006-1-2")
		//html, e := this.mAnalystHttpClient.MatchByCountryHtml(day1)
		//if e != nil {
		//	continue
		//}
		//mpEnRaceInfo := this.GetHtml2EnRaceInfoList(html)

		day2 := tTime.Format("20060102")

		html, e := this.mAnalystHttpClient.GetHcRaceInfoHtml(day2)
		arrFaRaceInfo, mpFaTeamInfo, e := this.GetHtml2HcRaceTmpList(html, tTime, nil)
		if e != nil {
			continue
		}

		for _, data := range mpFaTeamInfo {
			data.AddUpdate(o)
		}
		for _, data := range arrFaRaceInfo {
			data.AddUpdate(o)
		}
	}
}
func (this *AnalystWork) GetHtml2EnRaceInfoList(html string) map[int64]EnRaceInfo {
	mpRace := make(map[int64]EnRaceInfo)

	arr := strings.Split(html, `; A`)
	for _, item := range arr {
		s := stringKit.GetBetweenStr(item, `=[`, `];`)
		if len(s) > 0 {
			dataItem := strings.Split(s, `,`)
			RaceInfoId := strconvEx.StrTry2Int64(dataItem[0], 0)
			ATeamID := strconvEx.StrTry2Int64(dataItem[2], 0)
			BTeamID := strconvEx.StrTry2Int64(dataItem[3], 0)
			if RaceInfoId > 0 {
				mpRace[RaceInfoId] = EnRaceInfo{
					RaceInfoId: RaceInfoId,
					ATeamId:    ATeamID,
					BTeamId:    BTeamID,
				}
			}
		}
	}

	return mpRace
}

func (this *AnalystWork) GetHtml2HcRaceTmpList(html string, time2 time.Time, mpEnRaceInfo map[int64]EnRaceInfo) ([]models.FaRaceInfo, map[int64]models.FaTeamInfo, error) {
	if mpEnRaceInfo == nil {
		mpEnRaceInfo = make(map[int64]EnRaceInfo)
	}
	mapFaTeamInfo := make(map[int64]models.FaTeamInfo)
	arrRace := make([]models.FaRaceInfo, 0)
	dom, e := goquery.ParseString(html)
	if e != nil {
		return arrRace, mapFaTeamInfo, e
	}

	tableNodes := dom.Find("#table_live")
	arrStrTable := goqueryKit.GetTableNodes2Arr(tableNodes)
	for i, arrTd := range arrStrTable {
		if i < 1 {
			continue
		}
		if len(arrTd) < 9 {
			continue
		}
		LeagueName := arrTd[0].Text
		strTimes := strings.Split(arrTd[1].Text, "日")
		day := strconvEx.StrTry2Int(strTimes[0], 0)
		arrHM := strings.Split(strTimes[1], ":")
		hour := strconvEx.StrTry2Int(arrHM[0], 0)
		min := strconvEx.StrTry2Int(arrHM[1], 0)

		RaceTime := time.Date(time2.Year(), time2.Month(), day, hour, min, 0, 0, time.Local)

		//analysis(1962981)
		strRaceId := stringKit.GetBetweenStr(arrTd[9].Node.Child[1].Attr[1].Val, "(", ")")
		RaceInfoId := strconvEx.StrTry2Int64(strRaceId, 0)
		ATeamCName := arrTd[3].Text
		BTeamCName := arrTd[5].Text

		ARanking := stringKit.GetBetweenStr(arrTd[3].Text, "[", "]")
		BRanking := stringKit.GetBetweenStr(arrTd[5].Text, "[", "]")
		if ARanking != "" {
			ATeamCName = stringKit.GetLaterStr(ATeamCName, "]")
		}
		if BRanking != "" {
			BTeamCName = stringKit.GetBeforeStr(BTeamCName, "[")
		}

		AFirstHalfScore := 0
		BFirstHalfScore := 0
		firstHalfScoreItems := strings.Split(arrTd[6].Text, "-")
		if len(firstHalfScoreItems) == 2 {
			AFirstHalfScore = strconvEx.StrTry2Int(firstHalfScoreItems[0], 0)
			BFirstHalfScore = strconvEx.StrTry2Int(firstHalfScoreItems[1], 0)
		}

		AScore := 0
		BScore := 0
		scoreItems := strings.Split(arrTd[4].Text, "-")
		if len(firstHalfScoreItems) == 2 {
			AScore = strconvEx.StrTry2Int(scoreItems[0], 0)
			BScore = strconvEx.StrTry2Int(scoreItems[1], 0)
		}

		EndCOdds2 := strings.Replace(arrTd[7].Text, "*", "", -1)
		EndOdds2 := footballKit.GoalCn2Goal(EndCOdds2)

		aFaRaceInfo := models.FaRaceInfo{}
		aFaRaceInfo.ATeamCName = ATeamCName
		aFaRaceInfo.ATeamEName = ATeamCName
		aFaRaceInfo.BTeamCName = BTeamCName
		aFaRaceInfo.BTeamEName = BTeamCName
		aFaRaceInfo.RaceTime = RaceTime
		aFaRaceInfo.RaceInfoId = RaceInfoId
		aFaRaceInfo.ARanking = ARanking
		aFaRaceInfo.BRanking = BRanking
		aFaRaceInfo.LeagueName = LeagueName

		aFaRaceInfo.AFirstHalfScore = AFirstHalfScore
		aFaRaceInfo.BFirstHalfScore = BFirstHalfScore
		aFaRaceInfo.AScore = AScore
		aFaRaceInfo.BScore = BScore
		aFaRaceInfo.EndCOdds2 = EndCOdds2
		aFaRaceInfo.EndEOdds2 = EndOdds2
		aFaRaceInfo.SysId = this.SysId

		if aEnRaceInfo, ok := mpEnRaceInfo[RaceInfoId]; ok {

			aFaTeamInfoA := models.FaTeamInfo{
				BaseInfo:  models.BaseInfo{Id: aEnRaceInfo.ATeamId},
				TeamEName: ATeamCName,
				TeamCName: ATeamCName,
			}
			mapFaTeamInfo[aFaTeamInfoA.Id] = aFaTeamInfoA

			aFaTeamInfoB := models.FaTeamInfo{
				BaseInfo:  models.BaseInfo{Id: aEnRaceInfo.BTeamId},
				TeamEName: BTeamCName,
				TeamCName: BTeamCName,
			}
			mapFaTeamInfo[aFaTeamInfoB.Id] = aFaTeamInfoB
			aFaRaceInfo.ATeamId = aFaTeamInfoA.Id
			aFaRaceInfo.BTeamId = aFaTeamInfoB.Id
		}

		arrRace = append(arrRace, aFaRaceInfo)

	}

	return arrRace, mapFaTeamInfo, nil
}

func (this *AnalystWork) GetAndSaveRaceData(Ids ...int64) {
	o := orm.NewOrm()
	for _, id := range Ids {
		aRaceInfo := models.FaRaceInfo{}
		aRaceInfo.Id = id
		if e := aRaceInfo.Read(o); e != nil {
			continue
		}

		aABHistoryRaceData, e1 := this.GetABHistoryRaceData(aRaceInfo)
		if e1 != nil {
			ttLog.LogDebug("GetABHistoryRaceData:", e1)
			continue
		}
		this.saveABHistoryRaceData(aABHistoryRaceData)

		//----------------------------------------
		arrAsiaOddsInfo, e1 := this.GetAsiaOddsInfoData(aRaceInfo)
		if e1 != nil {
			ttLog.LogDebug("GetAsiaOddsInfoData:", e1)
		}
		e1 = models.MultiSaveAsiaOddsInfo(o, arrAsiaOddsInfo)
		if e1 != nil {
			ttLog.LogDebug("MultiSaveAsiaOddsInfo:", e1)
		}

		//----------------------------------------
		arrEuropeOddsInfo, e1 := this.GetEuropeOddsInfoData(aRaceInfo)
		if e1 != nil {
			ttLog.LogDebug("GetEuropeOddsInfoData:", e1)
		}
		e1 = models.MultiSaveEuropeOddsInfo(o, arrEuropeOddsInfo)
		if e1 != nil {
			ttLog.LogDebug("MultiSaveEuropeOddsInfo:", e1)
		}

		//----------------------------------------
		arrGSOddsInfo, e1 := this.GetGSOddsInfoData(aRaceInfo)
		if e1 != nil {
			ttLog.LogDebug("GetGSOddsInfoData:", e1)
		}
		e1 = models.MultiSaveGSOddsInfo(o, arrGSOddsInfo)
		if e1 != nil {
			ttLog.LogDebug("MultiSaveGSOddsInfo:", e1)
		}

		this.mExcelFormulaData.LoadAsiaOddsData(arrAsiaOddsInfo)
		this.mExcelFormulaData.LoadHistoryRaceInfo(aABHistoryRaceData.HistoryRaceInfoExtA,
			aABHistoryRaceData.HistoryRaceInfoExtB, aRaceInfo.ATeamCName, aRaceInfo.BTeamCName)
		this.mExcelFormulaData.LoadEuropeOddsData(arrEuropeOddsInfo)
		data := this.mExcelFormulaData.GetExcelFormulaData()
		//this.mExcelFormulaData.xlFile.SaveAs(`d:\\tmp\a.xlsx`)
		data.SysId = this.SysId
		data.MainId = aRaceInfo.Id
		data.RaceInfoId = aRaceInfo.RaceInfoId
		data.DataType = models.ExcelFData_DataType_Whole
		data.AddUpdate(o)
	}
}
func (this *AnalystWork) CaleSaveExcelData(o orm.Ormer, aRaceInfo models.FaRaceInfo, lstAsiaOddsInfo []models.FaAsiaOddsInfo,
	lstEuropeOddsInfo []models.FaEuropeOddsInfo, arrAHistoryRace, arrBHistoryRace []models.FaHistoryRaceInfoExt,
	ATeamName, BTeamName string) {
	this.mExcelFormulaData.LoadAsiaOddsData(lstAsiaOddsInfo)
	this.mExcelFormulaData.LoadHistoryRaceInfo(arrAHistoryRace, arrBHistoryRace, ATeamName, BTeamName)
	this.mExcelFormulaData.LoadEuropeOddsData(lstEuropeOddsInfo)
	data := this.mExcelFormulaData.GetExcelFormulaData()

	data.SysId = this.SysId
	data.MainId = aRaceInfo.Id
	data.RaceInfoId = aRaceInfo.RaceInfoId
	data.DataType = models.ExcelFData_DataType_Whole
	data.AddUpdate(o)

	t := time.Now()
	if aRaceInfo.RaceTime.Sub(t) < 0 {
		t = aRaceInfo.RaceTime
	}

	oneHour := time.Hour
	lstEuropeOddsInfoOneHour := make([]models.FaEuropeOddsInfo, 0)
	lstEuropeOddsInfoAuthority := make([]models.FaEuropeOddsInfo, 0)
	for _, EuropeOddsInfo := range lstEuropeOddsInfo {
		if EuropeOddsInfo.EndOddsTime.Add(oneHour).Sub(t) >= 0 {
			lstEuropeOddsInfoOneHour = append(lstEuropeOddsInfoOneHour, EuropeOddsInfo)
		}
		if EuropeOddsInfo.TypeE == 1 {
			lstEuropeOddsInfoAuthority = append(lstEuropeOddsInfoAuthority, EuropeOddsInfo)
		}
	}

	this.mExcelFormulaData.LoadEuropeOddsData(lstEuropeOddsInfoOneHour)
	dataOneHour := this.mExcelFormulaData.GetExcelFormulaData()
	dataOneHour.SysId = this.SysId
	dataOneHour.MainId = aRaceInfo.Id
	dataOneHour.RaceInfoId = aRaceInfo.RaceInfoId
	dataOneHour.DataType = models.ExcelFData_DataType_One
	dataOneHour.AddUpdate(o)

	this.mExcelFormulaData.LoadEuropeOddsData(lstEuropeOddsInfoAuthority)
	dataAuthority := this.mExcelFormulaData.GetExcelFormulaData()
	dataAuthority.SysId = this.SysId
	dataAuthority.MainId = aRaceInfo.Id
	dataAuthority.RaceInfoId = aRaceInfo.RaceInfoId
	dataAuthority.DataType = models.ExcelFData_DataType_Auth
	dataAuthority.AddUpdate(o)

}

func (this *AnalystWork) GetRaceDataList(arrIds []int64, isJs bool) (interface{}, error) {
	mpData := make(map[int64]interface{})
	wg := sync.WaitGroup{}
	for _, id := range arrIds {
		wg.Add(1)
		o := orm.NewOrm()
		aFaRaceInfo := models.FaRaceInfo{}
		aFaRaceInfo.Id = id
		e := aFaRaceInfo.Read(o)
		if e != nil {
			continue
		}

		if aFaRaceInfo.IsGet == 0 {
			this.GetAndSaveRaceData(aFaRaceInfo.Id)
		}

		aFaRaceSetup := models.FaRaceSetup{
			BaseInfo: models.BaseInfo{SysId: this.SysId}, RaceInfoId: aFaRaceInfo.RaceInfoId,
		}
		aFaRaceSetup.ReadEx(o)

		HistoryFightInfoList, e := this.GetHistoryFightInfoList(o, id)
		if e != nil {
			ttLog.LogError(e)
		}

		AFutureRaceInfoList, e := this.GetFutureRaceInfoList(o, id, models.FutureType_A)
		if e != nil {
			ttLog.LogError(e)
		}
		BFutureRaceInfoList, e := this.GetFutureRaceInfoList(o, id, models.FutureType_B)
		if e != nil {
			ttLog.LogError(e)
		}

		AIntegralRankingInfo, e := this.GetIntegralRankingInfo(o, id, models.IntegralType_A)
		if e != nil {
			ttLog.LogError(e)
		}
		BIntegralRankingInfo, e := this.GetIntegralRankingInfo(o, id, models.IntegralType_B)
		if e != nil {
			ttLog.LogError(e)
		}
		ExcelFData, e := this.GetExcelFData(o, id, models.ExcelFData_DataType_Whole)
		ExcelFDataOneHour, e := this.GetExcelFData(o, id, models.ExcelFData_DataType_One)
		ExcelFDataAuthority, e := this.GetExcelFData(o, id, models.ExcelFData_DataType_Auth)

		EuropeOddCalculateOneHour, e := this.GetEuropeOddCalculate(o, id, models.EuropeOddCalculateType_OneHour)
		EuropeOddCalculateAuthority, e := this.GetEuropeOddCalculate(o, id, models.EuropeOddCalculateType_Authority)
		EuropeOddCalculateAvg, e := this.GetEuropeOddCalculate(o, id, models.EuropeOddCalculateType_Avg)

		mpRaceData := make(map[string]interface{})

		mpRaceData["RaceInfo"] = aFaRaceInfo
		mpRaceData["RaceSetup"] = aFaRaceSetup
		mpRaceData["HistoryFightInfoList"] = HistoryFightInfoList
		mpRaceData["AFutureRaceInfoList"] = AFutureRaceInfoList
		mpRaceData["BFutureRaceInfoList"] = BFutureRaceInfoList

		mpRaceData["AIntegralRankingInfo"] = AIntegralRankingInfo
		mpRaceData["BIntegralRankingInfo"] = BIntegralRankingInfo

		mpRaceData["ExcelFData"] = ExcelFData
		mpRaceData["ExcelFDataOneHour"] = ExcelFDataOneHour
		mpRaceData["ExcelFDataAuthority"] = ExcelFDataAuthority

		mpRaceData["EuropeOddCalculateOneHour"] = EuropeOddCalculateOneHour
		mpRaceData["EuropeOddCalculateAuthority"] = EuropeOddCalculateAuthority
		mpRaceData["EuropeOddCalculateAvg"] = EuropeOddCalculateAvg

		mpData[id] = mpRaceData
		wg.Done()
	}

	return mpData, nil
}

type ABHistoryRaceData struct {
	RaceInfo                 models.FaRaceInfo
	HistoryRaceInfoExtA      []models.FaHistoryRaceInfoExt
	HistoryRaceInfoExtB      []models.FaHistoryRaceInfoExt
	HistoryRaceInfoExtH      []models.FaHistoryRaceInfoExt
	InjuryInfoList           []models.FaInjury
	FutureRaceInfoDataA      []models.FaFutureRaceInfo
	FutureRaceInfoDataB      []models.FaFutureRaceInfo
	IntegralRankingInfoDataA models.FaIntegralRankingInfo
	IntegralRankingInfoDataB models.FaIntegralRankingInfo
}

func (this *AnalystWork) GetABHistoryRaceData(aRaceInfo models.FaRaceInfo) (ABHistoryRaceData, error) {
	aABHistoryRaceData := ABHistoryRaceData{}
	strHtml, e := this.mAnalystHttpClient.GetABHistoryRaceDataHtml(aRaceInfo.RaceInfoId)
	if e != nil {
		return aABHistoryRaceData, e
	}
	HomeTeam := stringKit.GetBetweenStr(strHtml, "hometeam = ", ";")
	GuestTeam := stringKit.GetBetweenStr(strHtml, "guestteam = ", ";")
	HomeTeam = strings.Trim(HomeTeam, `"`)
	GuestTeam = strings.Trim(GuestTeam, `"`)

	aRaceInfo.HomeTeam = HomeTeam
	aRaceInfo.GuestTeam = GuestTeam

	//h_data=[['16-01-23',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>(中)',2696,'<span title="韩国女足  排名:17">韩国女足</span>',2,0,'2-0','-0.75',1,1,-1,1206956],['16-01-21',428,'四国赛','#cc9900',1933,'<span title="中国女足  排名:14">中国女足</span>',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>',0,0,'0-0','1.5',0,1,-1,1207004],['15-12-21',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>(中)',7812,'<span title="特立尼达和多巴哥女足  排名:48">特立尼达</span>',2,1,'1-0','2.5',1,-1,1,1202511],['15-12-17',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>',7812,'<span title="特立尼达和多巴哥女足  排名:48">特立尼达</span>',3,0,'1-0','2.25',1,1,1,1201853],['15-12-14',1366,'国际友谊','#4666bb',4883,'<span title="巴西女足  排名:6">巴西女足</span>',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>',6,0,'4-0','2',-1,-1,1,1201882],['15-12-10',1366,'国际友谊','#4666bb',2337,'<span title="加拿大女足  排名:11">加拿大女</span>(中)',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>',3,0,'3-0','0.25',-1,-1,1,1201147],['15-07-25',487,'泛美女足','#b00900',2337,'<span title="加拿大女足  排名:9">加拿大女</span>',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',1,2,'0-2','0.25',1,1,1,1156312],['15-07-23',487,'泛美女足','#b00900',4883,'<span title="巴西女足  排名:8">巴西女足</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',4,2,'2-1','1.25',-1,-1,1,1153119],['15-07-19',487,'泛美女足','#b00900',7812,'<span title="特立尼达和多巴哥女足  排名:45">特立尼达</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',1,3,'0-2','-1.25',1,1,1,1143811],['15-07-15',487,'泛美女足','#b00900',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',5937,'<span title="阿根廷女足  排名:36">阿根廷女</span>',3,1,'1-0','0.5',1,1,1,1143809],['15-07-11',487,'泛美女足','#b00900',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',12344,'<span title="哥伦比亚女足  排名:28">哥伦比亚</span>',0,1,'0-0','0',-1,-1,-1,1143807],['15-06-18',388,'女世界杯','#696900',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',2361,'<span title="法国女足  排名:3">法国女足</span>',0,5,'0-4','-1.75',-1,-1,1,1127589],['15-06-14',388,'女世界杯','#696900',2359,'<span title="英格兰女足  排名:6">英格兰女</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',2,1,'0-0','1.25',-1,1,1,1127588],['15-06-10',388,'女世界杯','#696900',12344,'<span title="哥伦比亚女足  排名:28">哥伦比亚</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',1,1,'0-1','-0.25',0,-1,-1,1127586],['15-05-29',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',7829,'<span title="哥斯达黎加女足  排名:37">哥斯达黎</span>',3,0,'3-0','1.25',1,1,1,1128820],['15-05-26',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',7829,'<span title="哥斯达黎加女足  排名:37">哥斯达黎</span>',2,1,'1-1','',1,-2,1,1128589],['15-05-18',1366,'国际友谊','#4666bb',3862,'<span title="美国女足  排名:2">美国女足</span>',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',5,1,'1-1','3.5',-1,-1,1,1127927],['15-03-11',1366,'国际友谊','#4666bb',2362,'<span class=hp>1</span><span title="意大利女足  排名:14">意大利女</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',2,3,'0-1','0.75',1,1,1,1106456],['15-03-09',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',5564,'<span title="比利时女足  排名:26">比利时女</span>',0,0,'0-0','0.25',0,-1,-1,1105992],['15-03-06',1366,'国际友谊','#4666bb',4519,'<span title="捷克女足  排名:30">捷克女足</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span><span class=hp>1</span>',0,1,'0-0','-0.5',1,1,-1,1102793],['15-03-04',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',9935,'<span title="南非女足  排名:60">南非女足</span>',2,0,'0-0','1.25',1,1,-1,1103069],['15-02-08',1366,'国际友谊','#4666bb',5247,'<span class=hp>1</span><span title="墨西哥女足  排名:25">墨西哥女</span>',5938,'<span title="厄瓜多尔女足  排名:46">厄瓜多尔</span>',2,0,'0-0','',1,-2,-1,1092095],['15-02-06',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',5938,'<span title="厄瓜多尔女足  排名:46">厄瓜多尔</span>',1,0,'1-0','',1,-2,-1,1090409],['15-01-15',184,'女足四国','#FF2F73',2696,'<span title="韩国女足  排名:">韩国女足</span>(中)',5247,'<span title="墨西哥女足  排名:">墨西哥女</span>',2,1,'1-0','0.75',-1,-1,1,1080891],['15-01-13',184,'女足四国','#FF2F73',2337,'<span title="加拿大女足  排名:9">加拿大女</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',2,1,'0-0','1.25',-1,1,1,1080244],['15-01-11',1366,'国际友谊','#4666bb',1933,'<span title="中国女足  排名:13">中国女足</span>',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',0,0,'0-0','1.5',0,1,-1,1079683],['14-11-28',1371,'女美加杯','#888566',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',12344,'<span title="哥伦比亚女足  排名:31">哥伦比亚</span>',2,0,'0-0','',1,-2,-1,1071281],['14-11-26',1371,'女美加杯','#888566',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',7829,'<span title="哥斯达黎加女足  排名:40">哥斯达黎</span>',1,0,'0-0','',1,-2,-1,1071027],['14-11-22',1371,'女美加杯','#888566',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',16109,'<span title="海地女足  排名:56">海地女足</span><span class=hp>1</span>',1,0,'0-0','',1,-2,-1,1070698],['14-11-20',1371,'女美加杯','#888566',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',7812,'<span title="特立尼达和多巴哥女足  排名:125">特立尼达</span>',6,0,'1-0','',1,-2,1,1070696]];
	//主队的历史对战 有主队的
	h_data := stringKit.GetBetweenStr(strHtml, "h_data=[[", "]];")
	if h_data == "" {
		h_data = stringKit.GetBetweenStr(strHtml, "h_data = [[", "]];")
	}

	//客队的历史对战 有客队的
	a_data := stringKit.GetBetweenStr(strHtml, "a_data=[[", "]];")
	if a_data == "" {
		a_data = stringKit.GetBetweenStr(strHtml, "a_data = [[", "]];")
	}

	//主客队的历史对战 有主客队的
	v_data := stringKit.GetBetweenStr(strHtml, "v_data=[[", "]];")
	if v_data == "" {
		v_data = stringKit.GetBetweenStr(strHtml, "v_data = [[", "]];")
	}

	//亚赔数据
	//Vs_hOdds=[[940737,24,'1.01','-0.75','0.83','1.09','0','0.75','3','1.5'],[941285,24,'0.87','2.25','0.97','0.87','2.25','0.97','3.25','3.25'],[941539,24,'0.84','0','0.96','0.78','0','1.02','2.5','2.5'],[981800,3,'0.91','1.25','0.85','0.50','1.5','1.35','3.25','3.25'],[981800,8,'0.88','0.75','0.93','0.53','1.5','1.43','2.25','3.25'],[981800,12,'0.65','1.25','1.30','0.65','1.25','1.30','3.25','3.25'],[982222,1,'0.80','2.75','0.90','0.90','2.75','0.80','3.75','3.75'],[982222,3,'0.75','2.5','0.95','0.85','2.5','0.97','3.75','3.75'],[982222,8,'0.85','2.5','0.95','0.90','2.75','0.90','3.75','3.75'],[982222,12,'0.90','2.75','1.05','0.90','2.75','1.05','3.75','3.75'],[982222,24,'0.70','2.75','1.16','0.86','2.75','0.98','3.75','3.75'],[982225,1,'1.20','-4.5','0.50','0.75','-4.25','0.95','5.25','5.25'],[982225,3,'0.90','-4.25','0.80','0.81','-4.25','1.01','5.25','5.25'],[982225,8,'1.00','-3.5','0.80','0.83','-4.25','0.98','4.5','5.25'],[982225,12,'1.00','-4.25','0.95','0.85','-4.5','1.10','5.25','5.25'],[982225,24,'1.20','-4.5','0.67','0.80','-4.5','1.04','5.25','5.25'],[982226,3,'0.75','-1.75','0.95','1.00','-2','0.70','3','3.25'],[982226,8,'0.90','-2','0.90','1.35','-1.75','0.58','2.75','3'],[982226,12,'0.85','-1.5','1.00','1.25','-2','0.60','2.75','3'],[982226,24,'1.21','-1.5','0.58','1.16','-1.75','0.62','2.75','3'],[983765,1,'0.60','0.25','1.10','0.70','0.25','1.00','2.5','2.25'],[983765,3,'0.90','0.25','0.80','0.61','0.25','1.17','2.5','2.25'],[983765,8,'0.95','0.5','0.85','0.70','0.25','1.10','2.5','2.25'],[983765,12,'0.95','0.25','0.95','0.75','0.25','1.15','2.5','2.5'],[983765,24,'0.86','0.25','0.90','0.74','0.25','1.02','2.5','2'],[1054205,1,'0.70','4.25','1.00','0.72','4.25','0.98','4.5','5.25'],[1054205,3,'0.85','4.25','0.91','0.72','4.25','1.11','5.25','5.25'],[1054205,8,'0.90','3.5','0.90','0.75','4.25','1.05','4.5','5.25'],[1054205,12,'0.90','4.25','1.05','0.88','4.25','1.07','5.25','5.25'],[1054205,24,'0.74','4.25','1.11','0.86','4.25','0.98','5.25','5.25'],[1054207,3,'0.61','2.5','1.09','0.82','2.5','0.94','3.75','3.75'],[1054207,8,'0.90','1.5','0.90','0.88','2.5','0.93','3','3.75'],[1054207,12,'0.70','2.5','1.25','0.70','2.5','1.25','3.75','3.75'],[1054207,24,'0.64','2.5','1.13','0.92','2.5','0.88','3.75','3.75'],[1054623,3,'0.77','2.25','0.99','0.60','2.25','1.28','3.75','3.75'],[1054623,24,'0.60','2.25','1.25','0.60','2.25','1.31','3.75','3.75'],[1060207,3,'0.80','0.75','0.96','0.94','1','0.88','3','2.75'],[1060207,8,'0.83','0.75','0.98','0.73','0.75','1.08','2.5','2.75'],[1060207,12,'0.82','0.75','1.13','0.83','0.75','1.12','2.75','2.75'],[1060207,24,'0.88','0.75','0.96','0.74','0.75','1.11','2.75','2.75'],[1060994,1,'0.75','-4.75','0.95','0.75','-4.75','0.95','4.5','5.75'],[1060994,3,'0.95','-4.75','0.75','0.78','-4.75','0.98','5.75','5.75'],[1060994,8,'1.05','-4.75','0.75','0.78','-4.75','1.03','5.75','5.75'],[1060994,12,'0.88','-4.75','1.02','0.85','-4.75','1.05','5.75','5.75'],[1060994,24,'0.82','-4.75','0.98','0.94','-4.75','0.86','5.75','5.75'],[1062705,1,'1.02','3.75','0.68','0.98','3.5','0.72','4.5','4.25'],[1062705,3,'0.70','3.75','1.06','1.00','3.25','0.82','4.75','4.25'],[1062705,8,'0.80','3.75','1.00','1.00','3.5','0.80','4.5','4.25'],[1062705,12,'1.10','3.5','0.80','1.10','3.5','0.80','4.5','4.25'],[1062705,24,'0.84','3.75','1.00','0.98','3.25','0.82','4.75','4.25'],[1070696,8,'0.83','2.25','0.98','0.83','2.25','0.98','3.5','4'],[1070696,24,'1.00','2.25','0.84','1.00','2.25','0.84','4','4'],[1070698,8,'0.90','3.5','0.90','0.70','3.75','1.10','4.5','5.25'],[1070698,24,'0.54','3.75','1.42','0.54','3.75','1.42','5.25','5.25'],[1071027,8,'0.85','1.25','0.95','0.85','1.25','0.95','3.25','3.25'],[1071027,24,'0.89','1.25','0.95','0.89','1.25','0.95','3.25','3.25'],[1071281,8,'0.85','0.25','0.95','0.85','0.25','0.95','2.5','2.5'],[1071281,24,'0.85','0.25','0.99','0.85','0.25','0.99','2.5','2.5'],[1079683,3,'0.85','1.5','0.91','0.94','1.5','0.88','2.75','2.25'],[1079683,8,'1.00','1.5','0.80','0.78','1.25','1.03','2.75','2.5'],[1079683,12,'0.90','1','1.05','0.90','1','1.05','2.5','2.5'],[1079683,24,'0.92','1.5','0.92','0.80','1.25','1.04','2.75','2.75'],[1080244,3,'0.97','1.25','0.85','0.88','1.25','0.94','2.5','2.5'],[1080244,8,'0.98','1.25','0.83','0.93','1.25','0.88','2.5','2.5'],[1080244,12,'1.07','1.25','0.88','0.97','1.25','0.98','2.5','2.5'],[1080244,24,'0.99','1.25','0.85','0.96','1.25','0.88','2.5','2.5'],[1080891,3,'0.97','0.75','0.85','0.80','0.75','1.02','2.5','2.75'],[1080891,8,'0.98','0.75','0.83','0.78','0.75','1.03','2.5','2.5'],[1080891,12,'0.87','0.5','1.08','0.82','0.75','1.13','2.5','2.5'],[1080891,24,'1.02','0.75','0.82','0.78','0.75','1.06','2.5','2.5'],[1102793,3,'0.80','-0.5','0.96','0.88','-0.25','0.94','2.75','2.5'],[1102793,8,'0.63','-0.5','1.25','0.93','-0.25','0.88','2.75','2.5'],[1102793,12,'0.73','-0.25','1.22','1.02','-0.25','0.93','2.5','2.5'],[1102793,24,'0.80','-0.25','1.04','0.96','-0.25','0.88','2.5','2.5'],[1103069,3,'0.96','1.25','0.80','1.13','1','0.70','2.75','2.5'],[1103069,8,'0.98','1.25','0.83','1.08','1','0.73','2.75','2.5'],[1103069,12,'0.80','0.75','1.10','0.80','0.75','1.10','2.5','2.5'],[1103069,24,'0.98','1.25','0.82','1.06','1','0.78','2.75','2.5'],[1105992,3,'0.85','0.25','0.91','1.07','0','0.75','2.5','2.5'],[1105992,8,'0.88','0.25','0.93','1.08','0','0.73','2.5','2.5'],[1105992,12,'1.35','0','0.60','0.75','-0.25','1.20','2.5','2.5'],[1105992,24,'0.88','0.25','0.92','0.78','-0.25','1.06','2.5','2.5'],[1106456,3,'0.96','0.75','0.80','0.96','0.75','0.86','2.5','2.5'],[1106456,8,'0.95','0.75','0.85','0.90','0.75','0.90','2.75','2.5'],[1106456,12,'1.15','0.75','0.80','0.90','0.75','1.05','2.5','2.5'],[1106456,24,'1.08','0.75','0.76','0.88','0.75','0.96','2.5','2.5'],[1125469,3,'0.77','1','0.99','0.72','1.25','1.11','3.25','3.25'],[1125469,8,'0.95','0.75','0.85','0.83','1.25','0.98','3.25','3'],[1125469,12,'0.97','1','0.98','0.77','1','1.18','3','3'],[1125469,24,'0.96','1.25','0.88','0.86','1.25','0.98','3','3'],[1125472,3,'0.61','4.25','1.17','0.35','4.25','1.88','5.25','5.25'],[1125472,8,'0.48','4.25','1.60','0.35','4.25','2.10','5.25','5.25'],[1125472,24,'0.64','4.25','1.25','0.35','4.25','1.96','5.25','5.25'],[1125473,3,'0.70','4.25','1.00','0.65','4.25','1.20','5.25','5.25'],[1125473,8,'0.70','4.25','1.10','0.63','4.25','1.25','5.25','5.25'],[1125473,12,'0.45','4.25','1.50','0.62','4.25','1.33','5.25','5.25'],[1125473,24,'0.66','4.25','1.11','0.68','4.5','1.19','5.25','5.5'],[1126547,3,'0.80','-0.25','0.96','0.80','0','1.02','3','3'],[1126547,8,'0.83','-0.25','0.98','0.80','0','1.00','3','3'],[1126547,12,'0.75','-0.25','1.20','0.88','0','1.07','3','3'],[1126547,24,'0.75','-0.25','1.09','0.85','0','0.99','3','3'],[1126867,3,'0.77','-0.25','0.99','0.70','0','1.13','2.75','2.5'],[1126867,8,'0.88','0','0.93','0.73','0','1.08','2.75','2.5'],[1126867,12,'1.25','0','0.70','0.90','0','1.05','2.75','2.75'],[1126867,24,'0.72','-0.25','1.08','0.72','0','1.13','2.75','2.75'],[1127586,1,'0.84','-0.5','0.96','1.00','0','0.80','2.25','2.25'],[1127586,3,'1.04','-0.25','0.78','1.04','0','0.78','2.25','2.25'],[1127586,8,'1.08','-0.25','0.73','1.05','0','0.75','2.5','2.25'],[1127586,12,'0.88','-0.5','1.07','1.08','0','0.87','2.25','2.25'],[1127586,24,'1.11','-0.25','0.74','0.76','-0.25','1.08','2.5','2.25'],[1127588,1,'0.94','1.25','0.86','1.10','1','0.70','2.5','2.5'],[1127588,3,'0.99','1.25','0.83','0.84','0.75','0.98','2.5','2.5'],[1127588,8,'0.98','1.25','0.83','0.83','0.75','0.98','2.5','2.5'],[1127588,12,'1.00','1.25','0.95','0.87','0.75','1.08','2.5','2.25'],[1127588,24,'1.04','1.25','0.80','0.82','0.75','1.02','2.5','2.5'],[1127589,1,'0.94','-1.75','0.86','0.60','-1.75','1.20','2.5','2.75'],[1127589,3,'0.85','-1.75','0.97','0.81','-1.5','1.01','2.75','2.75'],[1127589,8,'0.90','-1.75','0.90','0.90','-1.5','0.90','2.75','2.75'],[1127589,12,'0.85','-1.75','1.10','0.90','-1.5','1.05','2.75','2.75'],[1127589,24,'0.98','-1.75','0.86','0.84','-1.5','1.00','2.75','2.75'],[1127927,3,'0.75','3.5','1.01','0.98','3.25','0.84','4.5','4.25'],[1127927,8,'0.83','3.5','0.98','1.05','3.25','0.75','4.5','4.5'],[1127927,12,'0.90','3.5','1.05','1.00','3.25','0.90','4.5','4.25'],[1127927,24,'0.70','3.5','1.11','1.19','3.25','0.68','4.5','4.5'],[1128070,3,'0.77','3.25','0.99','1.02','3.25','0.80','4.25','4.25'],[1128070,8,'0.75','3.25','1.05','0.90','3.25','0.90','4.25','4.25'],[1128070,12,'0.75','3.25','1.20','1.05','3.25','0.90','4.25','4.25'],[1128070,24,'0.82','3.25','1.02','0.98','3.25','0.86','4.25','4.25'],[1128820,3,'0.85','1.25','0.91','0.91','1.25','0.91','3','3'],[1128820,8,'0.88','1.25','0.93','0.83','1.25','0.98','3','3'],[1128820,12,'1.02','1.25','0.93','1.02','1.25','0.93','3','3'],[1128820,24,'0.87','1.25','0.93','0.95','1.25','0.89','3','3'],[1143807,3,'0.80','0','0.96','0.80','0','1.02','2.5','2.5'],[1143807,8,'0.80','0','1.00','0.80','0','1.00','2.25','2.25'],[1143807,12,'0.85','0','1.05','0.85','0','1.05','2','2'],[1143807,24,'0.86','0','0.94','0.84','0','1.00','2.5','2.5'],[1143809,3,'0.85','0.5','0.91','1.07','1.25','0.75','2.5','2.75'],[1143809,8,'0.88','0.5','0.93','0.58','1','1.35','2.5','2.5'],[1143809,12,'0.70','1','1.10','0.70','1','1.10','2.5','2.5'],[1143809,24,'0.72','0.75','1.13','0.84','1','1.00','2.5','2.75'],[1143811,3,'0.91','-1.25','0.85','0.85','-1.5','0.97','2.75','3'],[1143811,8,'0.95','-1','0.85','0.90','-1.5','0.90','3','2.75'],[1143811,12,'1.00','-1.25','0.89','0.78','-1.5','1.12','2.75','2.75'],[1143811,24,'1.11','-1.25','0.74','0.97','-1.5','0.87','2.75','2.75'],[1153119,3,'0.85','1.25','0.91','0.89','1.5','0.93','2.75','2.75'],[1153119,8,'1.05','1.25','0.75','0.85','1.5','0.95','2.5','2.75'],[1153119,12,'0.82','1.5','1.03','0.87','1.5','0.98','2.75','2.75'],[1153119,24,'0.92','1.25','0.92','0.60','1.25','1.31','2.75','2.75'],[1156312,3,'0.95','0.25','0.81','0.80','0','1.02','2.5','2.5'],[1156312,8,'0.98','0.25','0.83','0.80','0','1.00','2.5','2.5'],[1156312,12,'0.93','0','0.96','0.85','0','1.04','2.5','2.5'],[1156312,24,'1.02','0.25','0.82','1.20','0.25','0.67','2.5','2.5'],[1186747,3,'0.85','-1.25','0.91','0.80','-0.75','1.02','3.25','3.25'],[1186747,8,'0.85','-0.75','0.95','0.73','-1','1.08','3','3.5'],[1186747,12,'0.65','-1.25','1.20','0.88','-1','0.97','3.5','3.5'],[1186747,24,'0.80','-1.25','1.04','0.54','-1.25','1.42','3.5','3.5'],[1186753,3,'0.77','2.75','0.99','0.97','2.75','0.85','3.75','3.75'],[1186753,8,'0.90','3','0.90','0.85','2.75','0.95','4','3.75'],[1186753,12,'0.65','2.75','1.25','1.12','2.75','0.77','3.75','3.75'],[1186753,24,'0.84','2.75','1.00','1.13','2.75','0.72','3.75','3.75'],[1186755,3,'0.85','0.25','0.91','0.99','0.5','0.83','2.5','2.75'],[1186755,8,'1.00','0.25','0.80','0.68','0.25','1.15','2.5','2.5'],[1186755,12,'0.76','0.25','1.10','0.67','0.25','1.24','2.5','2.5'],[1186755,24,'0.84','0.25','0.96','0.68','0.25','1.19','2.5','2.5'],[1189244,3,'0.91','-0.25','0.85','0.90','-0.25','0.92','2.75','2.75'],[1189244,8,'0.78','-0.25','1.03','0.90','-0.25','0.90','2.75','2.75'],[1189244,12,'0.90','-0.25','0.95','0.90','-0.25','0.95','2.75','2.75'],[1189244,24,'0.68','-0.25','1.19','0.80','-0.25','1.04','2.75','2.75'],[1201147,3,'0.91','0.25','0.85','0.80','0.25','1.02','2.75','2.75'],[1201147,8,'1.03','0.25','0.78','0.83','0.25','0.98','2.75','2.75'],[1201147,12,'1.06','0.25','0.84','0.81','0.25','1.08','2.5','2.5'],[1201147,24,'0.93','0.25','0.87','0.83','0.25','1.01','2.75','2.75'],[1201853,3,'0.75','2.25','1.01','0.70','2.25','1.06','3.75','3.75'],[1201853,8,'0.85','2','0.95','0.73','2.25','1.08','3.5','3.75'],[1201853,12,'0.55','2.25','1.41','0.62','2.25','1.31','3.75','3.75'],[1201853,24,'0.65','2.25','1.17','0.77','2.25','1.07','3.75','3.75'],[1201882,3,'0.77','2','0.99','0.65','2','1.20','3.75','3.75'],[1201882,8,'0.68','2','1.15','0.65','2','1.20','3.75','3.75'],[1201882,12,'0.71','2','1.17','0.62','2','1.31','3.75','3.75'],[1201882,24,'0.80','2','1.00','0.71','2','1.14','3.75','3.75'],[1202511,3,'0.77','2.5','0.99','0.97','2.5','0.85','3.5','3.5'],[1202511,8,'0.78','2.5','1.03','0.78','2.5','1.03','3.5','3.5'],[1202511,24,'0.80','2.5','1.04','0.94','2.5','0.90','3.5','3.25'],[1206955,3,'0.91','3.25','0.85','0.87','3.75','0.95','4.25','4.5'],[1206955,12,'0.91','3.25','0.86','0.83','3.75','0.94','4.25','4.25'],[1206955,24,'0.94','3.25','0.90','0.58','3.25','1.35','4.25','4.25'],[1206956,3,'0.96','-0.75','0.80','1.13','-0.75','0.70','2.5','2.5'],[1206956,12,'0.96','-0.75','0.82','1.09','-0.75','0.71','2.5','2.5'],[1206956,24,'0.93','-0.75','0.91','1.14','-0.75','0.71','2.5','2.5'],[1207001,3,'0.77','3.25','0.99','0.97','3.25','0.85','4.25','4.25'],[1207001,8,'0.85','3.5','0.95','0.95','3.25','0.85','4.5','4.25'],[1207001,12,'0.73','3.25','1.07','1.08','3.25','0.72','4.25','4.25'],[1207001,24,'0.79','3.25','1.05','0.96','3.25','0.88','4.25','4.25'],[1207004,3,'0.96','1.5','0.80','0.85','1','0.97','2.75','2.5'],[1207004,8,'0.80','1.5','1.00','1.08','1.25','0.72','2.75','2.5'],[1207004,12,'1.02','1.25','0.77','1.11','1.25','0.70','2.75','2.25'],[1207004,24,'1.02','1.5','0.82','1.11','1.25','0.74','2.75','2.5']];
	Vs_hOdds := stringKit.GetBetweenStr(strHtml, "Vs_hOdds=[[", "]];")
	if Vs_hOdds == "" {
		Vs_hOdds = stringKit.GetBetweenStr(strHtml, "Vs_hOdds = [[", "]];")
	} //欧赔数据
	Vs_eOdds := stringKit.GetBetweenStr(strHtml, "Vs_eOdds=[[", "]];")
	if Vs_eOdds == "" {
		Vs_eOdds = stringKit.GetBetweenStr(strHtml, "Vs_eOdds = [[", "]];")
	}

	mpAsiaOdds := this.getHtml2AsiaOdds(Vs_hOdds)
	mpEuropeOdds := this.getHtml2EuropeOdds(Vs_eOdds)

	AHistoryRaceItem := this.getHtml2HistoryRaceItem(h_data, aRaceInfo.RaceInfoId, models.HistoryType_A)
	BHistoryRaceItem := this.getHtml2HistoryRaceItem(a_data, aRaceInfo.RaceInfoId, models.HistoryType_B)
	AvBHistoryRaceItem := this.getHtml2HistoryRaceItem(v_data, aRaceInfo.RaceInfoId, models.HistoryType_V)

	AHistoryRaceItem = this.updateHistoryRaceInfo(aRaceInfo.Id, aRaceInfo.RaceInfoId, AHistoryRaceItem, mpAsiaOdds, mpEuropeOdds)
	BHistoryRaceItem = this.updateHistoryRaceInfo(aRaceInfo.Id, aRaceInfo.RaceInfoId, BHistoryRaceItem, mpAsiaOdds, mpEuropeOdds)
	AvBHistoryRaceItem = this.updateHistoryRaceInfo(aRaceInfo.Id, aRaceInfo.RaceInfoId, AvBHistoryRaceItem, mpAsiaOdds, mpEuropeOdds)
	arrInjuryInfo := this.getHtml2InjuryInfoData(strHtml, aRaceInfo.Id, aRaceInfo.RaceInfoId)
	AFutureRaceInfo, BFutureRaceInfo, _ := this.getHtml2FutureRaceInfo(strHtml, aRaceInfo.Id, aRaceInfo.RaceInfoId)
	aIntegralRankingInfoData, bIntegralRankingInfoData, _ := this.getHtml2IntegralRanking(strHtml, aRaceInfo.Id, aRaceInfo.RaceInfoId)

	aABHistoryRaceData.RaceInfo = aRaceInfo
	aABHistoryRaceData.HistoryRaceInfoExtA = AHistoryRaceItem
	aABHistoryRaceData.HistoryRaceInfoExtB = BHistoryRaceItem
	aABHistoryRaceData.HistoryRaceInfoExtH = AvBHistoryRaceItem
	aABHistoryRaceData.InjuryInfoList = arrInjuryInfo
	aABHistoryRaceData.FutureRaceInfoDataA = AFutureRaceInfo
	aABHistoryRaceData.FutureRaceInfoDataB = BFutureRaceInfo
	aABHistoryRaceData.IntegralRankingInfoDataA = aIntegralRankingInfoData
	aABHistoryRaceData.IntegralRankingInfoDataB = bIntegralRankingInfoData

	return aABHistoryRaceData, e
}

func (this *AnalystWork) saveABHistoryRaceData(aABHistoryRaceData ABHistoryRaceData) error {
	o := orm.NewOrm()
	aABHistoryRaceData.RaceInfo.AddUpdate(o, "HomeTeam", "GuestTeam")
	models.MultiSaveHistoryRaceInfoExt(o, aABHistoryRaceData.HistoryRaceInfoExtA)
	models.MultiSaveHistoryRaceInfoExt(o, aABHistoryRaceData.HistoryRaceInfoExtB)
	models.MultiSaveHistoryRaceInfoExt(o, aABHistoryRaceData.HistoryRaceInfoExtH)

	models.MultiSaveInjury(o, aABHistoryRaceData.InjuryInfoList)
	models.MultiSaveFutureRaceInfo(o, aABHistoryRaceData.FutureRaceInfoDataA)
	models.MultiSaveFutureRaceInfo(o, aABHistoryRaceData.FutureRaceInfoDataB)

	aABHistoryRaceData.IntegralRankingInfoDataA.AddUpdate(o)
	aABHistoryRaceData.IntegralRankingInfoDataB.AddUpdate(o)

	return nil
}

func (this *AnalystWork) GetAsiaOddsInfoData(aRaceInfo models.FaRaceInfo) ([]models.FaAsiaOddsInfo, error) {
	strHtml, e := this.mAnalystHttpClient.GetAsianOddsHtml(aRaceInfo.RaceInfoId)
	if e != nil {
		return nil, nil
	}
	arrData := this.getHtml2AsiaOddsInfoData(strHtml, aRaceInfo.Id, aRaceInfo.RaceInfoId)

	return arrData, nil
}

func (this *AnalystWork) getHtml2AsiaOddsInfoData(strHttpData string, MainId, RaceID int64) []models.FaAsiaOddsInfo {
	arrData := make([]models.FaAsiaOddsInfo, 0)
	dom, e := goquery.ParseString(strHttpData)
	if e != nil {
		return arrData
	}
	if strHttpData == "" {
		return arrData
	}

	vsNode := dom.Find(".vs")
	if len(vsNode) == 0 {
		return arrData
	}
	aa := vsNode.Eq(0)

	strTime := aa[0].Child[1].Child[2].Data
	arrRaceTimeItem := strings.Split(strTime, " ")
	newStrTime := fmt.Sprintf("%s-%s:00", arrRaceTimeItem[0], arrRaceTimeItem[1])
	raceTime, _ := time.Parse(newStrTime, timeKit.DateTimeLayout)
	fmt.Println(raceTime)

	tableNodes := dom.Find("#oddsDetail")
	arrStrTable := goqueryKit.GetTableNodes2Arr(tableNodes)

	tableRowCount := len(arrStrTable)
	if tableRowCount < 1 {
		return arrData
	}

	firstRowLen := len(arrStrTable[0])
	for i := 0; i < firstRowLen-2; i++ {
		aAsiaOddsInfo := models.FaAsiaOddsInfo{
			RaceInfoId: RaceID, CompanyName: arrStrTable[0][i].Text}
		aAsiaOddsInfo.SysId = this.SysId
		aAsiaOddsInfo.MainId = MainId

		arrData = append(arrData, aAsiaOddsInfo)
	}
	timeIndex := firstRowLen - 1
	scoreIndex := firstRowLen - 2

	for rowI := tableRowCount - 1; rowI >= 1; rowI-- {
		rowNode := arrStrTable[rowI]

		strScore := rowNode[scoreIndex].Text
		if strScore != "" {
			break
		}

		//4-18
		arrDay := strings.Split(rowNode[timeIndex].Node.Child[0].Data, "-")
		m := strconvEx.StrTry2Int(arrDay[0], 0)
		d := strconvEx.StrTry2Int(arrDay[1], 0)
		//17:52
		arrTime := strings.Split(rowNode[timeIndex].Node.Child[2].Data, ":")
		h := strconvEx.StrTry2Int(arrTime[0], 0)
		mi := strconvEx.StrTry2Int(arrTime[1], 0)
		OddsTime := time.Date(raceTime.Year(), time.Month(m), d, h, mi, 0, 0, time.Local)
		for j := 0; j < len(arrData); j++ {
			if len(rowNode[j].Node.Child) == 0 {
				continue
			}
			strOdds2 := goqueryKit.GetNodeText(rowNode[j].Node.Child[0])
			Odds2 := footballKit.GoalCn2Goal(strOdds2)
			Odds1 := strconvEx.StrTry2Float64(goqueryKit.GetNodeText(rowNode[j].Node.Child[3]), 0)
			Odds3 := strconvEx.StrTry2Float64(goqueryKit.GetNodeText(rowNode[j].Node.Child[5]), 0)
			arrData[j].EndOddsTime = OddsTime
			arrData[j].EndOdds2 = Odds2
			arrData[j].EndOdds1 = Odds1
			arrData[j].EndOdds3 = Odds3

			if arrData[j].BeginOdds1 == 0 {
				arrData[j].BeginOddsTime = OddsTime
				arrData[j].BeginOdds2 = Odds2
				arrData[j].BeginOdds1 = Odds1
				arrData[j].BeginOdds3 = Odds3
			}
		}
	}

	return arrData
}

func (this *AnalystWork) GetEuropeOddsInfoData(aRaceInfo models.FaRaceInfo) ([]models.FaEuropeOddsInfo, error) {
	strHtml, e := this.mAnalystHttpClient.EuropeOddsInfoDataHtml(aRaceInfo.RaceInfoId)
	if e != nil {
		return nil, nil
	}
	arrData, e := this.getHtml2EuropeOddsInfoData(strHtml, aRaceInfo.Id, aRaceInfo.RaceInfoId)

	return arrData, e
}

type GameData struct {
	CompanyId     int
	CompanyName   string
	BeginOddsTime time.Time
}

/*
*
115|119835766|William Hill|8|4.8|1.3|11.34|18.9|69.77|90.7|12|5.5|1.2|7.59|16.55|75.86|91.03|1.09|0.88|0.90|2022,11-1,19,18,50,00|威廉希*(英国)|1|0
*/
func (this *AnalystWork) getHtml2Game(rowItem string) (GameData, error) {
	const (
		Index_CompanyId     = 1
		Index_CompanyName   = 2
		Index_BeginOddsTime = 20
	)
	rowItem1 := strings.Split(rowItem, "|")
	CompanyId := strconvEx.StrTry2Int(rowItem1[Index_CompanyId], 0)
	CompanyName := rowItem1[Index_CompanyName]
	StrBeginOddsTime := rowItem1[Index_BeginOddsTime]
	BeginOddsTime, e := time.Parse("2006,1-2,15,04,05,00", StrBeginOddsTime)
	if e != nil {
		BeginOddsTime, _ = time.Parse("2006,1-2,15,4,5,00", StrBeginOddsTime)
	}

	aGameData := GameData{}
	aGameData.CompanyId = CompanyId
	aGameData.CompanyName = CompanyName
	aGameData.BeginOddsTime = BeginOddsTime

	return aGameData, nil
}

/*
*
119846245^8.6|5.6|1.33|11-20 02:07|0.78|0.90|1.00;7.4|5.3|1.38|11-19 18:08|0.67|0.85|1.03;7.9|5.5|1.36|11-19 07:28|0.72|0.88|1.02;8.6|5.7|1.33|11-19 05:41|0.78|0.91|1.00;
*/
func (this *AnalystWork) getHtml2EuropeOddsInfoAll(rowItem string) (models.FaEuropeOddsInfo, error) {
	EuropeOddsInfoBegin := models.FaEuropeOddsInfo{}
	rowItem1 := strings.Split(rowItem, "|")
	if len(rowItem1) < 2 {
		return EuropeOddsInfoBegin, fmt.Errorf("有问题的数据")
	}

	const (
		index_BeginOdds1 = 0
		index_BeginOdds2 = 1
		index_BeginOdds3 = 2

		index_Time = 3

		index_KellyOdds1 = 4
		index_KellyOdds2 = 5
		index_KellyOdds3 = 6
	)
	EuropeOddsInfoBegin.CompanyId = strconvEx.StrTry2Int(rowItem1[0], 0)

	oddsItem := strings.Split(rowItem1[1], ";")
	iLen := len(oddsItem)
	if iLen <= 0 {
		return EuropeOddsInfoBegin, fmt.Errorf("有问题的数据")
	}
	itemEnd := strings.Split(oddsItem[0], `|`)
	itemBegin := strings.Split(oddsItem[iLen-1], `|`)

	toEuropeOddsInfo := func(itemBegin []string) models.FaEuropeOddsInfo {
		strTime := itemBegin[index_Time]
		tt := time.Now()
		newTime := fmt.Sprintf("%d-%s", tt.Year(), strTime)

		aEuropeOddsInfo := models.FaEuropeOddsInfo{}
		aEuropeOddsInfo.BeginOdds1 = strconvEx.StrTry2Float64(itemBegin[index_BeginOdds1], 0)
		aEuropeOddsInfo.BeginOdds2 = strconvEx.StrTry2Float64(itemBegin[index_BeginOdds2], 0)
		aEuropeOddsInfo.BeginOdds3 = strconvEx.StrTry2Float64(itemBegin[index_BeginOdds3], 0)

		aEuropeOddsInfo.KellyOdds1 = strconvEx.StrTry2Float64(itemBegin[index_KellyOdds1], 0)
		aEuropeOddsInfo.KellyOdds2 = strconvEx.StrTry2Float64(itemBegin[index_KellyOdds2], 0)
		aEuropeOddsInfo.KellyOdds3 = strconvEx.StrTry2Float64(itemBegin[index_KellyOdds3], 0)
		aEuropeOddsInfo.BeginOddsTime, _ = time.Parse("2006-01-02 15:04", newTime)

		return aEuropeOddsInfo
	}
	EuropeOddsInfoBegin = toEuropeOddsInfo(itemBegin)
	EuropeOddsInfoEnd := toEuropeOddsInfo(itemEnd)
	EuropeOddsInfoBegin.EndOdds1 = EuropeOddsInfoEnd.BeginOdds1
	EuropeOddsInfoBegin.EndOdds2 = EuropeOddsInfoEnd.BeginOdds2
	EuropeOddsInfoBegin.EndOdds3 = EuropeOddsInfoEnd.BeginOdds3
	EuropeOddsInfoBegin.EndOddsTime = EuropeOddsInfoEnd.BeginOddsTime

	return EuropeOddsInfoBegin, nil
}

/*
*
8.6|5.6|1.33|11-20 02:07|0.78|0.90|1.00
*/
func (this *AnalystWork) getHtml2EuropeOddsInfo1(CompanyId, Year int, rowItem string) (models.FaEuropeOddsInfo, error) {
	EuropeOddsInfoBegin := models.FaEuropeOddsInfo{}
	rowItem1 := strings.Split(rowItem, "|")
	if len(rowItem1) < 2 {
		return EuropeOddsInfoBegin, fmt.Errorf("有问题的数据")
	}

	const (
		index_BeginOdds1 = 0
		index_BeginOdds2 = 1
		index_BeginOdds3 = 2

		index_Time = 3

		index_KellyOdds1 = 4
		index_KellyOdds2 = 5
		index_KellyOdds3 = 6
	)
	EuropeOddsInfoBegin.CompanyId = CompanyId
	EuropeOddsInfoBegin.BeginOdds1 = strconvEx.StrTry2Float64(rowItem1[index_BeginOdds1], 0)
	EuropeOddsInfoBegin.BeginOdds2 = strconvEx.StrTry2Float64(rowItem1[index_BeginOdds2], 0)
	EuropeOddsInfoBegin.BeginOdds3 = strconvEx.StrTry2Float64(rowItem1[index_BeginOdds3], 0)

	strTime := rowItem1[index_Time]
	newTime := ""
	if Year == 0 {
		tt := time.Now()
		newTime = fmt.Sprintf("%d-%s", tt.Year(), strTime)
	} else {
		newTime = fmt.Sprintf("%d-%s", Year, strTime)
	}
	EuropeOddsInfoBegin.BeginOddsTime, _ = time.Parse("2006-01-02 15:04", newTime)

	EuropeOddsInfoBegin.KellyOdds1 = strconvEx.StrTry2Float64(rowItem1[index_KellyOdds1], 0)
	EuropeOddsInfoBegin.KellyOdds2 = strconvEx.StrTry2Float64(rowItem1[index_KellyOdds2], 0)
	EuropeOddsInfoBegin.KellyOdds3 = strconvEx.StrTry2Float64(rowItem1[index_KellyOdds3], 0)

	return EuropeOddsInfoBegin, nil
}

/*
*
281|120603548|Bet 365|2.25|3.1|3.3|41.53|30.15|28.32|93.45|2.5|3|3|37.5|31.25|31.25|93.75|0.93|0.92|0.97|2022,12-1,30,17,00,00|36*(英国)|1|0
*/
func (this *AnalystWork) getHtml2EuropeOddsInfo2(rowItem string) (models.FaEuropeOddsInfo, error) {
	EuropeOddsInfoBegin := models.FaEuropeOddsInfo{}
	rowItem1 := strings.Split(rowItem, "|")
	if len(rowItem1) < 2 {
		return EuropeOddsInfoBegin, fmt.Errorf("有问题的数据")
	}

	const (
		index_CompanyId   = 1
		index_CompanyName = 2
		index_EndOdds1    = 10
		index_EndOdds2    = 11
		index_EndOdds3    = 12

		index_KellyOdds1 = 17
		index_KellyOdds2 = 18
		index_KellyOdds3 = 19

		index_Time  = 20
		index_TypeZ = 22
		index_TypeE = 23
	)
	//2022,12-1,30,17,00,00
	strTime := rowItem1[index_Time]
	arrTime := strings.Split(strTime, ",")
	arrMonth := strings.Split(arrTime[1], "-")

	y := strconvEx.StrTry2Int(arrTime[0], 0)
	d := strconvEx.StrTry2Int(arrTime[2], 0)
	m := strconvEx.StrTry2Int(arrMonth[0], 0)
	oddsTime := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)

	CompanyId := strconvEx.StrTry2Int(rowItem1[index_CompanyId], 0)
	EuropeOddsInfoBegin.CompanyId = CompanyId
	EuropeOddsInfoBegin.CompanyName = rowItem1[index_CompanyName]

	EuropeOddsInfoBegin.EndOddsTime = oddsTime
	EuropeOddsInfoBegin.EndOdds1 = strconvEx.StrTry2Float64(rowItem1[index_EndOdds1], 0)
	EuropeOddsInfoBegin.EndOdds2 = strconvEx.StrTry2Float64(rowItem1[index_EndOdds2], 0)
	EuropeOddsInfoBegin.EndOdds3 = strconvEx.StrTry2Float64(rowItem1[index_EndOdds3], 0)

	EuropeOddsInfoBegin.KellyOdds1 = strconvEx.StrTry2Float64(rowItem1[index_KellyOdds1], 0)
	EuropeOddsInfoBegin.KellyOdds2 = strconvEx.StrTry2Float64(rowItem1[index_KellyOdds2], 0)
	EuropeOddsInfoBegin.KellyOdds3 = strconvEx.StrTry2Float64(rowItem1[index_KellyOdds3], 0)
	EuropeOddsInfoBegin.TypeZ = strconvEx.StrTry2Int(rowItem1[index_TypeZ], 0)
	EuropeOddsInfoBegin.TypeE = strconvEx.StrTry2Int(rowItem1[index_TypeE], 0)

	return EuropeOddsInfoBegin, nil
}

func (this *AnalystWork) getHtml2EuropeOddsInfoData(strHttpData string, MainId, RaceID int64) ([]models.FaEuropeOddsInfo, error) {
	arrData := make([]models.FaEuropeOddsInfo, 0)
	if strHttpData == "" {
		return arrData, nil
	}
	strGame := stringKit.GetBetweenStr(strHttpData, `game=Array("`, `");`)
	gameItems := strings.Split(strGame, `","`)

	strGameDetail := stringKit.GetBetweenStr(strHttpData, `gameDetail=Array("`, `");`)
	GameDetailItems := strings.Split(strGameDetail, `","`)

	mapEuropeCOdds := make(map[string]string)
	mapBeginEurope := make(map[int]models.FaEuropeOddsInfo)

	for _, rowItem := range gameItems {
		Europe, _ := this.getHtml2EuropeOddsInfo2(rowItem)
		mapBeginEurope[Europe.CompanyId] = Europe
	}

	for _, rowItem := range GameDetailItems {
		rowItem1 := strings.Split(rowItem, "^")
		if len(rowItem1) < 2 {
			continue
		}
		CompanyId := strconvEx.StrTry2Int(rowItem1[0], 0)

		mapEuropeCOdds[rowItem1[0]] = ""
		oddsItem := strings.Split(rowItem1[1], ";")
		iLen := len(oddsItem)
		if iLen <= 0 {
			continue
		}

		Europe, _ := this.getHtml2EuropeOddsInfo1(CompanyId, 0, oddsItem[iLen-2])
		Europe1, _ := this.getHtml2EuropeOddsInfo1(CompanyId, 0, oddsItem[1])

		Europe.EndOdds1 = Europe1.BeginOdds1
		Europe.EndOdds2 = Europe1.BeginOdds2
		Europe.EndOdds3 = Europe1.BeginOdds3
		Europe.EndOddsTime = Europe1.BeginOddsTime

		mpE := mapBeginEurope[CompanyId]
		mpE.BeginOdds1 = Europe.BeginOdds1
		mpE.BeginOdds2 = Europe.BeginOdds2
		mpE.BeginOdds3 = Europe.BeginOdds3
		mpE.RaceInfoId = RaceID
		mpE.SysId = this.SysId
		mpE.MainId = MainId

		mpE.EndOddsTime = time.Date(mpE.EndOddsTime.Year(), mpE.EndOddsTime.Month(), Europe1.BeginOddsTime.Day(), Europe1.BeginOddsTime.Hour(), Europe1.BeginOddsTime.Minute(), Europe1.BeginOddsTime.Second(), 0, time.UTC)
		mpE.BeginOddsTime = Europe.BeginOddsTime
		if mpE.EndOddsTime.Sub(Europe.BeginOddsTime) < 0 {
			mpE.BeginOddsTime = time.Date(Europe.BeginOddsTime.Year()-1, Europe.BeginOddsTime.Month(), Europe.BeginOddsTime.Day(), Europe.BeginOddsTime.Hour(), Europe.BeginOddsTime.Minute(), Europe.BeginOddsTime.Second(), 0, time.UTC)
		}
		mapBeginEurope[mpE.CompanyId] = mpE
	}

	arrData = make([]models.FaEuropeOddsInfo, len(mapBeginEurope))
	i := 0
	for _, e := range mapBeginEurope {
		arrData[i] = e
		i++
	}

	return arrData, nil
}

func (this *AnalystWork) GetGSOddsInfoData(aRaceInfo models.FaRaceInfo) ([]models.FaGSOddsInfo, error) {
	strHtml, e := this.mAnalystHttpClient.GSOddsInfoDataHtml(aRaceInfo.RaceInfoId)
	if e != nil {
		return nil, nil
	}
	arrData, e := this.getHtml2GSOddsInfoData(strHtml, aRaceInfo.Id, aRaceInfo.RaceInfoId)

	return arrData, e
}

func (this *AnalystWork) getHtml2GSOddsInfoData(strHttpData string, MainId, RaceID int64) ([]models.FaGSOddsInfo, error) {
	arrData := make([]models.FaGSOddsInfo, 0)
	dom, e := goquery.ParseString(strHttpData)
	if e != nil {
		return arrData, e
	}

	vsNode := dom.Find(".vs")
	if len(vsNode) == 0 {
		return arrData, nil
	}
	aa := vsNode.Eq(0)
	if len(aa) == 0 {
		return arrData, nil
	}
	strTime := aa[0].Child[1].Child[2].Data
	arrRaceTimeItem := strings.Split(strTime, " ")
	newStrTime := fmt.Sprintf("%s-%s:00", arrRaceTimeItem[0], arrRaceTimeItem[1])
	raceTime, _ := time.Parse(newStrTime, timeKit.DateTimeLayout)
	fmt.Println(raceTime)

	tableNodes := dom.Find("#oddsDetail")
	arrStrTable := goqueryKit.GetTableNodes2Arr(tableNodes)

	tableRowCount := len(arrStrTable)
	if tableRowCount < 1 {
		return arrData, nil
	}

	firstRowLen := len(arrStrTable[0])
	for i := 0; i < firstRowLen-2; i++ {
		aOddsInfo := models.FaGSOddsInfo{
			RaceInfoId: RaceID, CompanyName: arrStrTable[0][i].Text}
		aOddsInfo.SysId = this.SysId
		aOddsInfo.MainId = MainId

		arrData = append(arrData, aOddsInfo)
	}
	timeIndex := firstRowLen - 1
	scoreIndex := firstRowLen - 2

	for rowI := tableRowCount - 1; rowI >= 1; rowI-- {
		rowNode := arrStrTable[rowI]

		strScore := rowNode[scoreIndex].Text
		if strScore != "" {
			break
		}

		//4-18
		arrDay := strings.Split(rowNode[timeIndex].Node.Child[0].Data, "-")
		m := strconvEx.StrTry2Int(arrDay[0], 0)
		d := strconvEx.StrTry2Int(arrDay[1], 0)
		//17:52
		arrTime := strings.Split(rowNode[timeIndex].Node.Child[2].Data, ":")
		h := strconvEx.StrTry2Int(arrTime[0], 0)
		mi := strconvEx.StrTry2Int(arrTime[1], 0)
		OddsTime := time.Date(raceTime.Year(), time.Month(m), d, h, mi, 0, 0, time.Local)
		for j := 0; j < len(arrData); j++ {
			if len(rowNode[j].Node.Child) < 5 {
				continue
			}
			strOdds2 := goqueryKit.GetNodeText(rowNode[j].Node.Child[0])
			Odds2 := footballKit.GoalCn2Goal(strOdds2)
			Odds1 := strconvEx.StrTry2Float64(goqueryKit.GetNodeText(rowNode[j].Node.Child[3]), 0)
			Odds3 := strconvEx.StrTry2Float64(goqueryKit.GetNodeText(rowNode[j].Node.Child[5]), 0)
			arrData[j].EndOddsTime = OddsTime
			arrData[j].EndOdds2 = Odds2
			arrData[j].EndOdds1 = Odds1
			arrData[j].EndOdds3 = Odds3

			if arrData[j].BeginOdds1 == 0 {
				arrData[j].BeginOddsTime = OddsTime
				arrData[j].BeginOdds2 = Odds2
				arrData[j].BeginOdds1 = Odds1
				arrData[j].BeginOdds3 = Odds3
			}
		}
	}
	return arrData, nil
}

func (this *AnalystWork) GetEur2EurCale(arrInfo []models.FaEuropeOddsInfo) models.FaEuropeOddCalculate {
	aEurOddsCalc := models.FaEuropeOddCalculate{}
	vagBeginZOdds0 := 0.0
	vagBeginHOdds1 := 0.0
	vagBeginKOdds2 := 0.0

	vagEndZOdds3 := 0.0
	vagEndHOdds4 := 0.0
	vagEndKOdds5 := 0.0

	iLen := len(arrInfo)
	fLen := float64(iLen)

	for _, eurOdds := range arrInfo {
		vagBeginZOdds0 += eurOdds.BeginOdds1
		vagBeginHOdds1 += eurOdds.BeginOdds2
		vagBeginKOdds2 += eurOdds.BeginOdds3

		vagEndZOdds3 += eurOdds.EndOdds1
		vagEndHOdds4 += eurOdds.EndOdds2
		vagEndKOdds5 += eurOdds.EndOdds3
	}
	vagBeginZOdds0 /= fLen
	vagBeginHOdds1 /= fLen
	vagBeginKOdds2 /= fLen

	vagEndZOdds3 /= fLen
	vagEndHOdds4 /= fLen
	vagEndKOdds5 /= fLen

	aEurOddsCalc.Begin00 = vagBeginZOdds0
	aEurOddsCalc.Begin01 = vagBeginHOdds1
	aEurOddsCalc.Begin02 = vagBeginKOdds2

	aEurOddsCalc.Begin03 = 1.0 / vagBeginZOdds0
	aEurOddsCalc.Begin04 = 1.0 / vagBeginHOdds1
	aEurOddsCalc.Begin05 = 1.0 / vagBeginKOdds2

	aEurOddsCalc.Begin06 = 1.0/vagBeginHOdds1 + 1.0/vagBeginKOdds2
	aEurOddsCalc.Begin07 = 1.0/vagBeginZOdds0 + 1.0/vagBeginKOdds2
	aEurOddsCalc.Begin08 = 1.0/vagBeginZOdds0 + 1.0/vagBeginHOdds1

	aEurOddsCalc.Begin09 = aEurOddsCalc.Begin06 + aEurOddsCalc.Begin07 - aEurOddsCalc.Begin08
	aEurOddsCalc.Begin10 = aEurOddsCalc.Begin07 + aEurOddsCalc.Begin08 - aEurOddsCalc.Begin06
	aEurOddsCalc.Begin11 = 1.0/vagBeginZOdds0 + 1.0/vagBeginHOdds1 + 1.0/vagBeginKOdds2

	aEurOddsCalc.Begin12 = aEurOddsCalc.Begin06 - aEurOddsCalc.Begin09
	aEurOddsCalc.Begin13 = aEurOddsCalc.Begin08 - aEurOddsCalc.Begin10

	aEurOddsCalc.Begin14 = math.Abs(aEurOddsCalc.Begin09 - aEurOddsCalc.Begin10)
	aEurOddsCalc.Begin15 = 1 / aEurOddsCalc.Begin11
	aEurOddsCalc.Begin16 = iLen

	J := math.Abs(aEurOddsCalc.Begin06) - (math.Abs(aEurOddsCalc.Begin05) + math.Abs(aEurOddsCalc.Begin04))
	if J < 0 {
		aEurOddsCalc.Begin17 = math.Abs(J)
	}
	if J > 0 {
		aEurOddsCalc.Begin18 = math.Abs(J)
	}

	J = math.Abs(aEurOddsCalc.Begin08) - (math.Abs(aEurOddsCalc.Begin03) + math.Abs(aEurOddsCalc.Begin04))
	if J < 0 {
		aEurOddsCalc.Begin19 = math.Abs(J)
	}
	if J > 0 {
		aEurOddsCalc.Begin20 = math.Abs(J)
	}

	if vagEndZOdds3 != 0 && vagEndKOdds5 != 0 && vagEndKOdds5 != 0 {

		aEurOddsCalc.End00 = vagEndZOdds3
		aEurOddsCalc.End01 = vagEndHOdds4
		aEurOddsCalc.End02 = vagEndKOdds5

		aEurOddsCalc.End03 = 1.0 / vagEndZOdds3
		aEurOddsCalc.End04 = 1.0 / vagEndHOdds4
		aEurOddsCalc.End05 = 1.0 / vagEndKOdds5

		aEurOddsCalc.End06 = 1.0/vagEndHOdds4 + 1.0/vagEndKOdds5
		aEurOddsCalc.End07 = 1.0/vagEndZOdds3 + 1.0/vagEndKOdds5
		aEurOddsCalc.End08 = 1.0/vagEndZOdds3 + 1.0/vagEndHOdds4

		aEurOddsCalc.End09 = aEurOddsCalc.End06 + aEurOddsCalc.End07 - aEurOddsCalc.End08
		aEurOddsCalc.End10 = aEurOddsCalc.End07 + aEurOddsCalc.End08 - aEurOddsCalc.End06

		aEurOddsCalc.End11 = 1.0/vagEndZOdds3 + 1.0/vagEndHOdds4 + 1.0/vagEndKOdds5

		aEurOddsCalc.End12 = aEurOddsCalc.End06 - aEurOddsCalc.End09
		aEurOddsCalc.End13 = aEurOddsCalc.End08 - aEurOddsCalc.End10

		aEurOddsCalc.End14 = math.Abs(aEurOddsCalc.End09 - aEurOddsCalc.End10)
		aEurOddsCalc.End15 = 1.0 / aEurOddsCalc.End11
		aEurOddsCalc.End16 = iLen

		J = math.Abs(aEurOddsCalc.End06) - (math.Abs(aEurOddsCalc.End05) + math.Abs(aEurOddsCalc.End04))
		if J < 0 {
			aEurOddsCalc.End17 = math.Abs(J)
		}
		if J > 0 {
			aEurOddsCalc.End18 = math.Abs(J)
		}

		J = math.Abs(aEurOddsCalc.End08) - (math.Abs(aEurOddsCalc.End03) + math.Abs(aEurOddsCalc.End04))
		if J < 0 {
			aEurOddsCalc.End19 = math.Abs(J)
		}
		if J > 0 {
			aEurOddsCalc.End20 = math.Abs(J)
		}
	}
	aEurOddsCalc.Begin21 = (aEurOddsCalc.Begin06 - aEurOddsCalc.End06) - (aEurOddsCalc.Begin07 - aEurOddsCalc.End07) - (aEurOddsCalc.Begin08 - aEurOddsCalc.End08)
	aEurOddsCalc.Begin22 = (aEurOddsCalc.Begin08 - aEurOddsCalc.End08) - (aEurOddsCalc.Begin07 - aEurOddsCalc.End07) - (aEurOddsCalc.Begin06 - aEurOddsCalc.End06)
	aEurOddsCalc.Begin23 = aEurOddsCalc.Begin09 - aEurOddsCalc.End09
	aEurOddsCalc.Begin24 = aEurOddsCalc.Begin10 - aEurOddsCalc.End10

	J = math.Abs(aEurOddsCalc.Begin06-aEurOddsCalc.End06) - math.Abs(aEurOddsCalc.Begin08-aEurOddsCalc.End08)
	if J > 0 {
		aEurOddsCalc.Begin25 = J
	}
	if J < 0 {
		aEurOddsCalc.Begin27 = J
	}
	aEurOddsCalc.Begin26 = math.Abs(aEurOddsCalc.Begin07 - aEurOddsCalc.End07)
	J = math.Abs(aEurOddsCalc.Begin03-aEurOddsCalc.End03) - math.Abs(aEurOddsCalc.Begin05-aEurOddsCalc.End05)
	if J > 0 {
		aEurOddsCalc.Begin28 = J
	}
	if J < 0 {
		aEurOddsCalc.Begin30 = J
	}
	aEurOddsCalc.Begin29 = math.Abs(aEurOddsCalc.Begin04 - aEurOddsCalc.End04)

	aEurOddsCalc.Begin31 = aEurOddsCalc.Begin12 - aEurOddsCalc.End12
	aEurOddsCalc.Begin32 = aEurOddsCalc.Begin13 - aEurOddsCalc.End13

	return aEurOddsCalc
}
