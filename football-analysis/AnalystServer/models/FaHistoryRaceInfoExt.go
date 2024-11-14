package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	HistoryType_A = 1
	HistoryType_B = 2
	HistoryType_V = 3
)

type FaHistoryRaceInfoExt struct {
	BaseInfo
	MainRaceInfoId int64     `orm:"description(主表赛事ID)"`
	HistoryType    int       `orm:"default(1); description(历史对战类型，1: A v N, 2:B v N, 3:A v B)"`
	RaceInfoId     int64     `orm:"description(赛事ID)"` //赛事ID
	RaceTime       time.Time //比赛时间
	LeagueName     string    `orm:"size(200);null;description(联赛名)"`  //联赛名
	ATeamId        int64     `orm:"description(A队ID)"`                //A队ID
	ATeamName      string    `orm:"size(200);null;description(A队名称)"` //A队名称
	ATeamRanking   string    `orm:"size(20);null;description(A队排名)"`  //A队排名
	BTeamId        int64     `orm:"description(B队ID)"`                //A队ID //B队ID
	BTeamName      string    `orm:"size(200);null;description(B队名称)"` //B队名称
	BTeamRanking   string    `orm:"size(20);null;description(B队排名)"`  //B队排名
	AScore         int       `orm:"description(A队全场分数)"`              //A队全场分数
	BScore         int       `orm:"description(B队全场分数)"`              //B队全场分数
	AHalfScore     int       `orm:"description(A队半场分数)"`              //A队半场分数
	BHalfScore     int       `orm:"description(B队半场分数)"`              //B队半场分数

	BeginAsiaOdds1  float64 `orm:"digits(8);decimals(3);null;description(初盘赔率1)"`   //初始赔率1
	BeginAsiaOdds2  float64 `orm:"digits(8);decimals(3);null;description(初盘赔率2)"`   //初始赔率2
	BeginAsiaOdds3  float64 `orm:"digits(8);decimals(3);null;description(初盘赔率3)"`   //初始赔率3
	EndAsiaOdds1    float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率1)"` //最后的赔率1
	EndAsiaOdds2    float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率2)"` //最后的赔率2
	EndAsiaOdds3    float64 `orm:"digits(8);decimals(3);null; description:(末盘赔率3)"` //最后的赔率3
	BeginCAsiaOdds2 string  `orm:"size(200);null;description(初盘赔率2)"`
	EndCAsiaOdds2   string  `orm:"size(200);null;description(末盘亚赔率2)"`

	BeginEuropeOdds1 float64 `orm:"digits(8);decimals(3);null;description(初始欧洲赔率1)"` //初始欧洲赔率1
	BeginEuropeOdds2 float64 `orm:"digits(8);decimals(3);null;description(初始欧洲赔率2)"` //初始欧洲赔率2
	BeginEuropeOdds3 float64 `orm:"digits(8);decimals(3);null;description(初始欧洲赔率3)"` //初始欧洲赔率3

	EndEuropeOdds1 float64 `orm:"digits(8);decimals(3);null;description(最后的欧洲赔率1)"` //最后的欧洲赔率1
	EndEuropeOdds2 float64 `orm:"digits(8);decimals(3);null;description(最后的欧洲赔率2)"` //最后的欧洲赔率2
	EndEuropeOdds3 float64 `orm:"digits(8);decimals(3);null;description(最后的欧洲赔率3)"` //最后的欧洲赔率3
}

func (this *FaHistoryRaceInfoExt) TableName() string {
	return mconst.TableName_FaHistoryRaceInfoExt
}

func (this *FaHistoryRaceInfoExt) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaHistoryRaceInfoExt) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaHistoryRaceInfoExt) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaHistoryRaceInfoExt) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaHistoryRaceInfoExt{}

	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).
		Filter("HistoryType", this.HistoryType).One(&data)
	if e == nil {
		this.Id = data.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}

func MultiSaveHistoryRaceInfoExt(o orm.Ormer, arr []FaHistoryRaceInfoExt, cols ...string) error {
	for _, item := range arr {
		e := item.AddUpdate(o, cols...)
		if e != nil {
			ttLog.LogDebug(e)
		}
	}

	return nil
}
