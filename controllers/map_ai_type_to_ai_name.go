package controllers

import (
	"ai-config-project/models"
	"github.com/astaxie/beego"
	"net/http"
)

type MapAiTypeToAiNameController struct {
	beego.Controller
}

func (c *MapAiTypeToAiNameController) GetAllMapAiTypeToAiName() {
	list,err := models.GetAllMapAiTypeToAiName()
	if err != nil{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = err.Error()
	}else{
		c.Data["json"] = list
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.ServeJSON()
	}
}
