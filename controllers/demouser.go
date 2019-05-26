package controllers

import (
	"encoding/json"
	"log"

	"github.com/xlgwr/demoApigo/services"

	"github.com/astaxie/beego"
	"github.com/xlgwr/demoApigo/models"
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

	var du models.DemoUser
	json.Unmarshal(o.Ctx.Input.RequestBody, &du)
	id, err := demoUserDB.Add(du)
	if err != nil {
		log.Fatal(err)
	}
	o.Data["json"] = map[string]int64{"ID": id}
	o.ServeJSON()
}

// Get method
// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:id [get]
func (o *DemoUserController) Get() {
	id := o.Ctx.Input.Param(":id")
	o.Data["json"] = map[string]string{"ID": id}
	o.ServeJSON()
}
