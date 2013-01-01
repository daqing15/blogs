package main

import (
	"blog/controllers"
	"blog/controllers/admin"
	"github.com/astaxie/beego"
)

func main() {
	beego.PprofOn = true
	beego.RegisterController("/", &controllers.MainController{})
	beego.RegisterController("/reg", &controllers.RegController{})
	beego.RegisterController("/login", &controllers.LoginController{})
	beego.RegisterController("/logout", &controllers.LogoutController{})
	beego.RegisterController("/:id([0-9]+)", &controllers.MainController{})
	beego.RegisterController("/admin/index", &admin.IndexController{})
	beego.RegisterController("/admin/login", &admin.LoginController{})
	beego.RegisterController("/admin/addblog", &admin.AddBlogController{})
	beego.RegisterController("/admin/editblog/:id([0-9]+)", &admin.EditBlogController{})
	beego.RegisterController("/admin/delblog/:id([0-9]+)", &admin.DelBlogController{})
	beego.Run()
}
