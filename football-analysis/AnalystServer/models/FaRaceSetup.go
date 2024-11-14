package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type FaRaceSetup struct {
	BaseInfo
	RaceInfoId int64 `orm:"description(赛事ID)"`

	RA1     int `orm:"description(A红1。值:1，0)"`
	RA2     int `orm:"description(A红2。值:1，0)"`
	GA1     int `orm:"description(A绿1。值:1，0)"`
	GA2     int `orm:"description(A绿2。值:1，0)"`
	BA      int `orm:"description(A黑1。值:1，0)"`
	BLA     int `orm:"description(A蓝1。值:1，0)"`
	FHA     int `orm:"description(A粉1。值:1，0)"`
	TeamAPm int `orm:"description(不知)"`

	RB1     int `orm:"description(B红1。值:1，0)"`
	RB2     int `orm:"description(B红2。值:1，0)"`
	GB1     int `orm:"description(B绿1。值:1，0)"`
	GB2     int `orm:"description(B绿2。值:1，0)"`
	BB      int `orm:"description(B黑1。值:1，0)"`
	FHB     int `orm:"description(B蓝1。值:1，0)"`
	BLB     int `orm:"description(B粉1。值:1，0)"`
	TeamBPm int `orm:"description(不知)"`

	RaceLc int `orm:"description(轮次)"`

	Info1 int `orm:"description(不知)"`
	Info2 int `orm:"description(不知)"`
	Info3 int `orm:"description(不知)"`
	Info4 int `orm:"description(不知)"`
}

func (this *FaRaceSetup) TableName() string {
	return mconst.TableName_FaRaceSetup
}

func (this *FaRaceSetup) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}
func (this *FaRaceSetup) ReadEx(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).One(this)
	return e
}

func (this *FaRaceSetup) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaRaceSetup) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaRaceSetup) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaRaceSetup{}

	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).One(&data)
	if e == nil {
		this.Id = data.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}
