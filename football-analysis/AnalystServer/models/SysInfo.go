package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type SysInfo struct {
	BaseInfo
}

func (a *SysInfo) TableName() string {
	return mconst.TableName_SysInfo
}
func InitSysInfo(o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	c, _ := o.QueryTable(mconst.TableName_SysInfo).Count()
	if c == 0 {
		arrData := make([]SysInfo, 0)
		arrData = append(arrData, SysInfo{BaseInfo{Id: 1}})
		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			logs.Error(e)
		}
	}
}
