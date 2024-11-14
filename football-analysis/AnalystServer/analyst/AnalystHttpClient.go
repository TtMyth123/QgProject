package analyst

import (
	"fmt"
	"github.com/TtMyth123/kit/httpClientKit"
	"github.com/TtMyth123/kit/timeKit"
	"github.com/astaxie/beego"
	"golang.org/x/text/encoding/simplifiedchinese"
	"time"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

type AnalystHttpClient struct {
	httpClient1       *httpClientKit.HttpClient
	httpClient1VipUrl string

	httpClient2                *httpClientKit.HttpClient
	httpClient2baseNowscoreUrl string
	httpClient2baseNowgoalUrl  string

	httpClient3      *httpClientKit.HttpClient
	httpClient3BfUrl string
	wwwVip           string
	wwwLive          string
	wwwBf            string
	www1x2d          string
	wwwZq            string
}

/*
*
分析
baseVipUrl = "http://vip.titan007.com"
baseLiveUrl = "http://live.titan007.com"
baseBfUrl = "http://bf.titan007.com/"
baseNowscoreUrl = "http://1x2.nowscore.com"
*/
func NewAnalystHttpClient(baseUrl, baseNowScoreUrl, baseNowGoalUrl string) *AnalystHttpClient {
	if baseUrl == "" {
		baseUrl = beego.AppConfig.String("football::baseUrl")
	}
	if baseNowScoreUrl == "" {
		baseNowScoreUrl = beego.AppConfig.String("football::NowScoreUrl")
	}
	if baseNowGoalUrl == "" {
		baseNowGoalUrl = beego.AppConfig.String("football::NowGoalUrl")
	}

	aHttpClient := new(AnalystHttpClient)
	aHttpClient.wwwVip = "vip"
	aHttpClient.wwwLive = "live"
	aHttpClient.wwwBf = "bf"
	aHttpClient.www1x2d = "1x2d"
	aHttpClient.wwwZq = "zq"

	aHttpClient.httpClient1 = httpClientKit.GetHttpClient("")
	aHttpClient.httpClient1VipUrl = baseUrl

	aHttpClient.httpClient2 = httpClientKit.GetHttpClient("")
	aHttpClient.httpClient2baseNowscoreUrl = baseNowScoreUrl
	aHttpClient.httpClient2baseNowgoalUrl = baseNowGoalUrl

	return aHttpClient
}

func (this *AnalystHttpClient) GetAsianOddsHtml(termId int64) (string, error) {
	//http://vip.win007.com/AsianOdds_n.aspx?id=1744933
	//fullUrl := fmt.Sprintf(`%s/AsianOdds_n.aspx?id=%d`, this.httpClient1VipUrl, termId)

	fullUrl := fmt.Sprintf(`%s/AsianOdds_n.aspx?id=%d`, this.getBaseUrl(this.wwwVip), termId)
	html, e := this.httpClient1.GetString(fullUrl)

	return html, e
}

/*
*
获取欧赔数据。
*/
func (this *AnalystHttpClient) EuropeOddsInfoDataHtml(termId int64) (string, error) {
	//EuropeOddsInfoDataURL := "http://1x2.nowscore.com/%d.js"
	//fullUrl := fmt.Sprintf(EuropeOddsInfoDataURL, termId)
	//fullUrl := fmt.Sprintf(`%s/%d.js`, this.httpClient2baseNowscoreUrl, termId)
	fullUrl := fmt.Sprintf(`%s/%d.js`, this.getBaseUrl(this.www1x2d), termId)
	//fullUrl := fmt.Sprintf(`%s/%d.js`, this.httpClient2baseNowscoreUrl, termId)
	html, e := this.httpClient1.GetString(fullUrl)
	return html, e
}

/*
*
球探大小盘口数据
*/
func (this *AnalystHttpClient) GSOddsInfoDataHtml(termId int64) (string, error) {
	//http://vip.win007.com/AsianOdds_n.aspx?id=1744933
	//EuropeOddsInfoDataURL := "http://vip.win007.com/OverDown_n.aspx?id=%d"
	//fullUrl := fmt.Sprintf(EuropeOddsInfoDataURL, termId)

	fullUrl := fmt.Sprintf(`%s/OverDown_n.aspx?id=%d`, this.getBaseUrl(this.wwwVip), termId)
	html, e := this.httpClient1.GetString(fullUrl)
	return html, e
}

/*
*
百家欧赔。
*/
func (this *AnalystHttpClient) HundredEOddsHtml(termId int) (string, error) {
	//http://vip.win007.com/AsianOdds_n.aspx?id=1744933
	//AsianOdds_nURL := "http://1x2d.win007.com/%d.js"
	//fullUrl := fmt.Sprintf(AsianOdds_nURL, termId)

	fullUrl := fmt.Sprintf(`%s/OverDown_n.aspx?id=%d`, this.getBaseUrl(this.www1x2d), termId)
	html, e := this.httpClient1.GetString(fullUrl)
	return html, e
}

/*
*
完成的比赛 (回查)
http://bf.titan007.com/football/Over_20221118.htm
*/
func (this *AnalystHttpClient) GetHcRaceInfoHtml(time string) (string, error) {
	fullUrl := fmt.Sprintf(`%s/football/Over_%s.htm`, this.getBaseUrl(this.wwwBf), time)
	html, e := this.httpClient1.GetBytes(fullUrl)
	strHtml := ConvertByte2String(html, GB18030)
	return strHtml, e
}

/*
*
完成的比赛 (回查)
http://data.nowgoal.com/MatchByCountry.aspx?date=2019-9-1&orderby=time&type=2
*/
func (this *AnalystHttpClient) MatchByCountryHtml(time string) (string, error) {
	orderby := "time"
	itype := 2
	//fullUrl := fmt.Sprintf(`http://data.nowgoal.com/MatchByCountry.aspx?date=%s&orderby=%s&type=%d`, time, orderby, itype)
	fullUrl := fmt.Sprintf(`%s/MatchByCountry.aspx?date=%s&orderby=%s&type=%d`, this.httpClient2baseNowgoalUrl, time, orderby, itype)
	html, e := this.httpClient1.GetString(fullUrl)
	return html, e
}
func (this *AnalystHttpClient) getBaseUrl(www string) string {
	strUrl := fmt.Sprintf(`%s%s.%s`, `http://`, www, this.httpClient1VipUrl)
	return strUrl
}

/*
*
获取赛事列表
*/
func (this *AnalystHttpClient) GetBFData() (string, error) {
	//http://live.titan007.com/vbsxml/bfdata_ut.js?r=0071667825746000
	t := time.Now()
	longT := timeKit.GetJavaTimeLong(t)
	//fullUrl := fmt.Sprintf(`%s/vbsxml/bfdata_ut.js?r=007%d`, this.httpClient2LiveUrl, longT)
	fullUrl := fmt.Sprintf(`%s/vbsxml/bfdata_ut.js?r=007%d`, this.getBaseUrl(this.wwwLive), longT)
	paramsHeader := make(map[string]string)
	//paramsHeader["Referer"] = this.httpClient1VipUrl
	paramsHeader["Referer"] = this.getBaseUrl(this.wwwLive)
	paramsHeader["cache-control"] = `no-cache`
	paramsHeader["Cookie"] = `win007BfCookie=null; bfWin007FirstMatchTime=2021,1,1,08,00,00`
	html, e := this.httpClient1.GetHeader(fullUrl, paramsHeader)

	return string(html), e
}

func (this *AnalystHttpClient) GetABHistoryRaceDataHtml(RaceID int64) (string, error) {
	//fullUrl := fmt.Sprintf(`http://zq.win007.com/analysis/%dcn.htm`, RaceID)

	fullUrl := fmt.Sprintf(`%s/analysis/%dcn.htm`, this.getBaseUrl(this.wwwZq), RaceID)
	html, e := this.httpClient1.GetString(fullUrl)

	return html, e
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
