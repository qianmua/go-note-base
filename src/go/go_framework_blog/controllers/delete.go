package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

/**
 * @author
 * @date 2020/11/11 15:35
 * create by Goland
 * @version 1.0
 */

type DeleteController struct {
	beego.Controller
}

func (this *DeleteController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params()[":id"])
	blog := models.GetBlog(id)
	this.Data["Post"] = blog
	models.DelBlog(blog)
	this.Ctx.Redirect(302, "/")
}