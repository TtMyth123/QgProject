package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
*
球队信息
*/
type FaTeamInfo struct {
	BaseInfo
	TeamEName string `orm:"size(200);description(球队名称英文)"`
	TeamCName string `orm:"size(200);description(球队名称中文)"`
}

func (this *FaTeamInfo) TableName() string {
	return mconst.TableName_FaTeamInfo
}

func (this *FaTeamInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaTeamInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaTeamInfo) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaTeamInfo) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}

	data := &FaTeamInfo{BaseInfo: BaseInfo{Id: this.Id}}
	e := data.Read(o)
	if e == nil {
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}
