package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
*
欧赔计算数据

0主 1和 2客 3主胜 4和胜 5客胜 6价值值(1) 7价值值(2) 8价值值(3) 9总和 10反总和 11和 12佘 13反佘 14差 15返回率 16记录数 17和客胜 18价值1 19主和胜  20价值3 21余差  22反余差  23总余  24反总余 25值1  26值2  27值3 28主1  29和1  30客1 31余C 32反余C
*/
type FaEuropeOddCalculate struct {
	BaseInfo
	RaceInfoId int64 `orm:"description(赛事ID)"`
	DataType   int   `orm:"description(数据类型.1:一小时。2：权威_标准。3：平均_全部，4：第一天，5：权威_E,6:选择)"`
	C          int   `orm:"description(参与计算的赔率公司数量)"`

	BeginOdds1 float64 `orm:"digits(12);decimals(3);description(初亚赔_主)"`
	BeginOdds2 float64 `orm:"digits(12);decimals(3);description(初亚赔_和)"`
	BeginOdds3 float64 `orm:"digits(12);decimals(3);description(初亚赔_客)"`

	Begin00 float64 `orm:"digits(12);decimals(3);description(初赔_0主)"`
	Begin01 float64 `orm:"digits(12);decimals(3);description(初赔_1和)"`
	Begin02 float64 `orm:"digits(12);decimals(3);description(初赔_2客)"`
	Begin03 float64 `orm:"digits(12);decimals(3);description(初赔_3主胜)"`
	Begin04 float64 `orm:"digits(12);decimals(3);description(初赔_4和胜)"`
	Begin05 float64 `orm:"digits(12);decimals(3);description(初赔_5客胜)"`
	Begin06 float64 `orm:"digits(12);decimals(3);description(初赔_6价值值1)"`
	Begin07 float64 `orm:"digits(12);decimals(3);description(初赔_7价值值2)"`
	Begin08 float64 `orm:"digits(12);decimals(3);description(初赔_8价值值3)"`
	Begin09 float64 `orm:"digits(12);decimals(3);description(初赔_9总和)"`
	Begin10 float64 `orm:"digits(12);decimals(3);description(初赔_10反总和)"`
	Begin11 float64 `orm:"digits(12);decimals(3);description(初赔_11和)"`
	Begin12 float64 `orm:"digits(12);decimals(3);description(初赔_12佘)"`
	Begin13 float64 `orm:"digits(12);decimals(3);description(初赔_13反佘)"`
	Begin14 float64 `orm:"digits(12);decimals(3);description(初赔_14差)"`
	Begin15 float64 `orm:"digits(12);decimals(3);description(初赔_15返回率)"`
	Begin16 int     `orm:"digits(12);decimals(3);description(初赔_16记录数)"`
	Begin17 float64 `orm:"digits(12);decimals(3);description(初赔_17和客胜)"`
	Begin18 float64 `orm:"digits(12);decimals(3);description(初赔_18价值1)"`
	Begin19 float64 `orm:"digits(12);decimals(3);description(初赔_19主和胜)"`
	Begin20 float64 `orm:"digits(12);decimals(3);description(初赔_20价值3)"`
	Begin21 float64 `orm:"digits(12);decimals(3);description(初赔_21余差)"`
	Begin22 float64 `orm:"digits(12);decimals(3);description(初赔_22反余差)"`
	Begin23 float64 `orm:"digits(12);decimals(3);description(初赔_23总余)"`
	Begin24 float64 `orm:"digits(12);decimals(3);description(初赔_24反总余)"`
	Begin25 float64 `orm:"digits(12);decimals(3);description(初赔_25值1)"`
	Begin26 float64 `orm:"digits(12);decimals(3);description(初赔_26值2)"`
	Begin27 float64 `orm:"digits(12);decimals(3);description(初赔_27值3)"`
	Begin28 float64 `orm:"digits(12);decimals(3);description(初赔_28主1)"`
	Begin29 float64 `orm:"digits(12);decimals(3);description(初赔_29和1)"`
	Begin30 float64 `orm:"digits(12);decimals(3);description(初赔_30客1)"`
	Begin31 float64 `orm:"digits(12);decimals(3);description(初赔_31余C)"`
	Begin32 float64 `orm:"digits(12);decimals(3);description(初赔_32反余C)"`

	EndOdds1 float64 `orm:"digits(12);decimals(3);description(末亚赔_主)"`
	EndOdds2 float64 `orm:"digits(12);decimals(3);description(末亚赔_和)"`
	EndOdds3 float64 `orm:"digits(12);decimals(3);description(末亚赔_客)"`

	End00 float64 `orm:"digits(12);decimals(3);description(末赔_0主)"`
	End01 float64 `orm:"digits(12);decimals(3);description(末赔_1和)"`
	End02 float64 `orm:"digits(12);decimals(3);description(末赔_2客)"`
	End03 float64 `orm:"digits(12);decimals(3);description(末赔_3主胜)"`
	End04 float64 `orm:"digits(12);decimals(3);description(末赔_4和胜)"`
	End05 float64 `orm:"digits(12);decimals(3);description(末赔_5客胜)"`
	End06 float64 `orm:"digits(12);decimals(3);description(末赔_6价值值1)"`
	End07 float64 `orm:"digits(12);decimals(3);description(末赔_7价值值2)"`
	End08 float64 `orm:"digits(12);decimals(3);description(末赔_8价值值3)"`
	End09 float64 `orm:"digits(12);decimals(3);description(末赔_9总和)"`
	End10 float64 `orm:"digits(12);decimals(3);description(末赔_10反总和)"`
	End11 float64 `orm:"digits(12);decimals(3);description(末赔_11和)"`
	End12 float64 `orm:"digits(12);decimals(3);description(末赔_12佘)"`
	End13 float64 `orm:"digits(12);decimals(3);description(末赔_13反佘)"`
	End14 float64 `orm:"digits(12);decimals(3);description(末赔_14差)"`
	End15 float64 `orm:"digits(12);decimals(3);description(末赔_15返回率)"`
	End16 int     `orm:"digits(12);decimals(3);description(末赔_16记录数)"`
	End17 float64 `orm:"digits(12);decimals(3);description(末赔_17和客胜)"`
	End18 float64 `orm:"digits(12);decimals(3);description(末赔_18价值1)"`
	End19 float64 `orm:"digits(12);decimals(3);description(末赔_19主和胜)"`
	End20 float64 `orm:"digits(12);decimals(3);description(末赔_20价值3)"`
	End21 float64 `orm:"digits(12);decimals(3);description(末赔_21余差)"`
	End22 float64 `orm:"digits(12);decimals(3);description(末赔_22反余差)"`
	End23 float64 `orm:"digits(12);decimals(3);description(末赔_23总余)"`
	End24 float64 `orm:"digits(12);decimals(3);description(末赔_24反总余)"`
	End25 float64 `orm:"digits(12);decimals(3);description(末赔_25值1)"`
	End26 float64 `orm:"digits(12);decimals(3);description(末赔_26值2)"`
	End27 float64 `orm:"digits(12);decimals(3);description(末赔_27值3)"`
	End28 float64 `orm:"digits(12);decimals(3);description(末赔_28主1)"`
	End29 float64 `orm:"digits(12);decimals(3);description(末赔_29和1)"`
	End30 float64 `orm:"digits(12);decimals(3);description(末赔_30客1)"`
	End31 float64 `orm:"digits(12);decimals(3);description(末赔_31余C)"`
	End32 float64 `orm:"digits(12);decimals(3);description(末赔_32反余C)"`
}

const (
	EuropeOddCalculateType_OneHour   = 1
	EuropeOddCalculateType_Authority = 2
	EuropeOddCalculateType_Avg       = 3
)

func (this *FaEuropeOddCalculate) TableName() string {
	return mconst.TableName_FaEuropeOddCalculate
}

func (this *FaEuropeOddCalculate) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaEuropeOddCalculate) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaEuropeOddCalculate) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaEuropeOddCalculate) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaEuropeOddCalculate{}

	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).One(&data)
	if e == nil {
		this.Id = data.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}
