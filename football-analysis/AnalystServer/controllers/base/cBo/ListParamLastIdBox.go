package cBo

type ListParamLastIdBox struct {
	LastId int
	C      int
	Other  map[string]interface{}
}

func GetListParamLastIdBox() ListParamLastIdBox {
	aListParamBox := ListParamLastIdBox{}
	aListParamBox.Other = make(map[string]interface{})
	return aListParamBox
}
