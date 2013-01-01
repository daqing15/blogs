package admin

import (
	"blog/models"
	"blog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"strconv"
	"time"
)

type EditBlogController struct {
	beego.Controller
}

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func (this *EditBlogController) Prepare() {
	sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	sess_uid := sess.Get("userid")
	sess_username := sess.Get("username")
	if sess_uid == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
	}
	this.Data["Username"] = sess_username
}

func (this *EditBlogController) Get() {
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/editblog.tpl"
	this.Ctx.Request.ParseForm()
	id, _ := strconv.Atoi(this.Ctx.Request.Form.Get(":id"))
	blogInfo := models.GetBlogInfoById(id)
	this.Data["Id"] = blogInfo.Id
	this.Data["Content"] = blogInfo.Content
	this.Data["Title"] = blogInfo.Title
}

func (this *EditBlogController) Post() {
	this.Ctx.Request.ParseForm()
	id_str := this.Ctx.Request.Form.Get("id")
	id, _ := strconv.Atoi(id_str)
	blogInfo := models.GetBlogInfoById(id)
	title := this.Ctx.Request.Form.Get("title")
	content := this.Ctx.Request.Form.Get("content")
	blogInfo.Title = title
	blogInfo.Content = content
	//打印生成日志
	defer utils.Info("editBlog: ", "id:"+id_str, "title:"+title, "content:"+content)
	//获取系统当前时间
	now := beego.Date(time.Now(), "Y-m-d H:i:s")
	blogInfo.Created = now
	models.UpdateBlogInfo(blogInfo)
	this.Ctx.Redirect(302, "/admin/index")
}
