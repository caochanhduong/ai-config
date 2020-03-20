package controllers

import (
	"ai-config-project/models"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type MapAiTypeToAiNameController struct {
	beego.Controller
}

func (c *MapAiTypeToAiNameController) GetAllMapAiTypeToAiName() {
	list, err := models.GetAllMapAiTypeToAiName()
	if err != nil {
		c.Ctx.Abort(http.StatusBadRequest, err.Error())
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = list
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.ServeJSON()
	}
}

func (c *MapAiTypeToAiNameController) AddMap() {
	var ob models.MapAiTypeToAiName
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)
	ai_type := ob.AiType
	ai_name := ob.AiName
	fmt.Println("------------------ai_type")
	fmt.Println(ai_type)
	fmt.Println("------------------ai_name")
	fmt.Println(ai_name)

	// TODO:Validate ai_type as int64, ai_name as string

	map_ai := models.MapAiTypeToAiName{}
	map_ai.AiType = ai_type
	map_ai.AiName = ai_name

	if _, err := models.AddMap(&map_ai); err == nil{
		c.Ctx.Output.SetStatus(http.StatusCreated)
		c.Data["json"] = map_ai
	} else {
		c.Ctx.Abort(http.StatusBadRequest, err.Error())
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *MapAiTypeToAiNameController) ExistMapByAIType() {
	var ob models.MapAiTypeToAiName
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)
	
	ai_type := ob.AiType
	fmt.Println("------------------ai_type")
	fmt.Println(ai_type)

	// TODO:Validate ai_type as int64



	if exist, err := models.ExistMapByAIType(ai_type); err != nil && err != orm.ErrNoRows {
		c.Ctx.Abort(http.StatusBadRequest, err.Error())
		c.Data["json"] = err.Error()
	} else {
		if exist {
			c.Data["json"] = "existed"
			c.Ctx.Output.SetStatus(http.StatusOK)
		} else {
			c.Data["json"] = "not existed"
			c.Ctx.Output.SetStatus(http.StatusOK)
		}
	}
	c.ServeJSON()
}

func (c *MapAiTypeToAiNameController) UpdateMapByAiType() {
	var ob models.MapAiTypeToAiName
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)
	ai_type := ob.AiType
	ai_name := ob.AiName
	fmt.Println("------------------ai_type")
	fmt.Println(ai_type)
	fmt.Println("------------------ai_name")
	fmt.Println(ai_name)

	// TODO:Validate ai_type as int64, ai_name as string

	map_ai := models.MapAiTypeToAiName{}
	map_ai.AiType = ai_type
	map_ai.AiName = ai_name

	exist, err := models.ExistMapByAIType(ai_type)
	if err != nil &&  err != orm.ErrNoRows{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = "Can not update AI Type"
		c.ServeJSON()
		return
	}

	if !exist {
		c.Data["json"] = "AI Type not existed"
		c.ServeJSON()
		return
	}

	if _, err := models.UpdateMapByAiType(&map_ai); err == nil {
		// fmt.Println("duongcc")
		c.Ctx.Output.SetStatus(http.StatusCreated)
		c.Data["json"] = map_ai
	}
	
	c.ServeJSON()
}
