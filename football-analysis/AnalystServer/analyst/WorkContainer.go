package analyst

import "sync"

type WorkContainer struct {
	basePath1 string
	basePath2 string
	basePath3 string
	mpWork    sync.Map
}

func NewWorkContainer() *WorkContainer {
	aWorkContainer := new(WorkContainer)

	return aWorkContainer
}

func (this *WorkContainer) GetWork(SysId int64) *AnalystWork {
	aWork, _ := this.mpWork.LoadOrStore(SysId, this.newWork(SysId))
	aAnalystWork := aWork.(*AnalystWork)
	return aAnalystWork
}

func (this *WorkContainer) newWork(SysId int64) *AnalystWork {
	aWork := NewAnalystWork(2, "./conf/A.xlsx", this.basePath1, this.basePath2, this.basePath3)

	return aWork
}
