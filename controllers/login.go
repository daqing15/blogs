package controllers

import (
	"blog/models"
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"io"
	"time"
)

type LoginController struct {
	beego.Controller
}

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func (this *LoginController) Post() {
	this.TplNames = "login.tpl"
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")
	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	now := beego.Date(time.Now(), "Y-m-d H:i:s")

	userInfo := models.GetUserInfo(username)
	if userInfo.Password == newPass {
		var users models.User
		users.Last_logintime = now
		models.UpdateUserInfo(users)

		//登录成功设置session
		sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		sess.Set("uid", userInfo.Id)
		sess.Set("uname", userInfo.Username)

		this.Ctx.Redirect(302, "/")
	}

	this.Ctx.Redirect(302, "/")
}
