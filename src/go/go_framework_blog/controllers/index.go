package controllers

import "github.com/astaxie/beego"

/**
 * @author
 * @date 2020/11/11 15:36
 * create by Goland
 * @version 1.0
 */

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["blogs"] = models.GetAll()
	this.Layout = "layout.tpl"
	this.TplName = "index.tpl"
}