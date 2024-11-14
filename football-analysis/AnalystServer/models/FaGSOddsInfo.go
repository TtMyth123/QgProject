package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

type FaGSOddsInfo struct {
	//Id    int64
	//sysId int64
	BaseInfo
	CompanyName string  `orm:"size(200);null;description(赔率公司名)"`
	RaceInfoId  int64   `orm:"null;description(赛事ID)"`
	BeginOdds1  float64 `orm:"digits(8);decimals(3);null;description(初盘赔率1)"` //初始赔率1
	BeginOdds2  float64 `orm:"digits(8);decimals(3);null;description(初盘赔率2)"` //初始赔率2
	BeginOdds3  float64 `orm:"digits(8);decimals(3);null;description(初盘赔率3)"` //初始赔率3
	BeginCOdds2 string  `orm:"size(200);null;description(初盘亚赔率)"`
	EndOdds1    float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率1)"` //最后的赔率1
	EndOdds2    float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率2)"` //最后的赔率2
	EndOdds3    float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率3)"` //最后的赔率3
	EndCOdds2   string  `orm:"size(200);null;description(末盘亚赔率2)"`

	BeginOddsTime time.Time `orm:"null;description(初始时间)"`
	EndOddsTime   time.Time `orm:"null;description(末盘时间)"`
}

func (this *FaGSOddsInfo) TableName() string {
	return mconst.TableName_FaGSOddsInfo
}

func (this *FaGSOddsInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaGSOddsInfo) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *FaGSOddsInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaGSOddsInfo) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaGSOddsInfo) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaGSOddsInfo{}

	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).
		Filter("CompanyName", this.CompanyName).One(&data)
	if e == nil {
		this.Id = data.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}

	return e
}

func MultiSaveGSOddsInfo(o orm.Ormer, arr []FaGSOddsInfo, cols ...string) error {
	for _, item := range arr {
		e := item.AddUpdate(o, cols...)
		if e != nil {
			ttLog.LogDebug(e)
		}
	}

	return nil
}
