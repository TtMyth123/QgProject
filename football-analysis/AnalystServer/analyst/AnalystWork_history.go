package analyst

import (
	"fmt"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/footballKit"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models"
	"github.com/TtMyth123/kit/goqueryKit"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/opesun/goquery"
	"strings"
	"time"
)

func (this *AnalystWork) getHtml2AsiaOdds(Vs_hOdds string) map[int64]models.FaHistoryAsiaOdds {
	Vs_hOdds = strings.Replace(Vs_hOdds, `'`, ``, -1)
	mpOddsInfo := make(map[int64]models.FaHistoryAsiaOdds)
	arrAsiaOdds := strings.Split(Vs_hOdds, "],[")
	for _, aAsiaOddsI := range arrAsiaOdds {
		aOddsI := strings.Split(aAsiaOddsI, ",")
		if len(aOddsI) < 8 {
			continue
		}
		aOddsInfo := models.FaHistoryAsiaOdds{}
		aOddsInfo.SysId = this.SysId
		aOddsInfo.RaceInfoId = strconvEx.StrTry2Int64(aOddsI[0], 0)
		aOddsInfo.CompanyId = strconvEx.StrTry2Int(aOddsI[1], 0)
		aOddsInfo.BeginOdds1 = strconvEx.StrTry2Float64(aOddsI[2], 0)
		aOddsInfo.BeginOdds2 = strconvEx.StrTry2Float64(aOddsI[3], 0)
		aOddsInfo.BeginOdds3 = strconvEx.StrTry2Float64(aOddsI[4], 0)

		aOddsInfo.EndOdds1 = strconvEx.StrTry2Float64(aOddsI[5], 0)
		aOddsInfo.EndOdds2 = strconvEx.StrTry2Float64(aOddsI[6], 0)
		aOddsInfo.EndOdds3 = strconvEx.StrTry2Float64(aOddsI[7], 0)

		aOddsInfo.BeginCOdds2 = footballKit.Goal2GoalCn(aOddsInfo.BeginOdds2)
		aOddsInfo.EndCOdds2 = footballKit.Goal2GoalCn(aOddsInfo.EndOdds2)

		mpOddsInfo[aOddsInfo.RaceInfoId] = aOddsInfo
	}
	return mpOddsInfo
}

func (this *AnalystWork) getHtml2EuropeOdds(Vs_eOdds string) map[int64]models.FaHistoryEuropeOdds {
	Vs_eOdds = strings.Replace(Vs_eOdds, `'`, ``, -1)
	mpOddsInfo := make(map[int64]models.FaHistoryEuropeOdds)
	arrAsiaOdds := strings.Split(Vs_eOdds, "],[")
	for _, aAsiaOddsI := range arrAsiaOdds {
		aOddsI := strings.Split(aAsiaOddsI, ",")
		if len(aOddsI) < 8 {
			continue
		}
		aOddsInfo := models.FaHistoryEuropeOdds{}
		aOddsInfo.SysId = this.SysId
		aOddsInfo.RaceInfoId = strconvEx.StrTry2Int64(aOddsI[0], 0)
		aOddsInfo.CompanyId = strconvEx.StrTry2Int(aOddsI[1], 0)
		aOddsInfo.BeginOdds1 = strconvEx.StrTry2Float64(aOddsI[2], 0)
		aOddsInfo.BeginOdds2 = strconvEx.StrTry2Float64(aOddsI[3], 0)
		aOddsInfo.BeginOdds3 = strconvEx.StrTry2Float64(aOddsI[4], 0)

		aOddsInfo.EndOdds1 = strconvEx.StrTry2Float64(aOddsI[5], 0)
		aOddsInfo.EndOdds2 = strconvEx.StrTry2Float64(aOddsI[6], 0)
		aOddsInfo.EndOdds3 = strconvEx.StrTry2Float64(aOddsI[7], 0)

		aOddsInfo.BeginCOdds2 = footballKit.Goal2GoalCn(aOddsInfo.BeginOdds2)
		aOddsInfo.EndCOdds2 = footballKit.Goal2GoalCn(aOddsInfo.EndOdds2)

		mpOddsInfo[aOddsInfo.RaceInfoId] = aOddsInfo
	}
	return mpOddsInfo
}

