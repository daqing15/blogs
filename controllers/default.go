package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"strconv"
)

type MainController struct {
	beego.Controller
}

var welcome = `
# Golang 技术博客

`

func (this *MainController) Get() {
	this.Ctx.Request.ParseForm()
	id, _ := strconv.Atoi(this.Ctx.Request.Form.Get(":id"))
	if id == 0 {
		id = 1
	}
	blogInfo := models.GetBlogInfoById(id)
	this.Data["BlogInfo"] = blogInfo
	bloglist := models.GetAllBlogList()
	pm := make([]map[string]interface{}, len(bloglist))
	for k, pk := range bloglist {
		m := make(map[string]interface{}, 2)
		m["Blog"] = pk
		if this.Ctx.Params[":title"] == pk.Title {
			m["Cru"] = true
		} else {
			m["Cru"] = false
		}
		pm[k] = m
	}
	this.Data["BlogList"] = pm
	if blogInfo.Id == 0 {
		this.Data["Content"] = welcome
	} else {
		this.Data["Content"] = blogInfo.Content
	}

	this.Data["Username"] = ""
	sess := this.StartSession()
	username := sess.Get("uname")

	if username != nil {
		this.Data["Username"] = username
	}

	this.TplNames = "index.tpl"
}
