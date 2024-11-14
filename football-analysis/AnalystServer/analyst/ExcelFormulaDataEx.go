package analyst

import (
	"fmt"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models"
	"github.com/TtMyth123/kit/ExcelFormula/ExcelKit"
	"github.com/TtMyth123/kit/strconvEx"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type ExcelFormulaDataEx struct {
	xlFile *excelize.File
}

func NewExcelFormulaDataEx(excelFormulaFile string) *ExcelFormulaDataEx {
	if excelFormulaFile == "" {
		excelFormulaFile = "./conf/A.xlsx"
	}

	aExcelFormulaDataEx := ExcelFormulaDataEx{}
	f, err := excelize.OpenFile(excelFormulaFile)
	if err != nil {
		return nil
	}
	aExcelFormulaDataEx.xlFile = f
	return &aExcelFormulaDataEx
}
func (this *ExcelFormulaDataEx) LoadAsiaOddsData(lstAsiaOddsInfo []models.FaAsiaOddsInfo) {
	for row := 4; row <= 17; row++ {
		for col := 1; col <= 8; col++ {
			strCell := ExcelKit.GetCellXY2Str(row, col)
			this.xlFile.SetCellValue(SheetName_AsiaOdds, strCell, "")
		}
	}
	iLen := len(lstAsiaOddsInfo)
	if iLen > 13 {
		iLen = 13
	}
	for i := 0; i < iLen; i++ {
		row := i + 4
		this.xlFile.SetCellValue(SheetName_AsiaOdds, ExcelKit.GetCellXY2Str(row, 1), lstAsiaOddsInfo[i].CompanyName)
		this.xlFile.SetCellValue(SheetName_AsiaOdds, ExcelKit.GetCellXY2Str(row, 3), lstAsiaOddsInfo[i].BeginOdds1)
		this.xlFile.SetCellValue(SheetName_AsiaOdds, ExcelKit.GetCellXY2Str(row, 4), lstAsiaOddsInfo[i].BeginCOdds2)
		this.xlFile.SetCellValue(SheetName_AsiaOdds, ExcelKit.GetCellXY2Str(row, 5), lstAsiaOddsInfo[i].BeginOdds3)
		this.xlFile.SetCellValue(SheetName_AsiaOdds, ExcelKit.GetCellXY2Str(row, 6), lstAsiaOddsInfo[i].EndOdds1)
		this.xlFile.SetCellValue(SheetName_AsiaOdds, ExcelKit.GetCellXY2Str(row, 7), lstAsiaOddsInfo[i].EndCOdds2)
		this.xlFile.SetCellValue(SheetName_AsiaOdds, ExcelKit.GetCellXY2Str(row, 8), lstAsiaOddsInfo[i].EndOdds3)
	}
}

func (this *ExcelFormulaDataEx) LoadEuropeOddsData(lstEuropeOddsInfo []models.FaEuropeOddsInfo) {
	for row := 3; row <= 300; row++ {
		for col := 1; col <= 15; col++ {
			this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, col), "")
		}
	}
	iLen := len(lstEuropeOddsInfo)
	OddsMax := models.FaEuropeOddsInfo{}
	OddsMin := models.FaEuropeOddsInfo{}
	OddsSum := models.FaEuropeOddsInfo{}
	if iLen > 0 {
		OddsMin = lstEuropeOddsInfo[0]
		OddsMax = lstEuropeOddsInfo[0]
	}

	i := 0
	for i = 0; i < iLen; i++ {
		row := i + 6

		this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 1), i)
		this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 2), lstEuropeOddsInfo[i].CompanyName)
		this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 3), lstEuropeOddsInfo[i].BeginOdds1)
		this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 4), lstEuropeOddsInfo[i].BeginOdds2)
		this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 5), lstEuropeOddsInfo[i].BeginOdds3)
		this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 6), lstEuropeOddsInfo[i].EndOdds1)
		this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 7), lstEuropeOddsInfo[i].EndOdds2)
		this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 8), lstEuropeOddsInfo[i].EndOdds3)

		if OddsMax.BeginOdds1 < lstEuropeOddsInfo[i].BeginOdds1 {
			OddsMax.BeginOdds1 = lstEuropeOddsInfo[i].BeginOdds1
		}
		if OddsMax.BeginOdds2 < lstEuropeOddsInfo[i].BeginOdds2 {
			OddsMax.BeginOdds2 = lstEuropeOddsInfo[i].BeginOdds2
		}
		if OddsMax.BeginOdds3 < lstEuropeOddsInfo[i].BeginOdds3 {
			OddsMax.BeginOdds3 = lstEuropeOddsInfo[i].BeginOdds3
		}
		if OddsMax.EndOdds1 < lstEuropeOddsInfo[i].EndOdds1 {
			OddsMax.EndOdds1 = lstEuropeOddsInfo[i].EndOdds1
		}
		if OddsMax.EndOdds2 < lstEuropeOddsInfo[i].EndOdds2 {
			OddsMax.EndOdds2 = lstEuropeOddsInfo[i].EndOdds2
		}
		if OddsMax.EndOdds3 < lstEuropeOddsInfo[i].EndOdds3 {
			OddsMax.EndOdds3 = lstEuropeOddsInfo[i].EndOdds3
		}

		if OddsMin.BeginOdds1 > lstEuropeOddsInfo[i].BeginOdds1 {
			OddsMin.BeginOdds1 = lstEuropeOddsInfo[i].BeginOdds1
		}
		if OddsMin.BeginOdds2 > lstEuropeOddsInfo[i].BeginOdds2 {
			OddsMin.BeginOdds2 = lstEuropeOddsInfo[i].BeginOdds2
		}
		if OddsMin.BeginOdds3 > lstEuropeOddsInfo[i].BeginOdds3 {
			OddsMin.BeginOdds3 = lstEuropeOddsInfo[i].BeginOdds3
		}
		if OddsMin.EndOdds1 > lstEuropeOddsInfo[i].EndOdds1 {
			OddsMin.EndOdds1 = lstEuropeOddsInfo[i].EndOdds1
		}
		if OddsMin.EndOdds2 > lstEuropeOddsInfo[i].EndOdds2 {
			OddsMin.EndOdds2 = lstEuropeOddsInfo[i].EndOdds2
		}
		if OddsMin.EndOdds3 > lstEuropeOddsInfo[i].EndOdds3 {
			OddsMin.EndOdds3 = lstEuropeOddsInfo[i].EndOdds3
		}

		OddsSum.BeginOdds1 += lstEuropeOddsInfo[i].BeginOdds1
		OddsSum.BeginOdds2 += lstEuropeOddsInfo[i].BeginOdds2
		OddsSum.BeginOdds3 += lstEuropeOddsInfo[i].BeginOdds3
		OddsSum.EndOdds1 += lstEuropeOddsInfo[i].EndOdds1
		OddsSum.EndOdds2 += lstEuropeOddsInfo[i].EndOdds2
		OddsSum.EndOdds3 += lstEuropeOddsInfo[i].EndOdds3

		if row == 300 {
			break
		}
	}
	OddsSum.BeginOdds1 = OddsSum.BeginOdds1 / float64(i)
	OddsSum.BeginOdds2 = OddsSum.BeginOdds2 / float64(i)
	OddsSum.BeginOdds3 = OddsSum.BeginOdds3 / float64(i)
	OddsSum.EndOdds1 = OddsSum.EndOdds1 / float64(i)
	OddsSum.EndOdds2 = OddsSum.EndOdds2 / float64(i)
	OddsSum.EndOdds3 = OddsSum.EndOdds3 / float64(i)

	row := 3
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 2), "最大值")
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 3), OddsMax.BeginOdds1)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 4), OddsMax.BeginOdds2)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 5), OddsMax.BeginOdds3)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 6), OddsMax.EndOdds1)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 7), OddsMax.EndOdds2)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 8), OddsMax.EndOdds3)

	row = 4
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 2), "最小值")
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 3), OddsMin.BeginOdds1)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 4), OddsMin.BeginOdds2)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 5), OddsMin.BeginOdds3)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 6), OddsMin.EndOdds1)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 7), OddsMin.EndOdds2)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 8), OddsMin.EndOdds3)

	row = 5
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 2), "平均值")
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 3), OddsSum.BeginOdds1)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 4), OddsSum.BeginOdds2)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 5), OddsSum.BeginOdds3)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 6), OddsSum.EndOdds1)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 7), OddsSum.EndOdds2)
	this.xlFile.SetCellValue(SheetName_E, ExcelKit.GetCellXY2Str(row, 8), OddsSum.EndOdds3)
}