/*
*
'16-01-23',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>(中)',2696,'<span title="韩国女足  排名:17">韩国女足</span>',2,0,'2-0','-0.75',1,1,-1,1206956],['16-01-21',428,'四国赛','#cc9900',1933,'<span title="中国女足  排名:14">中国女足</span>',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>',0,0,'0-0','1.5',0,1,-1,1207004],['15-12-21',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>(中)',7812,'<span title="特立尼达和多巴哥女足  排名:48">特立尼达</span>',2,1,'1-0','2.5',1,-1,1,1202511],['15-12-17',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>',7812,'<span title="特立尼达和多巴哥女足  排名:48">特立尼达</span>',3,0,'1-0','2.25',1,1,1,1201853],['15-12-14',1366,'国际友谊','#4666bb',4883,'<span title="巴西女足  排名:6">巴西女足</span>',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>',6,0,'4-0','2',-1,-1,1,1201882],['15-12-10',1366,'国际友谊','#4666bb',2337,'<span title="加拿大女足  排名:11">加拿大女</span>(中)',5247,'<span title="墨西哥女足  排名:26">墨西哥女</span>',3,0,'3-0','0.25',-1,-1,1,1201147],['15-07-25',487,'泛美女足','#b00900',2337,'<span title="加拿大女足  排名:9">加拿大女</span>',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',1,2,'0-2','0.25',1,1,1,1156312],['15-07-23',487,'泛美女足','#b00900',4883,'<span title="巴西女足  排名:8">巴西女足</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',4,2,'2-1','1.25',-1,-1,1,1153119],['15-07-19',487,'泛美女足','#b00900',7812,'<span title="特立尼达和多巴哥女足  排名:45">特立尼达</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',1,3,'0-2','-1.25',1,1,1,1143811],['15-07-15',487,'泛美女足','#b00900',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',5937,'<span title="阿根廷女足  排名:36">阿根廷女</span>',3,1,'1-0','0.5',1,1,1,1143809],['15-07-11',487,'泛美女足','#b00900',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',12344,'<span title="哥伦比亚女足  排名:28">哥伦比亚</span>',0,1,'0-0','0',-1,-1,-1,1143807],['15-06-18',388,'女世界杯','#696900',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',2361,'<span title="法国女足  排名:3">法国女足</span>',0,5,'0-4','-1.75',-1,-1,1,1127589],['15-06-14',388,'女世界杯','#696900',2359,'<span title="英格兰女足  排名:6">英格兰女</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',2,1,'0-0','1.25',-1,1,1,1127588],['15-06-10',388,'女世界杯','#696900',12344,'<span title="哥伦比亚女足  排名:28">哥伦比亚</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',1,1,'0-1','-0.25',0,-1,-1,1127586],['15-05-29',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',7829,'<span title="哥斯达黎加女足  排名:37">哥斯达黎</span>',3,0,'3-0','1.25',1,1,1,1128820],['15-05-26',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',7829,'<span title="哥斯达黎加女足  排名:37">哥斯达黎</span>',2,1,'1-1',”,1,-2,1,1128589],['15-05-18',1366,'国际友谊','#4666bb',3862,'<span title="美国女足  排名:2">美国女足</span>',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',5,1,'1-1','3.5',-1,-1,1,1127927],['15-03-11',1366,'国际友谊','#4666bb',2362,'<span class=hp>1</span><span title="意大利女足  排名:14">意大利女</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',2,3,'0-1','0.75',1,1,1,1106456],['15-03-09',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',5564,'<span title="比利时女足  排名:26">比利时女</span>',0,0,'0-0','0.25',0,-1,-1,1105992],['15-03-06',1366,'国际友谊','#4666bb',4519,'<span title="捷克女足  排名:30">捷克女足</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span><span class=hp>1</span>',0,1,'0-0','-0.5',1,1,-1,1102793],['15-03-04',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>(中)',9935,'<span title="南非女足  排名:60">南非女足</span>',2,0,'0-0','1.25',1,1,-1,1103069],['15-02-08',1366,'国际友谊','#4666bb',5247,'<span class=hp>1</span><span title="墨西哥女足  排名:25">墨西哥女</span>',5938,'<span title="厄瓜多尔女足  排名:46">厄瓜多尔</span>',2,0,'0-0',”,1,-2,-1,1092095],['15-02-06',1366,'国际友谊','#4666bb',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',5938,'<span title="厄瓜多尔女足  排名:46">厄瓜多尔</span>',1,0,'1-0',”,1,-2,-1,1090409],['15-01-15',184,'女足四国','#FF2F73',2696,'<span title="韩国女足  排名:">韩国女足</span>(中)',5247,'<span title="墨西哥女足  排名:">墨西哥女</span>',2,1,'1-0','0.75',-1,-1,1,1080891],['15-01-13',184,'女足四国','#FF2F73',2337,'<span title="加拿大女足  排名:9">加拿大女</span>(中)',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',2,1,'0-0','1.25',-1,1,1,1080244],['15-01-11',1366,'国际友谊','#4666bb',1933,'<span title="中国女足  排名:13">中国女足</span>',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',0,0,'0-0','1.5',0,1,-1,1079683],['14-11-28',1371,'女美加杯','#888566',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',12344,'<span title="哥伦比亚女足  排名:31">哥伦比亚</span>',2,0,'0-0',”,1,-2,-1,1071281],['14-11-26',1371,'女美加杯','#888566',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',7829,'<span title="哥斯达黎加女足  排名:40">哥斯达黎</span>',1,0,'0-0',”,1,-2,-1,1071027],['14-11-22',1371,'女美加杯','#888566',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',16109,'<span title="海地女足  排名:56">海地女足</span><span class=hp>1</span>',1,0,'0-0',”,1,-2,-1,1070698],['14-11-20',1371,'女美加杯','#888566',5247,'<span title="墨西哥女足  排名:25">墨西哥女</span>',7812,'<span title="特立尼达和多巴哥女足  排名:125">特立尼达</span>',6,0,'1-0',”,1,-2,1,1070696
*/
func (this *AnalystWork) getHtml2HistoryRaceItem(strHistoryRace string, mainRaceId int64, HistoryType int) []models.FaHistoryRaceInfoExt {
	arrHistoryRaceInfoExt := make([]models.FaHistoryRaceInfoExt, 0)
	strHistoryRace = strings.Replace(strHistoryRace, "'", "", -1)
	bigItems := strings.Split(strHistoryRace, "],[")
	for _, aBigI := range bigItems {
		aHistoryRaceInfoExt := models.FaHistoryRaceInfoExt{}
		aHistoryRaceInfoExt.SysId = this.SysId
		aHistoryRaceInfoExt.MainRaceInfoId = mainRaceId
		aHistoryRaceInfoExt.HistoryType = HistoryType
		subItems := strings.Split(aBigI, ",")
		if len(subItems) < 10 {
			continue
		}
		//比赛时间  subItems[0]
		t, _ := time.Parse("06-01-02", subItems[0])
		aHistoryRaceInfoExt.RaceTime = t
		aHistoryRaceInfoExt.RaceInfoId = strconvEx.StrTry2Int64(subItems[15], 0)
		aHistoryRaceInfoExt.LeagueName = subItems[2]

		aHistoryRaceInfoExt.ATeamId = strconvEx.StrTry2Int64(subItems[4], 0)
		ATeamName := stringKit.GetBetweenStr(subItems[5], `>`, `<`)
		aHistoryRaceInfoExt.ATeamName = ATeamName
		aHistoryRaceInfoExt.ATeamRanking = stringKit.GetBetweenStr(subItems[5], `
	排名:
		`, `
		">`)

		aHistoryRaceInfoExt.BTeamId = strconvEx.StrTry2Int64(subItems[6], 0)
		BTeamName := stringKit.GetBetweenStr(subItems[7], `>`, `<`)
		aHistoryRaceInfoExt.BTeamName = BTeamName
		aHistoryRaceInfoExt.BTeamRanking = stringKit.GetBetweenStr(subItems[7], `排名:`, `">`)

		aHistoryRaceInfoExt.AScore = strconvEx.StrTry2Int(subItems[8], 0)
		aHistoryRaceInfoExt.BScore = strconvEx.StrTry2Int(subItems[9], 0)

		arrHalfScore := strings.Split(subItems[10], "-")
		if len(arrHalfScore) == 2 {
			aHistoryRaceInfoExt.AHalfScore = strconvEx.StrTry2Int(arrHalfScore[0], 0)
			aHistoryRaceInfoExt.BHalfScore = strconvEx.StrTry2Int(arrHalfScore[1], 0)
		}
		arrHistoryRaceInfoExt = append(arrHistoryRaceInfoExt, aHistoryRaceInfoExt)

	}
	return arrHistoryRaceInfoExt
}

