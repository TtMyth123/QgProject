package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type FaCustomRankingInfo struct {
	BaseInfo
	TeamId      int64  `orm:"description(球队ID)"`
	RankingName string `orm:"size(200);description(排名)"`
}

func (this *FaCustomRankingInfo) TableName() string {
	return mconst.TableName_FaCustomRankingInfo
}

func (this *FaCustomRankingInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaCustomRankingInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaCustomRankingInfo) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaCustomRankingInfo) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaCustomRankingInfo{}

	o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("TeamId", this.TeamId).One(&data)
	e := data.Read(o)
	if e == nil {
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}
