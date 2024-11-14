package controllers

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/controllers/Bll"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/controllers/base"
)

type FApiController struct {
	base.AuthorBaseController
}

func (c *FApiController) GetSysId() int64 {
	SysId, ok := c.GetInt64("SysId", 0)
	if ok != nil {
		SysId = c.SysId
	}
	return SysId
}

func (c *FApiController) UpdateTmpRaceList() {
	SysId := c.GetSysId()
	data, e := Bll.GetAndAddRaceTmpList(SysId)
	c.JsonResultEx(data, nil, e)
}

/*
*
添加赛事（将中间的赛事列表添加到右边）
*/
func (c *FApiController) AddRaceIds2RaceTable() {
	Ids := c.GetString("Ids")
	SysId := c.GetSysId()
	data, e := Bll.AddRaceIds2RaceTable(SysId, Ids)

	c.JsonResultEx(data, nil, e)
}

/*
*
删除赛事（将右边的赛事删除，）
*/
func (c *FApiController) DelJsRaceTable() {
	Ids := c.GetString("Ids")
	SysId := c.GetSysId()
	data, e := Bll.DelJsRaceTable(SysId, Ids)

	c.JsonResultEx(data, nil, e)
}

/*
*
“抓取数据” getandsaveracedata
*/
func (c *FApiController) GetAndSaveRaceData() {
	Ids := c.GetString("Ids")
	SysId := c.GetSysId()
	data, e := Bll.GetAndSaveRaceData(SysId, Ids)

	c.JsonResultEx(data, nil, e)
}

func (c *FApiController) GetRaceDataList() {
	Ids := c.GetString("Ids")
	Js, _ := c.GetInt("Js")
	SysId := c.GetSysId()

	data, e := Bll.GetRaceDataList(SysId, Ids, Js == 1)
	c.JsonResultEx(data, nil, e)
}
