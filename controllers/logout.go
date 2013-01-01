package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
)

type LogoutController struct {
	beego.Controller
}

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func (this *LogoutController) Get() {
	this.TplNames = "logout.tpl"
	//退出，销毁session
	sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	sess.Delete("uid")
	sess.Delete("uname")

	this.Ctx.Redirect(302, "/")
}
