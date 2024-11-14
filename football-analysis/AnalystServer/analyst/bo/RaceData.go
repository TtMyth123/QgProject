package bo

import "github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models"

type RaceData struct {
	RaceInfo             models.FaRaceInfo
	RaceSetup            models.FaRaceSetup
	HistoryFightInfoList []models.FaHistoryRaceInfoExt
	AIntegralRankingInfo models.FaIntegralRankingInfo
	BIntegralRankingInfo models.FaIntegralRankingInfo

	AFutureRaceInfoList []models.FaFutureRaceInfo
	BFutureRaceInfoList []models.FaFutureRaceInfo
	ExcelFData          models.FaExcelFData
}
