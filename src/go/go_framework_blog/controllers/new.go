package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

/**
 * @author
 * @date 2020/11/11 15:36
 * create by Goland
 * @version 1.0
 */

type NewController struct {
	beego.Controller
}

func (this *NewController) Get() {
	this.Layout = "layout.tpl"
	this.TplName = "new.tpl"
}

func (this *NewController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now()
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}