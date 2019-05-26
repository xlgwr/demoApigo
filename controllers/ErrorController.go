package controllers

import (
	"github.com/astaxie/beego"
	m "github.com/xlgwr/demoApigo/models"
)

type ErrorController struct {
	beego.Controller
}

func (o *ErrorController) Error404() {
	result := m.ApiMessage{404, false, o.Ctx.Request.URL.Path + " page not found", nil}
	o.Data["json"] = result
	o.ServeJSON()
}

func (o *ErrorController) Error501() {
	result := m.ApiMessage{501, false, "server error", nil}
	o.Data["json"] = result
	o.ServeJSON()
}

func (o *ErrorController) ErrorDb() {
	result := m.ApiMessage{-1, false, "database is now down", nil}
	o.Data["json"] = result
	o.ServeJSON()
}
