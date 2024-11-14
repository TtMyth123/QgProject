package controllers

import "github.com/TtMyth123/QgProject/football-analysis/AnalystServer/controllers/Bll"

func (c *FApiController) UpdateHcTmpRaceList() {
	SysId := c.GetSysId()

	Days := c.GetString("Days")
	data, e := Bll.UpdateHcTmpRaceList(SysId, Days)
	c.JsonResultEx(data, nil, e)
}

func (c *FApiController) AddHcRaceIds2RaceTable() {
	Ids := c.GetString("Ids")
	SysId := c.GetSysId()
	data, e := Bll.AddHcRaceIds2RaceTable(SysId, Ids)

	c.JsonResultEx(data, nil, e)
}

func (c *FApiController) DelHcRaceTable() {
	RaceInfoTmpIds := c.GetString("RaceInfoTmpIds")
	SysId := c.GetSysId()
	data, e := Bll.DelHcRaceTable(SysId, RaceInfoTmpIds)

	c.JsonResultEx(data, nil, e)
}
