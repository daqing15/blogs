package controllers

import (
	"blog/models"
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"regexp"
	"time"
)

type RegController struct {
	beego.Controller
}

func (this *RegController) Get() {
	this.TplNames = "reg.tpl"
}

func (this *RegController) Post() {
	this.TplNames = "reg.tpl"
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")
	usererr := checkUsername(username)
	fmt.Println(usererr)
	if usererr == false {
		this.Data["UsernameErr"] = "Username error, Please to again"
		return
	}

	passerr := checkPassword(password)
	if passerr == false {
		this.Data["PasswordErr"] = "Password error, Please to again"
		return
	}

	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	now := beego.Date(time.Now(), "Y-m-d H:i:s")

	userInfo := models.GetUserInfo(username)

	if userInfo.Username == "" {
		var users models.User
		users.Username = username
		users.Password = newPass
		users.Created = now
		users.Last_logintime = now
		models.AddUser(users)

		//登录成功设置session
		sess := this.StartSession()
		sess.Set("uid", userInfo.Id)
		sess.Set("uname", userInfo.Username)
		this.Ctx.Redirect(302, "/")
	} else {
		this.Data["UsernameErr"] = "User already exists"
	}

}

func checkPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func checkUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}
