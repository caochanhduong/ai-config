package controllers

import (
	"ai-config-backend/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type AdminUserController struct {
	beego.Controller
}

func (c *AdminUserController) GetAllAdminUser() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
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
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	var ob models.AdminUser
	json.Unmarshal(c.Ctx.Input.RequestBody,&ob)
	fmt.Println(ob)
	account := ob.Account
	password := ob.Password
	fmt.Println("------------------account")
	fmt.Println(account)
	fmt.Println("------------------password")
	fmt.Println(password)
	

    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
	if exist,err := models.ExistAdminUserbyAccount(account); err != nil && err!= orm.ErrNoRows{
		c.Ctx.Abort(http.StatusBadRequest,err.Error())
		c.Data["json"] = err.Error()
	}else{
		//if exist check password
		if exist {
			admin_user_db,err := models.FindAdminUserByAccount(account)
			if err != nil {
				c.Ctx.Abort(http.StatusBadRequest,err.Error())
				c.Data["json"] = err.Error()
			}
			//match password hash
			if bcrypt.CompareHashAndPassword( []byte(admin_user_db.Password), []byte(password)) == nil{
				c.Data["json"] = "match account"
			}else{
				c.Data["json"] = "no match account"
			}
			
			c.Ctx.Output.SetStatus(http.StatusOK)	
		}else{
			c.Data["json"] = "not existed"
			c.Ctx.Output.SetStatus(http.StatusOK)
		}
	}
	c.ServeJSON()
}

func (c *AdminUserController) ExistAdminUserByAccount(){
	var ob models.AdminUser
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	json.Unmarshal(c.Ctx.Input.RequestBody,&ob)
	fmt.Println(ob)
	account := ob.Account
	fmt.Println("------------------account")
	fmt.Println(account)


	if exist,err := models.ExistAdminUserbyAccount(account); err != nil && err != orm.ErrNoRows{
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

func (c *AdminUserController) AddAdminUser() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	var ob models.AdminUser
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	fmt.Println(ob)
	account := ob.Account
	password := ob.Password

	fmt.Println("------------------account")
	fmt.Println(account)
	fmt.Println("------------------password")
	fmt.Println(password)

	// TODO:Validate input
	admin_user := models.AdminUser{}
	admin_user.Account = account
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	admin_user.Password = string(hashPassword)
	if _, err := models.AddAdminUser(&admin_user); err == nil{
		c.Ctx.Output.SetStatus(http.StatusCreated)
		c.Data["json"] = admin_user
	} else {
		c.Ctx.Abort(http.StatusBadRequest, err.Error())
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}