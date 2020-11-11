package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

/**
 * @author
 * @date 2020/11/11 15:35
 * create by Goland
 * @version 1.0
 */

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplName = "edit.tpl"
}

func (this *EditController) Post() {
	inputs := this.Input()
	var blog models.Blog
	blog.Id, _ = strconv.Atoi(inputs.Get("id"))
	blog.Title = inputs.Get("title")
	blog.Content = inputs.Get("content")
	blog.Created = time.Now()
	models.SaveBlog(blog)
	this.Ctx.Redirect(302, "/")
}