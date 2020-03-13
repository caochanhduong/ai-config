package controllers

import (
	"ai-config-project/models"
	"github.com/astaxie/beego"
	"net/http"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) GetAllUser() {
	list,err := models.GetAllUser()
	if err != nil{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = err.Error()
	}else{
		c.Data["json"] = list
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.ServeJSON()
	}
}
