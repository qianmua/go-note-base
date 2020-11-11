package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

/**
 * @author
 * @date 2020/11/11 15:36
 * create by Goland
 * @version 1.0
 */

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplName = "view.tpl"
}