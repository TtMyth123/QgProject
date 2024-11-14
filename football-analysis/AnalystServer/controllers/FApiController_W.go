package controllers

import "github.com/TtMyth123/QgProject/football-analysis/AnalystServer/controllers/Bll"

func (c *FApiController) GetRaceData() {
	RaceInfoId, _ := c.GetInt64("RaceInfoId", 0)
	SysId := c.GetSysId()
	data, e := Bll.GetRaceData(SysId, RaceInfoId)

	c.JsonResultEx(data, nil, e)
}
