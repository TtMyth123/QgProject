package base

import (
	"encoding/json"
	"fmt"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/controllers/base/TtSession"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/controllers/base/cBo"
	"github.com/TtMyth123/QgProject/football-analysis/AnalystServer/controllers/base/enums"
	"github.com/TtMyth123/kit/httpKit"
	"github.com/TtMyth123/kit/stringKit"
	"github.com/TtMyth123/kit/ttLog"
	"github.com/astaxie/beego"
)

var (
	isDev       = false
	mapUserInfo map[int64]BaseSysUser
	mapSId2Id   map[string]*TtSession.SUserInfo
)

func init() {
	mapUserInfo = make(map[int64]BaseSysUser)
	mapSId2Id = make(map[string]*TtSession.SUserInfo)
}

type BaseController struct {
	beego.Controller
}
type BaseSysUser struct {
	UserId int64
	Name   string
	SId    string
}

func (this *BaseController) LoginOK(id int64, sId string, life int) *TtSession.SUserInfo {
	aUser := TtSession.NewSUserInfo(id, sId, life)
	mapSId2Id[aUser.SId] = aUser
	return aUser
}

func (this *BaseController) LogoutOK(sId string) error {
	aUser := mapSId2Id[sId]
	if aUser != nil {
		delete(mapUserInfo, aUser.Id)
	}

	delete(mapSId2Id, sId)

	return nil
}

func (this *BaseController) GetSUserInfo(sId string) *TtSession.SUserInfo {
	return mapSId2Id[sId]
}

func (this *BaseController) JsonResult(code enums.JsonResultCode, msg string, obj interface{}, mp interface{}) {
	if isDev {
		if code != enums.JRCodeSucc {
			_, Action := this.GetControllerAndAction()
			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
		}
	}
	//a := cBo.ArgsBox{}
	//this.GetJsonData(&a)
	res := JsonResult{Code: code, Msg: msg, Obj: obj, Mp: mp}
	//aTmp := cBo.ReplyBox{Result: true, Data: obj}
	//if code != enums.JRCodeSucc {
	//	aTmp.Result = false
	//	aTmp.ErrMsg = msg
	//}
	//res.Obj = obj
	this.Data["json"] = res
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) JsonResultEx(obj interface{}, mp interface{}, err error) {
	if err == nil {
		this.JsonResult(enums.JRCodeSucc, "", obj, mp)
	} else {
		this.JsonResult(enums.JRCodeFailed, err.Error(), obj, mp)
	}
}

//func (this *BaseController) JsonListResult(code enums.JsonResultCode, msg string, LastId int, obj interface{}) {
//	if isDev {
//		if code != enums.JRCodeSucc {
//			_, Action := this.GetControllerAndAction()
//			msg = fmt.Sprintf("M:%s +\n %s", Action, msg)
//		}
//	}
//
//	a := cBo.ArgsBox{}
//	this.GetJsonData(&a)
//	res := JsonResult{Code: code, Msg: msg, Obj: obj,  }
//	//aTmp := cBo.ReplyListBox{Result: true, ListData: obj, LastId: LastId}
//	//if code != enums.JRCodeSucc {
//	//	aTmp.Result = false
//	//	aTmp.ErrMsg = msg
//	//}
//	//res.Obj = aTmp
//	this.Data["json"] = res
//	this.ServeJSON()
//	this.StopRun()
//}

func (this *BaseController) GetJsonData(v interface{}) error {
	strJson := this.GetString("jsonData")
	controllerName, actionName := this.GetControllerAndAction()
	ttLog.LogDebug(controllerName, actionName, "jsonData:", strJson)
	e := json.Unmarshal([]byte(strJson), v)
	ttLog.LogDebug("Data:", stringKit.GetJsonStr(v))
	return e
}

func (this *BaseController) GetListParam(otherPs ...cBo.ArgsParam) cBo.ListParamBox {
	data := httpKit.GetParamValue(this.Ctx.Request)
	fmt.Println(data)
	param := cBo.GetListParamBox()
	param.PageIndex, _ = this.GetInt("PageIndex")
	param.PageSize, _ = this.GetInt("PageSize")
	param.MaxId, _ = this.GetInt64("MaxId")
	param.OrderBy = this.GetString("OrderBy")
	if param.PageSize == 0 {
		param.PageSize = 10
	}

	if param.PageIndex == 0 {
		param.PageIndex = 1
	}
	for _, p := range otherPs {
		switch p.T {
		case cBo.ParamTypeS:
			if def, ok := p.DefValue.(string); ok {
				param.Other[p.PName] = this.GetString(p.PName, def)
			} else {
				param.Other[p.PName] = this.GetString(p.PName)
			}
		case p.T:
			if def, ok := p.DefValue.(int); ok {
				param.Other[p.PName], _ = this.GetInt(p.PName, def)
			} else {
				param.Other[p.PName], _ = this.GetInt(p.PName)
			}
		}
	}
	return param
}

func (this *BaseController) GetListParamLastId(otherPs ...cBo.ArgsParam) cBo.ListParamLastIdBox {
	data := httpKit.GetParamValue(this.Ctx.Request)
	fmt.Println(data)
	param := cBo.GetListParamLastIdBox()
	param.LastId, _ = this.GetInt("LastId")
	param.C, _ = this.GetInt("C")
	if param.C == 0 {
		param.C = 10
	}

	for _, p := range otherPs {
		switch p.T {
		case cBo.ParamTypeS:
			if def, ok := p.DefValue.(string); ok {
				param.Other[p.PName] = this.GetString(p.PName, def)
			} else {
				param.Other[p.PName] = this.GetString(p.PName)
			}
		case p.T:
			if def, ok := p.DefValue.(int); ok {
				param.Other[p.PName], _ = this.GetInt(p.PName, def)
			} else {
				param.Other[p.PName], _ = this.GetInt(p.PName)
			}
		}
	}
	return param
}
