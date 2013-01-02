package admin

import (
	"blog/models"
	"github.com/astaxie/beego"
	"strconv"
)

type DelBlogController struct {
	beego.Controller
}

func (this *DelBlogController) Prepare() {
	sess := this.StartSession()
	sess_uid := sess.Get("userid")
	sess_username := sess.Get("username")
	if sess_uid == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
	}
	this.Data["Username"] = sess_username
}

func (this *DelBlogController) Get() {
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/delblog.tpl"
	this.Ctx.Request.ParseForm()
	id, _ := strconv.Atoi(this.Ctx.Request.Form.Get(":id"))
	blogInfo := models.GetBlogInfoById(id)
	models.DelBlogById(blogInfo)
	this.Ctx.Redirect(302, "/admin/index")
}
