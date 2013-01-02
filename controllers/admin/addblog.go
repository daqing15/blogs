package admin

import (
	"blog/models"
	"blog/utils"
	"github.com/astaxie/beego"
	"time"
	//"fmt"
)

type AddBlogController struct {
	beego.Controller
}

func (this *AddBlogController) Prepare() {
	sess := this.StartSession()
	sess_uid := sess.Get("userid")
	sess_username := sess.Get("username")
	if sess_uid == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
	}
	this.Data["Username"] = sess_username
}

func (this *AddBlogController) Get() {
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/addblog.tpl"
}

func (this *AddBlogController) Post() {
	this.Ctx.Request.ParseForm()
	title := this.Ctx.Request.Form.Get("title")
	content := this.Ctx.Request.Form.Get("content")
	//打印生成日志
	defer utils.Info("addblog: ", "title:"+title, "content:"+content)
	var data models.Blogs
	data.Title = title
	data.Content = content
	//获取系统当前时间
	now := beego.Date(time.Now(), "Y-m-d H:i:s")
	data.Created = now
	models.InsertBlogs(data)
	this.Ctx.Redirect(302, "/admin/index")
}
