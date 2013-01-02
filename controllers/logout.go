package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	this.TplNames = "logout.tpl"
	//退出，销毁session
	sess := this.StartSession()
	sess.Delete("uid")
	sess.Delete("uname")

	this.Ctx.Redirect(302, "/")
}
