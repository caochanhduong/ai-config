package routers

import (
	"ai-config-project/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ai-types", &controllers.MapAiTypeToAiNameController{},"get:GetAllMapAiTypeToAiName")
	beego.Router("/users", &controllers.UserController{},"get:GetAllUser")
	beego.Router("/admin-users", &controllers.AdminUserController{},"get:GetAllAdminUser")
	beego.Router("/admin-users/exist", &controllers.AdminUserController{},"post:ExistAdminUserByAccountAndPassword")
}
