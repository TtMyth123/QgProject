package base

type AuthorBaseController struct {
	BaseController
	SysId int64
	SId   string
}

func (this *AuthorBaseController) Prepare() {
	this.SysId = 1

	//strToken := this.Ctx.Input.Header("token")
	//aUser := this.GetSUserInfo(strToken)
	//if aUser==nil {
	//	this.JsonResultEx("",nil,errors.New("授权过期"))
	//}
	//this.SId = strToken
	//this.SysId = aUser.Id
}

func (this *AuthorBaseController) DoLogout() {
	this.LogoutOK(this.SId)
	this.JsonResultEx("", "", nil)
}