func (this *AnalystWork) updateHistoryRaceInfo(MainId, MainRaceInfoId int64, arr []models.FaHistoryRaceInfoExt,
	mpAsiaOdds map[int64]models.FaHistoryAsiaOdds, mpEuropeOdds map[int64]models.FaHistoryEuropeOdds) []models.FaHistoryRaceInfoExt {
	for i, _ := range arr {
		RaceId := arr[i].RaceInfoId
		arr[i].MainId = MainId
		arr[i].MainRaceInfoId = MainRaceInfoId
		arr[i].BeginAsiaOdds1 = mpAsiaOdds[RaceId].BeginOdds1
		arr[i].BeginAsiaOdds2 = mpAsiaOdds[RaceId].BeginOdds2
		arr[i].BeginAsiaOdds3 = mpAsiaOdds[RaceId].BeginOdds3
		arr[i].BeginCAsiaOdds2 = mpAsiaOdds[RaceId].BeginCOdds2

		arr[i].EndAsiaOdds1 = mpAsiaOdds[RaceId].EndOdds1
		arr[i].EndAsiaOdds2 = mpAsiaOdds[RaceId].EndOdds2
		arr[i].EndAsiaOdds3 = mpAsiaOdds[RaceId].EndOdds3
		arr[i].EndCAsiaOdds2 = mpAsiaOdds[RaceId].EndCOdds2

		arr[i].EndEuropeOdds1 = mpEuropeOdds[RaceId].EndOdds1
		arr[i].EndEuropeOdds2 = mpEuropeOdds[RaceId].EndOdds2
		arr[i].EndEuropeOdds3 = mpEuropeOdds[RaceId].EndOdds3

		arr[i].BeginEuropeOdds1 = mpEuropeOdds[RaceId].BeginOdds1
		arr[i].BeginEuropeOdds2 = mpEuropeOdds[RaceId].BeginOdds2
		arr[i].BeginEuropeOdds3 = mpEuropeOdds[RaceId].BeginOdds3
	}

	return arr
}

