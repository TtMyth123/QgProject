package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
*
亚赔数据
*/
type FaHistoryAsiaOdds struct {
	BaseInfo
	RaceInfoId  int64   `orm:"description(赛事ID)"`                               //赛事ID
	CompanyId   int     `orm:"description(赔率公司ID)"`                             //赔率公司ID
	BeginOdds1  float64 `orm:"digits(8);decimals(3);null;description(初盘赔率1)"`   //初始赔率1
	BeginOdds2  float64 `orm:"digits(8);decimals(3);null;description(初盘赔率2)"`   //初始赔率2
	BeginOdds3  float64 `orm:"digits(8);decimals(3);null;description(初盘赔率3)"`   //初始赔率3
	EndOdds1    float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率1)"` //最后的赔率1
	EndOdds2    float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率2)"` //最后的赔率2
	EndOdds3    float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率3)"` //最后的赔率3
	BeginCOdds2 string  `orm:"size(200);null;description(初盘赔率2)"`
	EndCOdds2   string  `orm:"size(200);null;description(末盘亚赔率2)"`
}

func (this *FaHistoryAsiaOdds) TableName() string {
	return mconst.TableName_FaHistoryAsiaOdds
}

func (this *FaHistoryAsiaOdds) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *FaHistoryAsiaOdds) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaHistoryAsiaOdds) ReadEx(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).One(this)
	return e
}

func (this *FaHistoryAsiaOdds) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaHistoryAsiaOdds) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaHistoryAsiaOdds) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaHistoryAsiaOdds{}

	o.QueryTable(this.TableName()).Filter("SysId", this.SysId).
		Filter("RaceInfoId", this.RaceInfoId).Filter("CompanyId", this.CompanyId).One(&data)
	e := data.Read(o)
	if e == nil {
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}

func MultiSaveHistoryAsiaOdds(o orm.Ormer, arr []FaHistoryAsiaOdds, cols ...string) error {
	for _, item := range arr {
		e := item.AddUpdate(o, cols...)
		if e != nil {
			ttLog.LogDebug(e)
		}
	}

	return nil
}
