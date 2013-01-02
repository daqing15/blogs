package admin

import (
	"blog/models"
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"io"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplNames = "admin/signin.tpl"
}

func (this *LoginController) Post() {
	//数据处理
	this.TplNames = "admin/signin.tpl"
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")

	if username == "" {
		this.Data["UsernameNull"] = "username is not null"
		return
	}

	if password == "" {
		this.Data["PasswordNull"] = "password is not null"
		return
	}

	adminInfo := models.GetAdminInfo(username)

	if adminInfo.Username == "" {
		this.Data["UsernameNull"] = "用户不存在"
		return
	}

	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	if adminInfo.Password != newPass {
		this.Data["PasswordNull"] = "密码错误"
		return
	}

	//登录成功设置session
	sess := this.StartSession()
	sess.Set("userid", adminInfo.Id)
	sess.Set("username", adminInfo.Username)

	this.Ctx.Redirect(302, "/admin/index")
}