/*
*
获取伤停数据
*/
func (this *AnalystWork) getHtml2InjuryInfoData(strHttpData string, MainId, RaceID int64) []models.FaInjury {
	arrInjuryInfo := make([]models.FaInjury, 0)
	doc, e := goquery.ParseString(string(strHttpData))
	if e != nil {
		return arrInjuryInfo
	}
	ns := doc.Find("#porlet_21")
	n_table := ns.Find("table")
	if len(n_table) < 4 {
		return arrInjuryInfo
	}
	aTeamA := goqueryKit.GetTableNodes2Arr(n_table.Eq(2))
	aTeamB := goqueryKit.GetTableNodes2Arr(n_table.Eq(3))
	arrInjuryInfoA := this.getInjuryInfoDataByTableNodes2Arr(aTeamA, MainId, RaceID, models.TeamType_1, models.BanType_1)
	arrInjuryInfoB := this.getInjuryInfoDataByTableNodes2Arr(aTeamB, MainId, RaceID, models.TeamType_2, models.BanType_1)

	arrInjuryInfo = append(arrInjuryInfo, arrInjuryInfoA...)
	arrInjuryInfo = append(arrInjuryInfo, arrInjuryInfoB...)

	return arrInjuryInfo
}

func (this *AnalystWork) getInjuryInfoDataByTableNodes2Arr(aTeamA [][]goqueryKit.TdNode, MainId, RaceID int64, TeamType int, BanType int) []models.FaInjury {
	arrInjuryInfo := make([]models.FaInjury, 0)
	rowCount := len(aTeamA)
	for iRow := 2; iRow < rowCount; iRow++ {
		row := aTeamA[iRow]
		iLen := len(row)
		if iLen < 2 {
			break
		}
		BanCrew := strings.Trim(row[0].Text, "")
		if BanCrew == "" {
			break
		}
		aInjuryInfo := models.FaInjury{RaceInfoId: RaceID, BanCrew: BanCrew, TeamType: TeamType, BanType: BanType}
		aInjuryInfo.SysId = this.SysId
		aInjuryInfo.MainId = MainId
		arrInjuryInfo = append(arrInjuryInfo, aInjuryInfo)
	}
	return arrInjuryInfo
}

