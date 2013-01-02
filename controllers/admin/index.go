package admin

import (
	"blog/models"
	"blog/utils"
	"github.com/astaxie/beego"
	"strconv"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Prepare() {
	sess := this.StartSession()

	sess_uid := sess.Get("userid")
	sess_username := sess.Get("username")
	if sess_uid == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
	}

	this.Data["Username"] = sess_username
}

func (this *IndexController) Get() {
	this.Ctx.Request.ParseForm()

	page, _ := strconv.Atoi(this.Ctx.Request.Form.Get("page"))
	offset := 1
	if page == 0 {
		page = 1
	}

	start := (page - 1) * offset
	list := models.GetBlogsList(start, offset)
	countInfo, _ := models.GetAllBlogsCount()
	count := string(countInfo[0]["count(*)"])
	totalCount, _ := strconv.Atoi(count)
	newList := make([]map[string]interface{}, len(list))
	for k, v := range list {
		m := make(map[string]interface{})
		m["id"] = v.Id
		m["title"] = v.Title
		m["content"] = v.Content
		m["created"] = v.Created
		newList[k] = m
	}

	this.Data["Blogs"] = newList
	//分页配置
	conf := utils.Config{
		PageUrl:       "/admin/index",
		PageSize:      1,
		RowsNum:       totalCount,
		AnchorClass:   "",
		LinksNum:      1,
		PageNum:       page,
		CurrentClass:  "",
		InfoTagOpen:   "<li>",
		InfoTagClose:  "</li>",
		FirstTagOpen:  "<li>",
		FirstTagClose: "</li>",
		LastTagOpen:   "<li>",
		LastTagClose:  "</li>",
		CurTagOpen:    "<li>",
		CurTagClose:   "</li>",
		NextTagOpen:   "<li>",
		NextTagClose:  "</li>",
		PrevTagOpen:   "<li>",
		PrevTagClose:  "</li>",
		NumTagOpen:    "<li>",
		NumTagClose:   "</li>",
	}

	pageStr, err := utils.CreateLinks(conf)
	if err != nil {
		beego.BeeLogger.Fatal(err)
	}

	if pageStr == "404" {
		this.Ctx.Redirect(302, "/admin/index")
	}
	this.Data["PageStr"] = "<ul>" + pageStr + "</ul>"
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/index.tpl"
}
