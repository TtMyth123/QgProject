package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

type FaEuropeOddsInfo struct {
	BaseInfo
	//Id    int64
	//sysId int64
	CompanyName string `orm:"size(200);null;description(赔率公司名)"`
	RaceInfoId  int64  `orm:"null;description(赛事ID)"`

	CompanyId  int     `orm:"description(赔率公司ID)"`                             //赔率公司ID
	BeginOdds1 float64 `orm:"digits(8);decimals(3);null;description(初盘赔率1)"`   //初始赔率1
	BeginOdds2 float64 `orm:"digits(8);decimals(3);null;description(初盘赔率2)"`   //初始赔率2
	BeginOdds3 float64 `orm:"digits(8);decimals(3);null;description(初盘赔率3)"`   //初始赔率3
	EndOdds1   float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率1)"` //最后的赔率1
	EndOdds2   float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率2)"` //最后的赔率2
	EndOdds3   float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率3)"` //最后的赔率3

	BeginCOdds2 string `orm:"size(200);null;description(初盘亚赔率)"`
	EndCOdds2   string `orm:"size(200);null;description(末盘亚赔率2)"`

	KellyOdds1 float64 `orm:"digits(8);decimals(3);null; description:(凯利赔率1)"` //凯利赔率1
	KellyOdds2 float64 `orm:"digits(8);decimals(3);null; description:(凯利赔率2)"` //凯利赔率1
	KellyOdds3 float64 `orm:"digits(8);decimals(3);null; description:(凯利赔率3)"` //凯利赔率1

	BeginOddsTime time.Time `orm:"null;description(初始时间)"`
	EndOddsTime   time.Time `orm:"null;description(末盘时间)"`
	AsiaBOdds1    float64   `orm:"digits(8);decimals(3);null;description(转成亚初盘赔率1)"`
	AsiaBOdds3    float64   `orm:"digits(8);decimals(3);null;description(转成亚初盘赔率3)"`
	AsiaEOdds1    float64   `orm:"digits(8);decimals(3);null;description(转成亚最后赔率1)"`
	AsiaEOdds3    float64   `orm:"digits(8);decimals(3);null;description(转成亚最后赔率3)"`
	TypeZ         int       `orm:"null;description(0:不是权威，1：权威_主流)"`
	TypeE         int       `orm:"null;description(0:不是交易所，1：交易所)"`
}

func (this *FaEuropeOddsInfo) TableName() string {
	return mconst.TableName_FaEuropeOddsInfo
}

func (this *FaEuropeOddsInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaEuropeOddsInfo) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *FaEuropeOddsInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaEuropeOddsInfo) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaEuropeOddsInfo) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaEuropeOddsInfo{}

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

func MultiSaveEuropeOddsInfo(o orm.Ormer, arr []FaEuropeOddsInfo, cols ...string) error {
	for _, item := range arr {
		e := item.AddUpdate(o, cols...)
		if e != nil {
			ttLog.LogDebug(e)
		}
	}

	return nil
}
