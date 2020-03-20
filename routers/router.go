package routers

import (
	"ai-config-project/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/map", &controllers.MapAiTypeToAiNameController{},"get:GetAllMapAiTypeToAiName")
	beego.Router("/users", &controllers.UserController{},"get:GetAllUser")
	beego.Router("/admin-users", &controllers.AdminUserController{},"get:GetAllAdminUser")
	beego.Router("/admin-users/exist", &controllers.AdminUserController{},"post:ExistAdminUserByAccountAndPassword")
	beego.Router("/users/exist", &controllers.UserController{},"post:ExistUserByID")
	beego.Router("/map/exist", &controllers.MapAiTypeToAiNameController{},"post:ExistMapByAIType")
	beego.Router("/map", &controllers.MapAiTypeToAiNameController{},"post:AddMap")
	beego.Router("/users", &controllers.UserController{},"post:AddUser")
	beego.Router("/map", &controllers.MapAiTypeToAiNameController{},"put:UpdateMapByAiType")
	beego.Router("/users", &controllers.UserController{},"put:UpdateUserById")
}
