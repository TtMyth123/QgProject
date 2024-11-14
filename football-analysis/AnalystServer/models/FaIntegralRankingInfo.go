package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type FaIntegralRankingInfo struct {
	BaseInfo
	RaceInfoId        int64 `orm:"description(赛事ID)"`      //赛事ID
	IntegralType      int   `orm:"description(1：A队，2：B队)"` //1：A队，2：B队
	TotalTeamWin      int   `orm:"description(总胜)"`        //总胜
	TotalTeamFlat     int   `orm:"description(总平)"`        //总平
	TotalTeamLose     int   `orm:"description(总负)"`        //总负
	TotalTeamGoal     int   `orm:"description(总进球)"`       //总进球
	TotalTeamFumble   int   `orm:"description(总失球)"`       //总失球
	HomeTeamWin       int   `orm:"description(主场胜)"`       //主场胜
	HomeTeamFlat      int   `orm:"description(主场平)"`       //主场平
	HomeTeamLose      int   `orm:"description(主场负)"`       //主场负
	HomeTeamGoal      int   `orm:"description(主场进球)"`      //主场进球
	HomeTeamFumble    int   `orm:"description(主场失球)"`      //主场失球
	GuestTeamWin      int   `orm:"description(客场胜)"`       //客场胜
	GuestTeamFlat     int   `orm:"description(客场平)"`       //客场平
	GuestTeamLose     int   `orm:"description(客场负)"`       //客场负
	GuestTeamGoal     int   `orm:"description(客场进球)"`      //客场进球
	GuestTeamFumble   int   `orm:"description(客场失球)"`      //客场失球
	Lately6TeamWin    int   `orm:"description(最近6场胜)"`     //最近6场胜
	Lately6TeamFlat   int   `orm:"description(最近6场平)"`     //最近6场平
	Lately6TeamLose   int   `orm:"description(最近6场负)"`     //最近6场负
	Lately6TeamGoal   int   `orm:"description(最近6场进球)"`    //最近6场进球
	Lately6TeamFumble int   `orm:"description(最近6场失球)"`    //最近6场失球
	TotalTeamRanking  int   `orm:"description(排名)"`        //排名
	HomeTeamRanking   int   `orm:"description(主场排名)"`      //主场排名
	GuestTeamRanking  int   `orm:"description(客场排名)"`      //客场排名
}

const (
	IntegralType_A = 1
	IntegralType_B = 2
)

func (this *FaIntegralRankingInfo) TableName() string {
	return mconst.TableName_FaIntegralRankingInfo
}

func (this *FaIntegralRankingInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaIntegralRankingInfo) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *FaIntegralRankingInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaIntegralRankingInfo) Update(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	this.UpdatedAt = time.Now()
	if cols != nil {
		cols = append(cols, "UpdatedAt")
	}

	_, e := o.Update(this, cols...)
	return e
}

func (this *FaIntegralRankingInfo) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaIntegralRankingInfo{}

	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).
		Filter("IntegralType", this.IntegralType).One(&data)
	if e == nil {
		this.Id = data.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}
