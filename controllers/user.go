package controllers

import (
	"ai-config-project/models"
	"github.com/astaxie/beego"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	
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