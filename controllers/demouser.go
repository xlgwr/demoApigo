package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/xlgwr/demoApigo/services"

	"github.com/astaxie/beego"
	m "github.com/xlgwr/demoApigo/models"
)

var demoUserDB services.DemoUserServices

type DemoUserController struct {
	beego.Controller
}

// Post ..
// @Title Create
// @Description create DemoUser
// @Param body body models.DemoUser	true "The DemoUser content"
// @Success 200 {int64} models.DemoUser.ID
// @Failure 403 body is empty
// @router / [post]
func (o *DemoUserController) Post() {
	result := m.ApiMessage{0, true, "", nil}
	var du m.DemoUser
	json.Unmarshal(o.Ctx.Input.RequestBody, &du)
	id, err := demoUserDB.Add(du)
	if err != nil {
		result.Success = false
		result.Message = err.Error()
	} else {
		result.Data = map[string]int64{"ID": id}
	}
	o.Data["json"] = result
	o.ServeJSON()
}

// Get method
// @Title Get
// @Description find DemoUser by id
// @Param id path string true "the id you want to get"
// @Success 200 {DemoUser} models.DemoUser
// @Failure 403 :id is empty
// @router /:id([0-9]+) [get]
func (o *DemoUserController) Get() {
	result := m.ApiMessage{0, true, "", nil}
	id := o.Ctx.Input.Param(":id")
	intid, err := strconv.Atoi(id)
	if err != nil {
		result.Success = false
		result.Message = err.Error()
		o.Data["json"] = result
		o.ServeJSON()
		return
	}
	u, err := demoUserDB.Get(intid)
	if err != nil {
		result.Success = false
		result.Message = err.Error()
	} else {
		result.Data = u
	}
	o.Data["json"] = result
	o.ServeJSON()
}
