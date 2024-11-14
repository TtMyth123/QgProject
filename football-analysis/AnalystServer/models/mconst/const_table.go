package mconst

const (
	TableName_SysInfo = "sys_info"
	//TableName_FaRaceInfo = "fa_race_info_tmp"
	TableName_FaTeamInfo = "fa_team_info"
	//TableName_FaRaceInfo          = "fa_race_info_hc_tmp"
	TableName_FaRaceInfo = "fa_race_info"
	//TableName_FaRaceInfoHc        = "fa_race_info_hc"
	TableName_FaCustomRankingInfo = "fa_custom_ranking_info"

	TableName_FaRaceSetup = "fa_race_setup"

	TableName_FaHistoryRaceInfoExt = "fa_history_race_info_ext"
	TableName_FaHistoryAsiaOdds    = "fa_history_asia_odds"
	TableName_FaHistoryEuropeOdds  = "fa_history_europe_odds"

	TableName_FaInjury              = "fa_injury"
	TableName_FaFutureRaceInfo      = "fa_future_race_info"
	TableName_FaIntegralRankingInfo = "fa_integral_ranking_info"
	TableName_FaAsiaOddsInfo        = "fa_asia_odds_info"
	TableName_FaEuropeOddsInfo      = "fa_europe_odds_info"
	TableName_FaGSOddsInfo          = "fa_gs_odds_info"
	TableName_FaExcelFData          = "fa_excel_f_data"

	TableName_FaEuropeOddCalculate    = "fa_europe_odd_calculate"
	TableName_FaAuthorityCompany      = "fa_authority_company"
	TableName_FaAuthorityCompanyAlias = "fa_authority_company_alias"
)

type a struct {
	FaAuthorityCompanyAlias bool `json:"fa_authority_company_alias"`
	FaAuthorityCompany      bool `json:"fa_authority_company"`
	FaEuropeOddCalculate    bool `json:"fa_europe_odd_calculate"`
	FaExcelFData            bool `json:"fa_excel_f_data"`
	FaGSOddsInfo            bool `json:"fa_gs_odds_info"`
	FaEuropeOddsInfo        bool `json:"fa_europe_odds_info"`
	FaAsiaOddsInfo          bool `json:"fa_asia_odds_info"`
	FaIntegralRankingInfo   bool `json:"fa_integral_ranking_info"`
	FaHistoryAsiaOdds       bool `json:"fa_history_asia_odds"`
	FaFutureRaceInfo        bool `json:"fa_future_race_info"`
	FaInjury                bool `json:"fa_injury"`
	FaEuropeOdds            bool `json:"fa_europe_odds"`
	FaAsiaOdds              bool `json:"fa_asia_odds"`
	FaHistoryRaceInfoExt    bool `json:"fa_history_race_info_ext"`
	FaRaceSetup             bool `json:"fa_race_setup"`
	FaCustomRankingInfo     bool `json:"fa_custom_ranking_info"`
	FaRaceInfo              bool `json:"fa_race_info"`
	//FaRaceInfo              bool `json:"fa_race_info_hc_tmp"`
	FaTeamInfo bool `json:"fa_team_info"`
	//FaRaceInfo bool `json:"fa_race_info_tmp"`
	SysInfo bool `json:"sys_info"`
}
