package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	FutureType_A = 1 //A队
	FutureType_B = 2 //B队
)

/*
*
未来赛事数据
*/
type FaFutureRaceInfo struct {
	BaseInfo
	RaceInfoId int64  `orm:"description(赛事ID)"`
	FutureType int    `orm:" ;description(类型. 1:A队.2:B队)"`                  //初始赔率1
	LeagueName string `orm:"size(200);null;description(联赛名称)"`              //联赛名称
	ClashName  string `orm:"size(400);null;description(对阵,如：名古屋鲸八 - 大阪钢巴)"` //对阵,如：名古屋鲸八 - 大阪钢巴
}

func (a *FaFutureRaceInfo) TableName() string {
	return mconst.TableName_FaFutureRaceInfo
}

func (this *FaFutureRaceInfo) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *FaFutureRaceInfo) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaFutureRaceInfo) ReadEx(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).One(this)
	return e
}

func (this *FaFutureRaceInfo) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaFutureRaceInfo) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaFutureRaceInfo) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaFutureRaceInfo{}
	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).
		Filter("FutureType", this.FutureType).One(&data)
	if e == nil {
		this.Id = data.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}

func MultiSaveFutureRaceInfo(o orm.Ormer, arr []FaFutureRaceInfo, cols ...string) error {
	for _, item := range arr {
		e := item.AddUpdate(o, cols...)
		if e != nil {
			ttLog.LogDebug(e)
		}
	}

	return nil
}
