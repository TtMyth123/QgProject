package cBo

type ReplyBox struct {
	Result  bool
	ErrMsg  string
	Data    interface{}
	ApiName string
}

type ReplyListBox struct {
	Result   bool
	ErrMsg   string
	ListData interface{}
	LastId   int
}
