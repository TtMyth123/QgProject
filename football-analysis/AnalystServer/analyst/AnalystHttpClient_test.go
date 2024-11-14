package analyst

import (
	"fmt"
	"testing"
)

const (
	basePath1 = "titan007.com"
	basePath2 = "http://1x2.nowscore.com"
	basePath3 = "http://data.nowgoal.com"
)

func Test_001GetBFData(t *testing.T) {
	aHttpClient := NewAnalystHttpClient(basePath1, basePath2, basePath3)

	htmlData, e := aHttpClient.GetBFData()
	fmt.Println(e, htmlData)
	fmt.Println()
}

func Test_001GetAsianOddsHtml(t *testing.T) {
	aHttpClient := NewAnalystHttpClient(basePath1, basePath2, basePath3)

	termId := int64(4918)
	htmlData, e := aHttpClient.GetAsianOddsHtml(termId)

	fmt.Println(e, htmlData)
}

func Test_001GetHcRaceInfoHtml(t *testing.T) {
	aHttpClient := NewAnalystHttpClient(basePath1, basePath2, basePath3)

	strDay := "20221118"
	htmlData, e := aHttpClient.GetHcRaceInfoHtml(strDay)

	fmt.Println(e, htmlData)
}