/*
*
未来赛事数据
*/
func (this *AnalystWork) getHtml2FutureRaceInfo(strHttpData string, MainId, RaceID int64) ([]models.FaFutureRaceInfo, []models.FaFutureRaceInfo, error) {
	doc, e := goquery.ParseString(string(strHttpData))
	if e != nil {
		return nil, nil, e
	}
	ns := doc.Find("#porlet_20")
	n_table := ns.Find("table")
	if len(n_table) < 3 {
		return nil, nil, fmt.Errorf("有问题的数据。")
	}
	aTeamA := goqueryKit.GetTableNodes2Arr(n_table.Eq(1))
	aTeamB := goqueryKit.GetTableNodes2Arr(n_table.Eq(2))
	AFutureRaceInfoList := this.getFutureRaceInfoByTableNodes2Arr(aTeamA, MainId, RaceID, models.FutureType_A)
	BFutureRaceInfoList := this.getFutureRaceInfoByTableNodes2Arr(aTeamB, MainId, RaceID, models.FutureType_B)

	return AFutureRaceInfoList, BFutureRaceInfoList, nil
}

func (this *AnalystWork) getFutureRaceInfoByTableNodes2Arr(aTeamA [][]goqueryKit.TdNode, MainId, RaceID int64, FutureType int) []models.FaFutureRaceInfo {
	arrFutureRaceInfo := make([]models.FaFutureRaceInfo, 0)
	rowCount := len(aTeamA)
	for iRow := 2; iRow < rowCount; iRow++ {
		row := aTeamA[iRow]
		iLen := len(row)
		if iLen < 2 {
			break
		}
		LeagueName := strings.Trim(row[1].Text, "")
		ClashName := strings.Trim(row[2].Text, "")
		if ClashName == "" {
			break
		}

		aInjuryInfo := models.FaFutureRaceInfo{RaceInfoId: RaceID, LeagueName: LeagueName, ClashName: ClashName}
		aInjuryInfo.FutureType = FutureType
		aInjuryInfo.SysId = this.SysId
		aInjuryInfo.MainId = MainId
		arrFutureRaceInfo = append(arrFutureRaceInfo, aInjuryInfo)
	}
	return arrFutureRaceInfo
}

/*
*
获取联赛积分排名
*/
func (this *AnalystWork) getHtml2IntegralRanking(strHttpData string, MainId, RaceID int64) (models.FaIntegralRankingInfo, models.FaIntegralRankingInfo, error) {
	doc, e := goquery.ParseString(string(strHttpData))
	if e != nil {
		return models.FaIntegralRankingInfo{}, models.FaIntegralRankingInfo{}, e
	}
	ns := doc.Find("#porlet_5")
	n_table := ns.Find("table")
	if len(n_table) < 3 {
		return models.FaIntegralRankingInfo{}, models.FaIntegralRankingInfo{}, fmt.Errorf("有问题的数据")
	}

	aTeamA := goqueryKit.GetTableNodes2Arr(n_table.Eq(1))
	aTeamB := goqueryKit.GetTableNodes2Arr(n_table.Eq(2))
	AIntegralRankingInfo := this.getABIntegralRankingByTableNodes2Arr(aTeamA, MainId, RaceID)
	AIntegralRankingInfo.IntegralType = models.IntegralType_A
	BIntegralRankingInfo := this.getABIntegralRankingByTableNodes2Arr(aTeamB, MainId, RaceID)
	BIntegralRankingInfo.IntegralType = models.IntegralType_B
	return AIntegralRankingInfo, BIntegralRankingInfo, nil
}

