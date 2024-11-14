package models

//type FaRaceInfoHc struct {
//	FaRaceInfo
//	IsShow  int    `orm:"default(0); description(是否已显示过数据)"`
//	IsGet   int    `orm:"default(0); description(是否已显示过数据)"`
//	Html800 string `orm:"size(2000);description(网页html)"`
//}
//
//func (this *FaRaceInfoHc) TableName() string {
//	return mconst.TableName_FaRaceInfoHc
//}
//
//func (this *FaRaceInfoHc) Read(o orm.Ormer) error {
//	if o == nil {
//		o = orm.NewOrm()
//	}
//	e := o.Read(this)
//	return e
//}
//
//func (this *FaRaceInfoHc) Delete(o orm.Ormer) error {
//	if o == nil {
//		o = orm.NewOrm()
//	}
//	_, e := o.Delete(this)
//	return e
//}
//
//func (this *FaRaceInfoHc) Add(o orm.Ormer) error {
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
//func (this *FaRaceInfoHc) Update(o orm.Ormer, cols ...string) error {
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
//func (this *FaRaceInfoHc) AddUpdate(o orm.Ormer, cols ...string) error {
//	if o == nil {
//		o = orm.NewOrm()
//	}
//	data := FaRaceInfoHc{}
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
