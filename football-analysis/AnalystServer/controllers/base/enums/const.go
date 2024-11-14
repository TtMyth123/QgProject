package enums

type JsonResultCode int

const (
	JRCodeSucc   JsonResultCode = 200
	JRCodeFailed                = 500
	JRCode302                   = 302 //跳转至地址
	JRCode401                   = 401 //未授权访问
	JRCode402                   = 402 //权限不足
)
