package models

//type FaRaceInfo struct {
//	FaRaceInfo
//	ARanking string `orm:"size(40);null;description(A队排名)"`
//	BRanking string `orm:"size(40);null;description(B队排名)"`
//}
//
//func (a *FaRaceInfo) TableName() string {
//	return mconst.TableName_FaRaceInfo
//}
//
//func (this *FaRaceInfo) Delete(o orm.Ormer) error {
//	if o == nil {
//		o = orm.NewOrm()
//	}
//	_, e := o.Delete(this)
//	return e
//}
//
//func (this *FaRaceInfo) Read(o orm.Ormer) error {
//	if o == nil {
//		o = orm.NewOrm()
//	}
//	e := o.Read(this)
//	return e
//}
//
//func (this *FaRaceInfo) ReadEx(o orm.Ormer) error {
//	if o == nil {
//		o = orm.NewOrm()
//	}
//	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).One(this)
//	return e
//}
//
//func (this *FaRaceInfo) Add(o orm.Ormer) error {
//	if o == nil {
//		o = orm.NewOrm()
//	}
//
//	this.CreatedAt = time.Now()
//	this.UpdatedAt = this.CreatedAt
//	id, e := o.Insert(this)
//	this.Id = id
//	return e
//}
//
//func (this *FaRaceInfo) Update(o orm.Ormer, cols ...string) error {
//	if o == nil {
//		o = orm.NewOrm()
//	}
//	this.UpdatedAt = time.Now()
//	if cols != nil {
//		cols = append(cols, "UpdatedAt")
//	}
//
//	_, e := o.Update(this, cols...)
//	return e
//}
//
//func (this *FaRaceInfo) AddUpdate(o orm.Ormer, cols ...string) error {
//	if o == nil {
//		o = orm.NewOrm()
//	}
//	data := FaRaceInfo{}
//
//	e := o.QueryTable(this.TableName()).Filter("SysId", this.SysId).Filter("RaceInfoId", this.RaceInfoId).One(&data)
//	if e == nil {
//		this.Id = data.Id
//		e = this.Update(o, cols...)
//	} else {
//		e = this.Add(o)
//	}
//	return e
//}
