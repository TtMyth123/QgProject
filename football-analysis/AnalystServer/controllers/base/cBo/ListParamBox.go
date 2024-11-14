package cBo

type ListParamBox struct {
	PageIndex int
	PageSize  int
	MaxId     int64
	Other     map[string]interface{}
	OrderBy   string
}

func (this *ListParamBox) AddParam(p ArgsParam, value interface{}) {
	this.Other[p.PName] = value
}

func GetListParamBox() ListParamBox {
	aListParamBox := ListParamBox{}
	aListParamBox.Other = make(map[string]interface{})
	return aListParamBox
}
