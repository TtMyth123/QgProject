package models

import (
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/models/mconst"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

/*
*
伤停数据
*/
type FaAuthorityCompany struct {
	BaseInfo
	CompanyId int64
	CName     string `orm:"size(200);null;description(公司名称)"`
}

func (a *FaAuthorityCompany) TableName() string {
	return mconst.TableName_FaAuthorityCompany
}

func (this *FaAuthorityCompany) Delete(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	_, e := o.Delete(this)
	return e
}

func (this *FaAuthorityCompany) Read(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}
	e := o.Read(this)
	return e
}

func (this *FaAuthorityCompany) Add(o orm.Ormer) error {
	if o == nil {
		o = orm.NewOrm()
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = this.CreatedAt
	id, e := o.Insert(this)
	this.Id = id
	return e
}

func (this *FaAuthorityCompany) Update(o orm.Ormer, cols ...string) error {
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

func InitAuthorityCompany(o orm.Ormer) {
	if o == nil {
		o = orm.NewOrm()
	}
	c, _ := o.QueryTable(mconst.TableName_FaAuthorityCompany).Count()
	if c == 0 {
		arrData := make([]FaAuthorityCompany, 0)
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_01, CName: "SB"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_02, CName: "bet 365(英国)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_03, CName: "立博(英国)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_04, CName: "伟德(直布罗陀)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_05, CName: "易胜博(安提瓜和巴布达)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_06, CName: "明陞(菲律宾)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_07, CName: "澳彩"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_08, CName: "10BET(英国)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_09, CName: "金宝博(马恩岛)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_10, CName: "12BET(菲律宾)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_12, CName: "盈禾(菲律宾)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_13, CName: "18Bet"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_14, CName: "澳门"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_15, CName: "易胜博(安提瓜和巴布达)"})
		arrData = append(arrData, FaAuthorityCompany{CompanyId: mconst.CompanyId_16, CName: "12BET(菲律宾)"})
		_, e := o.InsertMulti(len(arrData), &arrData)
		if e != nil {
			logs.Error(e)
		}
	}
}
