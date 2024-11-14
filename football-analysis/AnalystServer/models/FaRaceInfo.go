package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type FaRaceInfo struct {
	BaseInfo
	//FaRaceInfo
	IsJs    int `orm:"description(是否及时数据)"`
	IsJsTmp int `orm:"description(是否及时数据_临时)"`
	IsHc    int `orm:"description(是否回查数据)"`
	IsHcTmp int `orm:"description(是否回查数据_临时)"`

	IsShow    int    `orm:"default(0); description(是否已显示过数据)"`
	IsGet     int    `orm:"default(0); description(是否已显示过数据)"`
	HomeTeam  string `orm:"size(200);description(主队名称)"`
	GuestTeam string `orm:"size(200);description(客队名称)"`
	Html800   string `orm:"size(2000);description(网页html)"`

	LeagueName      string    `orm:"size(200);description(联赛名称)"`
	RaceInfoId      int64     `orm:"description(赛事ID)"`
	RaceTime        time.Time `orm:"null;description(比赛时间)"`
	ATeamId         int64     `orm:"description(A球队ID)"`
	ATeamEName      string    `orm:"size(200);description(B球队E名称)"`
	ATeamCName      string    `orm:"size(200);description(B球队C名称)"`
	BTeamId         int64     `orm:"description(B球队ID)"`
	BTeamEName      string    `orm:"size(200);description(B球队E名称)"`
	BTeamCName      string    `orm:"size(200);description(B球队C名称)"`
	AFirstHalfScore int       `orm:"description(半场A进球)"`
	BFirstHalfScore int       `orm:"description(半场B进球)"`
	AScore          int       `orm:"description(A进球)"`
	BScore          int       `orm:"description(B进球)"`
	EndCOdds2       string    `orm:"description(盘口_终盘)"`
	EndEOdds2       float64   `orm:"digits(12);decimals(3);description(盘口_数值_终盘)"`
	BeginCOdds2     string    `orm:"description(盘口_初盘)"`
	BeginEOdds2     float64   `orm:"digits(12);decimals(3);description(盘口_数值_初盘)"`
	AvgFirstWet     float64   `orm:"digits(12);decimals(3);description(初水)"`
	StrengthBValue  float64   `orm:"digits(12);decimals(3);description(B实力)"`
	BeginValue      float64   `orm:"digits(12);decimals(3);description(第一个初盘)"`
	EndValue        float64   `orm:"digits(12);decimals(3);description(第一个终盘)"`
	WinValue        float64   `orm:"digits(12);decimals(3);description(取胜)"`
	ZHValue         float64   `orm:"digits(12);decimals(3);description(综合)"`

	ARanking string `orm:"size(40);null;description(A队排名)"`
	BRanking string `orm:"size(40);null;description(B队排名)"`
}

func (this *FaRaceInfo) TableName() string {
	return mconst.TableName_FaRaceInfo
}

func (this *FaRaceInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaRaceInfo) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *FaRaceInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaRaceInfo) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaRaceInfo) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaRaceInfo{}

	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).One(&data)
	if e == nil {
		this.Id = data.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}
