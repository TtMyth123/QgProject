package analyst

import (
	"fmt"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/orm"
)

type SHistoryFightInfo struct {
	ATeamId   int64
	ATeamName string
	BTeamId   int64
	BTeamName string
}

func (this *AnalystWork) GetHistoryFightInfoList(o orm.Ormer, Id int64) (interface{}, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	arrData := make([]SHistoryFightInfo, 0)
	sqlArgs := make([]interface{}, 0)
	history_type := 3

	sqlWhere := "a.history_type = ? and a.main_id=?"
	sqlArgs = append(sqlArgs, history_type, Id)

	sql := fmt.Sprintf(`select a.a_score, a.b_score, a.a_team_name, a.b_team_name from fa_history_race_info_ext a where %s`,
		sqlWhere)
	_, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)

	return arrData, e
}

type SFutureRaceInfoList struct {
	LeagueName string
}

func (this *AnalystWork) GetFutureRaceInfoList(o orm.Ormer, Id int64, future_type int) (interface{}, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	arrData := make([]SFutureRaceInfoList, 0)
	sqlArgs := make([]interface{}, 0)

	sqlWhere := "a.future_type = ? and a.main_id=?"
	sqlArgs = append(sqlArgs, future_type, Id)

	sql := fmt.Sprintf(`select a.league_name from %s a where %s`,
		mconst.TableName_FaFutureRaceInfo, sqlWhere)
	_, e := o.Raw(sql, sqlArgs).QueryRows(&arrData)

	return arrData, e
}

type SIntegralRankingInfo struct {
	TotalTeamWin      int //总胜
	TotalTeamFlat     int //总平
	TotalTeamLose     int //总负
	Lately6TeamWin    int //近6场,胜
	Lately6TeamFlat   int //近6场,平
	Lately6TeamLose   int //近6场,负
	Lately6TeamGoal   int //近6场,进球
	Lately6TeamFumble int //近6场,失球
}

func (this *AnalystWork) GetIntegralRankingInfo(o orm.Ormer, Id int64, IntegralType int) (interface{}, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	aData := SIntegralRankingInfo{}
	sqlArgs := make([]interface{}, 0)

	sqlWhere := "a.integral_type = ? and a.main_id=?"
	sqlArgs = append(sqlArgs, IntegralType, Id)

	sql := fmt.Sprintf(`select a.total_team_win,a.total_team_flat, a.total_team_lose, a.lately6_team_win, a.lately6_team_flat, a.lately6_team_lose, a.lately6_team_goal, a.lately6_team_fumble from %s a where %s`,
		mconst.TableName_FaIntegralRankingInfo, sqlWhere)
	e := o.Raw(sql, sqlArgs).QueryRow(&aData)

	return aData, e
}

type SExcelFData struct {
}

func (this *AnalystWork) GetExcelFData(o orm.Ormer, Id int64, DataType int) (interface{}, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	aData := models.FaExcelFData{}
	sqlArgs := make([]interface{}, 0)

	sqlWhere := "a.data_type = ? and a.main_id=?"
	sqlArgs = append(sqlArgs, DataType, Id)

	sql := fmt.Sprintf(`select a.* from %s a where %s`,
		mconst.TableName_FaExcelFData, sqlWhere)
	e := o.Raw(sql, sqlArgs).QueryRow(&aData)

	return aData, e
}

type SEuropeOddCalculate struct {
	models.FaEuropeOddCalculate
}

func (a *AnalystWork) GetEuropeOddCalculate(o orm.Ormer, Id int64, DataType int) (interface{}, error) {
	if o == nil {
		o = orm.NewOrm()
	}
	if o == nil {
		o = orm.NewOrm()
	}
	aData := SEuropeOddCalculate{}
	sqlArgs := make([]interface{}, 0)

	sqlWhere := "a.data_type = ? and a.main_id=?"
	sqlArgs = append(sqlArgs, DataType, Id)

	sql := fmt.Sprintf(`select a.* from %s a where %s`,
		mconst.TableName_FaEuropeOddCalculate, sqlWhere)
	e := o.Raw(sql, sqlArgs).QueryRow(&aData)

	return nil, e
}
