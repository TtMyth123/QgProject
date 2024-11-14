package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/orm"
	"time"
)

type FaExcelFData struct {
	BaseInfo
	RaceInfoId int64   `orm:"description(赛事ID)"`
	DataType   int     `orm:"description(数据类型.0:完整。 1:全部。2：1小时。3：权威)"`
	AStrength  string  `orm:"size(200);null;description(A实力亚盘)"` //A实力亚盘
	AStrength0 float64 `orm:"digits(12);decimals(3);description(A实力亚盘0)"`
	AStrength1 float64 `orm:"digits(12);decimals(3);description(A实力亚盘1)"`

	AAttack float64 `orm:"digits(12);decimals(3);description(A进攻分值)"`
	ADefend float64 `orm:"digits(12);decimals(3);description(A防守分值)"` //
	ADValue string  `orm:"size(200);null;description(主队差)"`           //主队差

	BStrength  string  `orm:"size(200);null;description(B实力亚盘)"` //B实力亚盘
	BStrength0 float64 `orm:"digits(12);decimals(3);description(B实力亚盘0)"`
	BStrength1 float64 `orm:"digits(12);decimals(3);description(B实力亚盘1)"`

	BAttack float64 `orm:"digits(12);decimals(3);description(B进攻分值)"` //B进攻分值
	BDefend float64 `orm:"digits(12);decimals(3);description(B防守分值)"` //B防守分值
	BDValue string  //客队差

	DValue string `orm:"size(200);null;description(差距)"` //差距

	GoalInGap  string  `orm:"size(200);null;description(一位进球差距)"`
	GoalInGap0 float64 `orm:"digits(12);decimals(3);description(一位进球差距0)"`
	GoalInGap1 float64 `orm:"digits(12);decimals(3);description(一位进球差距1)"`

	ColligateGap  string  `orm:"size(200);null;description(一位综合差距)"`
	ColligateGap0 float64 `orm:"digits(12);decimals(3);description(一位综合差距0)"`
	ColligateGap1 float64 `orm:"digits(12);decimals(3);description(一位综合差距1)"`

	WheelForce  string  `orm:"size(200);null;description(盘力指数)"` //盘力指数
	WheelForce0 float64 `orm:"digits(12);decimals(3);description(盘力指数0)"`
	WheelForce1 float64 `orm:"digits(12);decimals(3);description(盘力指数1)"`

	WinIndex  string  `orm:"size(200);null;description(取胜易度)"` //取胜易度
	WinIndex0 float64 `orm:"digits(12);decimals(3);description(取胜易度0)"`
	WinIndex1 float64 `orm:"digits(12);decimals(3);description(取胜易度1)"`

	Skewness1 float64 `orm:"digits(12);decimals(3);description(初盘偏度1)"` //初盘偏度1
	Skewness2 float64 `orm:"digits(12);decimals(3);description(初盘偏度2)"` //初盘偏度2
	Skewness3 float64 `orm:"digits(12);decimals(3);description(初盘偏度3)"` //初盘偏度3

	InitialKurtosis1 float64 `orm:"digits(12);decimals(3);description(初盘峰度1)"` //初盘峰度1
	InitialKurtosis2 float64 `orm:"digits(12);decimals(3);description(初盘峰度2)"` //初盘峰度2
	InitialKurtosis3 float64 `orm:"digits(12);decimals(3);description(初盘峰度3)"` //初盘峰度3

	LateKurtosis1 float64 `orm:"digits(12);decimals(3);description(尾盘峰度1)"` //尾盘峰度1
	LateKurtosis2 float64 `orm:"digits(12);decimals(3);description(尾盘峰度2)"` //尾盘峰度2
	LateKurtosis3 float64 `orm:"digits(12);decimals(3);description(尾盘峰度3)"` //尾盘峰度3

	MaxGL1 float64 `orm:"digits(12);decimals(3);description(最高凯利1)"` //最高凯利1
	MaxGL2 float64 `orm:"digits(12);decimals(3);description(最高凯利2)"` //最高凯利2
	MaxGL3 float64 `orm:"digits(12);decimals(3);description(最高凯利3)"` //最高凯利3

	UnitOffset1 float64 `orm:"digits(12);decimals(3);description(单位偏差1)"` //单位偏差1
	UnitOffset2 float64 `orm:"digits(12);decimals(3);description(单位偏差2)"` //单位偏差2
	UnitOffset3 float64 `orm:"digits(12);decimals(3);description(单位偏差3)"` //单位偏差3

	GLOffset1 float64 `orm:"digits(12);decimals(3);description(凯利偏差1)"` //凯利偏差1
	GLOffset2 float64 `orm:"digits(12);decimals(3);description(凯利偏差2)"` //凯利偏差2
	GLOffset3 float64 `orm:"digits(12);decimals(3);description(凯利偏差3)"` //凯利偏差3

	HistryACount int //A队历史数量
	HistryBCount int //B队历史数量

	S2_A52 int //A主场数
	S2_A53 int //A客场数

	S2_A64 int //B主场数
	S2_A65 int //B客场数

	SkipACount int
	SkipBCount int
}

// `orm:"description(数据类型.0:完整。 1：1小时。2：权威.3:全部。)""`
const (
	ExcelFData_DataType_Whole = 0
	ExcelFData_DataType_One   = 1
	ExcelFData_DataType_Auth  = 2
	ExcelFData_DataType_All   = 3
)

func (this *FaExcelFData) TableName() string {
	return mconst.TableName_FaExcelFData
}

func (this *FaExcelFData) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaExcelFData) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *FaExcelFData) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaExcelFData) Update(o orm.Ormer, cols ...string) error {
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

func (this *FaExcelFData) AddUpdate(o orm.Ormer, cols ...string) error {
	if o == nil {
		o = orm.NewOrm()
	}
	data := FaExcelFData{}

	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).Filter("DataType", this.DataType).One(&data)
	if e == nil {
		this.Id = data.Id
		e = this.Update(o, cols...)
	} else {
		e = this.Add(o)
	}
	return e
}