func (this *AnalystWork) getABIntegralRankingByTableNodes2Arr(aTeamA [][]goqueryKit.TdNode, MainId, RaceID int64) models.FaIntegralRankingInfo {
	aIntegralRankingInfo := models.FaIntegralRankingInfo{}
	aIntegralRankingInfo.SysId = this.SysId
	aIntegralRankingInfo.MainId = MainId
	rowCount := len(aTeamA)
	if rowCount < 6 {
		return aIntegralRankingInfo
	}

	aIntegralRankingInfo.RaceInfoId = RaceID
	iRow := 2
	aIntegralRankingInfo.TotalTeamWin = strconvEx.StrTry2Int(aTeamA[iRow][2].Text, 0)
	aIntegralRankingInfo.TotalTeamFlat = strconvEx.StrTry2Int(aTeamA[iRow][3].Text, 0)
	aIntegralRankingInfo.TotalTeamLose = strconvEx.StrTry2Int(aTeamA[iRow][4].Text, 0)
	aIntegralRankingInfo.TotalTeamGoal = strconvEx.StrTry2Int(aTeamA[iRow][6].Text, 0)
	aIntegralRankingInfo.TotalTeamFumble = strconvEx.StrTry2Int(aTeamA[iRow][7].Text, 0)
	aIntegralRankingInfo.TotalTeamRanking = strconvEx.StrTry2Int(aTeamA[iRow][9].Text, 0)

	iRow = 3
	aIntegralRankingInfo.HomeTeamWin = strconvEx.StrTry2Int(aTeamA[iRow][2].Text, 0)
	aIntegralRankingInfo.HomeTeamFlat = strconvEx.StrTry2Int(aTeamA[iRow][3].Text, 0)
	aIntegralRankingInfo.HomeTeamLose = strconvEx.StrTry2Int(aTeamA[iRow][4].Text, 0)
	aIntegralRankingInfo.HomeTeamGoal = strconvEx.StrTry2Int(aTeamA[iRow][6].Text, 0)
	aIntegralRankingInfo.HomeTeamFumble = strconvEx.StrTry2Int(aTeamA[iRow][7].Text, 0)
	aIntegralRankingInfo.HomeTeamRanking = strconvEx.StrTry2Int(aTeamA[iRow][9].Text, 0)

	iRow = 4
	aIntegralRankingInfo.GuestTeamWin = strconvEx.StrTry2Int(aTeamA[iRow][2].Text, 0)
	aIntegralRankingInfo.GuestTeamFlat = strconvEx.StrTry2Int(aTeamA[iRow][3].Text, 0)
	aIntegralRankingInfo.GuestTeamLose = strconvEx.StrTry2Int(aTeamA[iRow][4].Text, 0)
	aIntegralRankingInfo.GuestTeamGoal = strconvEx.StrTry2Int(aTeamA[iRow][6].Text, 0)
	aIntegralRankingInfo.GuestTeamFumble = strconvEx.StrTry2Int(aTeamA[iRow][7].Text, 0)
	aIntegralRankingInfo.GuestTeamRanking = strconvEx.StrTry2Int(aTeamA[iRow][9].Text, 0)

	iRow = 5
	aIntegralRankingInfo.Lately6TeamWin = strconvEx.StrTry2Int(aTeamA[iRow][2].Text, 0)
	aIntegralRankingInfo.Lately6TeamFlat = strconvEx.StrTry2Int(aTeamA[iRow][3].Text, 0)
	aIntegralRankingInfo.Lately6TeamLose = strconvEx.StrTry2Int(aTeamA[iRow][4].Text, 0)
	aIntegralRankingInfo.Lately6TeamGoal = strconvEx.StrTry2Int(aTeamA[iRow][6].Text, 0)
	aIntegralRankingInfo.Lately6TeamFumble = strconvEx.StrTry2Int(aTeamA[iRow][7].Text, 0)

	return aIntegralRankingInfo
}
