package controllers

import (
	"ai-config-project/models"
	"github.com/astaxie/beego"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
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

func (c *UserController) ExistUserByID(){
	var ob models.User
	json.Unmarshal(c.Ctx.Input.RequestBody,&ob)
	fmt.Println(ob)
	id := ob.Id

	
	fmt.Println("------------------id")
	fmt.Println(id)
	//TODO:Validate id as int64



	if exist,err := models.ExistUserById(id); err != nil && err != orm.ErrNoRows{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = err.Error()
	}else{
		if exist {
			c.Data["json"] = "existed"
			c.Ctx.Output.SetStatus(http.StatusOK)	
		}else{
			c.Data["json"] = "not existed"
			c.Ctx.Output.SetStatus(http.StatusOK)
		}
	}
	c.ServeJSON()
}

func (c *UserController) AddUser() {
	var ob models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)
	id := ob.Id
	account := ob.Account
	ai_type := ob.AiType
	start_time := ob.StartTime
	duration := ob.Duration
	end_time := start_time + duration

	fmt.Println("------------------id")
	fmt.Println(id)
	fmt.Println("------------------account")
	fmt.Println(account)
	fmt.Println("------------------ai_type")
	fmt.Println(ai_type)
	fmt.Println("------------------start_time")
	fmt.Println(start_time)
	fmt.Println("------------------duration")
	fmt.Println(duration)
	fmt.Println("------------------end_time")
	fmt.Println(end_time)

	// TODO:Validate input


	user := models.User{}
	user.Id = id
	user.Account = account
	user.AiType = ai_type
	user.StartTime = start_time
	user.EndTime = end_time
	user.Duration = duration

	if _, err := models.AddUser(&user); err == nil{
		c.Ctx.Output.SetStatus(http.StatusCreated)
		c.Data["json"] = user
	} else {
		c.Ctx.Abort(http.StatusBadRequest, err.Error())
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *UserController) UpdateUserById() {
	var ob models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)

	id := ob.Id
	account := ob.Account
	ai_type := ob.AiType
	start_time := ob.StartTime
	duration := ob.Duration
	end_time := start_time + duration

	fmt.Println("------------------id")
	fmt.Println(id)
	fmt.Println("------------------account")
	fmt.Println(account)
	fmt.Println("------------------ai_type")
	fmt.Println(ai_type)
	fmt.Println("------------------start_time")
	fmt.Println(start_time)
	fmt.Println("------------------duration")
	fmt.Println(duration)
	fmt.Println("------------------end_time")
	fmt.Println(end_time)

	// TODO:Validate input


	user := models.User{}
	user.Id = id
	user.Account = account
	user.AiType = ai_type
	user.StartTime = start_time
	user.EndTime = end_time
	user.Duration = duration

	exist, err := models.ExistUserById(id)
	if err != nil &&  err != orm.ErrNoRows {
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = "Can not update User"
		c.ServeJSON()
		return
	}

	if !exist {
		c.Data["json"] = "User not existed"
		c.ServeJSON()
		return
	}

	if _, err := models.UpdateUserById(&user); err == nil {
		// fmt.Println("duongcc")
		c.Ctx.Output.SetStatus(http.StatusCreated)
		c.Data["json"] = user
	}
	
	c.ServeJSON()
}

func (c*UserController) FindUser(){
	

	// if v, err := c.GetInt64("ai_type"); err == nil {
	// 	ai_type = v
	// }

	// ai_name = c.GetString("ai_name","")
	idStr := c.Ctx.Input.Param(":id")
	accountStr := c.Ctx.Input.Param(":account")
	aiTypeStr := c.Ctx.Input.Param(":ai_type")
	startTimeStr := c.Ctx.Input.Param(":start_time")
	durationStr := c.Ctx.Input.Param(":duration")
	endTimeStr := c.Ctx.Input.Param(":end_time")


	idInt, err := strconv.ParseInt(idStr, 10, 64)
	aiTypeInt, err := strconv.ParseInt(aiTypeStr, 10, 64)
	startTimeInt, err := strconv.ParseInt(startTimeStr, 10, 64)
	durationInt, err := strconv.ParseInt(durationStr, 10, 64)
	endTimeInt, err := strconv.ParseInt(endTimeStr, 10, 64)

	fmt.Println("------------------idInt")
	fmt.Println(idInt)
	fmt.Println("------------------aiTypeInt")
	fmt.Println(aiTypeInt)

	// TODO:Validate ai_type as int64, ai_name as string

	res, err := models.FindUser(idInt,accountStr,aiTypeInt,startTimeInt,durationInt,endTimeInt)

	if err != nil &&  err != orm.ErrNoRows{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = "Can not find user"
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = res

	
	c.ServeJSON()
}

func (c*UserController) DeleteUserById() {
	idStr := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.ParseInt(idStr, 10, 64)
	fmt.Println("------------------idInt")
	fmt.Println(idInt)


	// TODO:Validate ai_type as int64, ai_name as string

	err := models.DeleteUserById(idInt)

	if err != nil{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = "Error delete User"
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = "Delete success"

	c.ServeJSON()
}
//:id=1,2,3,4 => split => [1,2,3,4]
func (c*UserController) DeleteUserByIds() {
	idStr := c.Ctx.Input.Param(":ids")
	listIdStr := strings.Split(idStr, ",") 
	var listIdInt []int64
	for _,v:= range listIdStr{
		idInt, _ := strconv.ParseInt(v, 10, 64)
		listIdInt = append(listIdInt, idInt)
	}
	
	fmt.Println("------------------listIdInt")
	fmt.Println(listIdInt)


	// TODO:Validate ai_type as int64, ai_name as string

	err := models.DeleteUserByIds(listIdInt)

	if err != nil{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = "Error delete Users"
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = "Delete success"

	c.ServeJSON()
}