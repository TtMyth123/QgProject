package cBo

type ParamType int

const (
	ParamTypeS ParamType = 1
	ParamTypeI ParamType = 2
)

type ArgsParam struct {
	PName    string
	T        ParamType
	DefValue interface{}
}
