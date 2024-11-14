package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
*
伤停数据
*/
type FaInjury struct {
	BaseInfo
	RaceInfoId int64  `orm:"description(赛事ID)"`
	BanCrew    string `orm:"default(0); description(队员名称)"`
	TeamType   int    `orm:"default(1); description(队伍类型，1:A队，2:B队.)"`
	BanType    int    `orm:"default(1); description(受伤类型。1：停，2：伤)"`
}

const (
	TeamType_1 = 1
	TeamType_2 = 2
)
const (
	BanType_1 = 1
	BanType_2 = 2
)

func (a *FaInjury) TableName() string {
	return mconst.TableName_FaInjury
}

func (this *FaInjury) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *FaInjury) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaInjury) ReadEx(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).One(this)
	return e
}

func (this *FaInjury) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaInjury) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaInjury) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaInjury{}

	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).
		Filter("TeamType", this.TeamType).Filter("BanType", this.BanType).One(&data)
	if e == nil {
		this.Id = data.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}

func MultiSaveInjury(o orm.Ormer, arr []FaInjury, cols ...string) error {
	for _, item := range arr {
		e := item.AddUpdate(o, cols...)
		if e != nil {
			ttLog.LogDebug(e)
		}
	}

	return nil
}
