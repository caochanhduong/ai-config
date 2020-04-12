package controllers

import (
	"ai-config-backend/models"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type MapAiTypeToAiNameController struct {
	beego.Controller
}

func (c *MapAiTypeToAiNameController) GetAllMapAiTypeToAiName() {
	list, err := models.GetAllMapAiTypeToAiName()
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
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
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
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
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")


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
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
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

func (c*MapAiTypeToAiNameController) FindMapByAiTypeAndAiName(){
	

	// if v, err := c.GetInt64("ai_type"); err == nil {
	// 	ai_type = v
	// }

	// ai_name = c.GetString("ai_name","")
	aitypeStr := c.Ctx.Input.Param(":ai_type")
	ainameStr := c.Ctx.Input.Param(":ai_name")
	aitypeInt, err := strconv.ParseInt(aitypeStr, 10, 64)
	fmt.Println("------------------ai_type")
	fmt.Println(aitypeInt)
	fmt.Println("------------------ai_name")
	fmt.Println(ainameStr)

	// TODO:Validate ai_type as int64, ai_name as string

	res, err := models.FindMapByAiTypeAndAiName(aitypeInt,ainameStr)
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil &&  err != orm.ErrNoRows{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = "Can not find AI Type"
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = res

	
	c.ServeJSON()
}

func (c*MapAiTypeToAiNameController) DeleteMapByAiType() {
	aiTypeStr := c.Ctx.Input.Param(":ai_type")
	aiTypeInt, _ := strconv.ParseInt(aiTypeStr, 10, 64)
	fmt.Println("------------------aiTypeInt")
	fmt.Println(aiTypeInt)


	// TODO:Validate ai_type as int64, ai_name as string
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	err := models.DeleteMapByAiType(aiTypeInt)

	if err != nil{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = "Error delete Map"
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = "Delete success"

	c.ServeJSON()
}
//:id=1,2,3,4 => split => [1,2,3,4]
func (c*MapAiTypeToAiNameController) DeleteMapByAiTypes() {
	aiTypeStr := c.Ctx.Input.Param(":ai_types")
	listaiTypeStr := strings.Split(aiTypeStr, ",") 
	var listaiTypeInt []int64
	for _,v:= range listaiTypeStr{
		aiTypeInt, _ := strconv.ParseInt(v, 10, 64)
		listaiTypeInt = append(listaiTypeInt, aiTypeInt)
	}
	
	fmt.Println("------------------listaiTypeInt")
	fmt.Println(listaiTypeInt)

	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// TODO:Validate ai_type as int64, ai_name as string

	err := models.DeleteUserByAiTypes(listaiTypeInt)

	if err != nil{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = "Error delete Maps"
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = "Delete success"

	c.ServeJSON()
}