func (this *ExcelFormulaDataEx) LoadHistoryRaceInfo(arrAHistoryRace, arrBHistoryRace []models.FaHistoryRaceInfoExt, ATeamName, BTeamName string) {
	this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(10, 1), ATeamName+"近")
	this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(25, 1), BTeamName+"近")

	this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(1, 1), fmt.Sprintf(`%s(主)   VS   %s`, ATeamName, BTeamName))
	this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(2, 1), "")
	this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(3, 1), "")

	i := 0
	iLen := len(arrAHistoryRace)
	for row := 14; row <= 23; row++ {
		if i >= iLen {
			break
		}
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 1), arrAHistoryRace[i].LeagueName)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 2), arrAHistoryRace[i].RaceTime.Format(timeKit.DateTimeLayout))
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 3), arrAHistoryRace[i].ATeamName)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 4), fmt.Sprintf("%d-%d (%d-%d)", arrAHistoryRace[i].AScore, arrAHistoryRace[i].BScore, arrAHistoryRace[i].AHalfScore, arrAHistoryRace[i].BHalfScore))

		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 5), arrAHistoryRace[i].BTeamName)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 6), arrAHistoryRace[i].EndAsiaOdds1)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 7), arrAHistoryRace[i].EndCAsiaOdds2)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 8), arrAHistoryRace[i].EndAsiaOdds3)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 9), arrAHistoryRace[i].EndEuropeOdds1)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 10), arrAHistoryRace[i].EndEuropeOdds2)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 11), arrAHistoryRace[i].EndEuropeOdds3)

		i++
	}

	i = 0
	iLen = len(arrBHistoryRace)
	for row := 29; row <= 38; row++ {
		if i >= iLen {
			break
		}
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 1), arrBHistoryRace[i].LeagueName)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 2), arrBHistoryRace[i].RaceTime.Format(timeKit.DateTimeLayout))
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 3), arrBHistoryRace[i].ATeamName)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 4), fmt.Sprintf("%d-%d (%d-%d)", arrAHistoryRace[i].AScore, arrAHistoryRace[i].BScore, arrAHistoryRace[i].AHalfScore, arrAHistoryRace[i].BHalfScore))

		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 5), arrBHistoryRace[i].BTeamName)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 6), arrBHistoryRace[i].EndAsiaOdds1)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 7), arrBHistoryRace[i].EndCAsiaOdds2)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 8), arrBHistoryRace[i].EndAsiaOdds3)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 9), arrBHistoryRace[i].EndEuropeOdds1)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 10), arrBHistoryRace[i].EndEuropeOdds2)
		this.xlFile.SetCellValue(SheetName_2, ExcelKit.GetCellXY2Str(row, 11), arrBHistoryRace[i].EndEuropeOdds3)
		i++
	}
}

