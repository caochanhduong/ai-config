package controllers

import (
	"ai-config-project/models"
	"encoding/json"
	"fmt"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/astaxie/beego"
)

type AdminUserController struct {
	beego.Controller
}

func (c *AdminUserController) GetAllAdminUser() {
	if list,err := models.GetAllAdminUser(); err != nil{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = err.Error()
	}else{
		c.Data["json"] = list
		c.Ctx.Output.SetStatus(http.StatusOK)
		c.ServeJSON()
	}
}

func (c *AdminUserController) ExistAdminUserByAccountAndPassword(){
	var ob models.AdminUser
	json.Unmarshal(c.Ctx.Input.RequestBody,&ob)
	fmt.Println(ob)
	account := ob.Account
	password := ob.Password
	fmt.Println("------------------account")
	fmt.Println(account)
	fmt.Println("------------------password")
	fmt.Println(password)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = err.Error()
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    fmt.Println(string(hash))
	if exist,err := models.ExistAdminUserByAccountAndPassword(account,string(hash)); err != nil{
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