func (this *ExcelFormulaDataEx) GetExcelFormulaData() models.FaExcelFData {
	aExcelFData := models.FaExcelFData{}
	AStrength_B5, _ := this.xlFile.GetCellValue(SheetName_1, "B5")
	AAttack_C5, _ := this.xlFile.GetCellValue(SheetName_1, "C5")
	ADefend_D5, _ := this.xlFile.GetCellValue(SheetName_1, "D5")
	BStrength_B6, _ := this.xlFile.GetCellValue(SheetName_1, "B6")
	BAttack_C6, _ := this.xlFile.GetCellValue(SheetName_1, "C6")

	BDefend_D6, _ := this.xlFile.GetCellValue(SheetName_1, "D6")
	ADValue_B7, _ := this.xlFile.GetCellValue(SheetName_1, "B7")
	BDValue_C7, _ := this.xlFile.GetCellValue(SheetName_1, "C7")
	DValue_D7, _ := this.xlFile.GetCellValue(SheetName_1, "D7")
	GoalInGap_F4, _ := this.xlFile.GetCellValue(SheetName_1, "F4")
	ColligateGap_F5, _ := this.xlFile.GetCellValue(SheetName_1, "F5")
	WheelForce_F6, _ := this.xlFile.GetCellValue(SheetName_1, "F6")
	WinIndex_F7, _ := this.xlFile.GetCellValue(SheetName_1, "F7")
	Skewness1_G8, _ := this.xlFile.GetCellValue(SheetName_1, "G8")
	Skewness2_H8, _ := this.xlFile.GetCellValue(SheetName_1, "H8")
	Skewness3_I8, _ := this.xlFile.GetCellValue(SheetName_1, "I8")

	InitialKurtosis1_G9, _ := this.xlFile.GetCellValue(SheetName_1, "G9")
	InitialKurtosis2_H9, _ := this.xlFile.GetCellValue(SheetName_1, "H9")
	InitialKurtosis3_I9, _ := this.xlFile.GetCellValue(SheetName_1, "I9")

	LateKurtosis1_G10, _ := this.xlFile.GetCellValue(SheetName_1, "G10")
	LateKurtosis2_H10, _ := this.xlFile.GetCellValue(SheetName_1, "H10")
	LateKurtosis3_I10, _ := this.xlFile.GetCellValue(SheetName_1, "I10")

	MaxGL1_B8, _ := this.xlFile.GetCellValue(SheetName_1, "B8")
	MaxGL2_C8, _ := this.xlFile.GetCellValue(SheetName_1, "C8")
	MaxGL3_D8, _ := this.xlFile.GetCellValue(SheetName_1, "D8")

	UnitOffset1_B9, _ := this.xlFile.GetCellValue(SheetName_1, "B9")
	UnitOffset2_C9, _ := this.xlFile.GetCellValue(SheetName_1, "C9")
	UnitOffset3_D9, _ := this.xlFile.GetCellValue(SheetName_1, "D9")

	S2_A52, _ := this.xlFile.GetCellValue(SheetName_2, "A52")
	S2_A53, _ := this.xlFile.GetCellValue(SheetName_2, "A53")
	S2_A64, _ := this.xlFile.GetCellValue(SheetName_2, "A64")
	S2_A65, _ := this.xlFile.GetCellValue(SheetName_2, "A65")

	AStrength := this.getStringValue(AStrength_B5)
	aExcelFData.AStrength, aExcelFData.AStrength0, aExcelFData.AStrength1 = SplitNums(AStrength)
	aExcelFData.AAttack = this.getFloatValue(AAttack_C5)
	aExcelFData.ADefend = this.getFloatValue(ADefend_D5)

	BStrength := this.getStringValue(BStrength_B6)
	aExcelFData.BStrength, aExcelFData.BStrength0, aExcelFData.BStrength1 = SplitNums(BStrength)
	aExcelFData.BAttack = this.getFloatValue(BAttack_C6)
	aExcelFData.BDefend = this.getFloatValue(BDefend_D6)

	ADValueF_B7 := this.getFloatValue(ADValue_B7)
	if ADValueF_B7 != 0 {
		aExcelFData.ADValue = fmt.Sprintf("%g", ADValueF_B7)
	} else {
		aExcelFData.ADValue = this.getStringValue(ADValue_B7)
	}

	BDValueF_C7 := this.getFloatValue(BDValue_C7)
	if BDValueF_C7 != 0 {
		aExcelFData.BDValue = fmt.Sprintf("%g", BDValueF_C7)
	} else {
		aExcelFData.BDValue = this.getStringValue(BDValue_C7)
	}

	DValueF_D7 := this.getFloatValue(DValue_D7)
	if DValueF_D7 != 0 {
		aExcelFData.DValue = fmt.Sprintf("%g", DValueF_D7)
	} else {
		aExcelFData.DValue = this.getStringValue(DValue_D7)
	}

	GoalInGap := this.getStringValue(GoalInGap_F4)
	ColligateGap := this.getStringValue(ColligateGap_F5)
	WheelForce := this.getStringValue(WheelForce_F6)
	WinIndex := this.getStringValue(WinIndex_F7)
	aExcelFData.GoalInGap, aExcelFData.GoalInGap0, aExcelFData.GoalInGap1 = SplitNumsEx(GoalInGap)
	aExcelFData.ColligateGap, aExcelFData.ColligateGap0, aExcelFData.ColligateGap1 = SplitNumsEx(ColligateGap)
	aExcelFData.WheelForce, aExcelFData.WheelForce0, aExcelFData.WheelForce1 = SplitNumsEx(WheelForce)
	aExcelFData.WinIndex, aExcelFData.WinIndex0, aExcelFData.WinIndex1 = SplitNumsEx(WinIndex)

	aExcelFData.Skewness1 = this.getFloatValue(Skewness1_G8)
	aExcelFData.Skewness2 = this.getFloatValue(Skewness2_H8)
	aExcelFData.Skewness3 = this.getFloatValue(Skewness3_I8)

	aExcelFData.InitialKurtosis1 = this.getFloatValue(InitialKurtosis1_G9)
	aExcelFData.InitialKurtosis2 = this.getFloatValue(InitialKurtosis2_H9)
	aExcelFData.InitialKurtosis3 = this.getFloatValue(InitialKurtosis3_I9)

	aExcelFData.LateKurtosis1 = this.getFloatValue(LateKurtosis1_G10)
	aExcelFData.LateKurtosis2 = this.getFloatValue(LateKurtosis2_H10)
	aExcelFData.LateKurtosis3 = this.getFloatValue(LateKurtosis3_I10)

	aExcelFData.MaxGL1 = this.getFloatValue(MaxGL1_B8)
	aExcelFData.MaxGL2 = this.getFloatValue(MaxGL2_C8)
	aExcelFData.MaxGL3 = this.getFloatValue(MaxGL3_D8)

	aExcelFData.UnitOffset1 = this.getFloatValue(UnitOffset1_B9)
	aExcelFData.UnitOffset2 = this.getFloatValue(UnitOffset2_C9)
	aExcelFData.UnitOffset3 = this.getFloatValue(UnitOffset3_D9)

	aExcelFData.S2_A52 = strconvEx.StrTry2Int(this.getStringValue(S2_A52), 0)
	aExcelFData.S2_A53 = strconvEx.StrTry2Int(this.getStringValue(S2_A53), 0)
	aExcelFData.S2_A64 = strconvEx.StrTry2Int(this.getStringValue(S2_A64), 0)
	aExcelFData.S2_A65 = strconvEx.StrTry2Int(this.getStringValue(S2_A65), 0)

	return aExcelFData
}

func (this *ExcelFormulaDataEx) GetExcelFormulaDataEx(lstAsiaOddsInfo []models.FaAsiaOddsInfo,
	lstEuropeOddsInfo []models.FaEuropeOddsInfo, arrAHistoryRace, arrBHistoryRace []models.FaHistoryRaceInfoExt,
	ATeamName, BTeamName string) models.FaExcelFData {
	this.LoadAsiaOddsData(lstAsiaOddsInfo)
	this.LoadEuropeOddsData(lstEuropeOddsInfo)
	this.LoadHistoryRaceInfo(arrAHistoryRace, arrBHistoryRace, ATeamName, BTeamName)

	aExcelFData := this.GetExcelFormulaData()
	return aExcelFData
}

func (this *ExcelFormulaDataEx) getStringValue(tmp string) string {
	if (tmp == "-2146826281") || (tmp == "") {
		tmp = "-"
	}
	return tmp
}

func (this *ExcelFormulaDataEx) getFloatValue(cellValue string) float64 {
	tmp := strconvEx.StrTry2Float64(cellValue, 0)
	tmp, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", tmp), 64)

	return tmp
}
