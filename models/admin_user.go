package models

import (
	

	"github.com/astaxie/beego/orm"
)

const (
	MySqlTableAdminUser = "admin_user"
)

type AdminUser struct {
	Account     string  `orm:"column(account);type(text);pk" json:"account"`
	Password string `orm:"column(password);type(text)" json:"password"`
}


// }
// func (t *AIType) TableUnique() [][]string {
// 	return [][]string{
// 		[]string{"ID", "AIName"},
// 	}
// }

func init() {
	orm.RegisterModel(new(AdminUser))
}

func GetAllAdminUser() ([] AdminUser,error){
	var res []AdminUser
	o := orm.NewOrm()
	// model := new(MapAiTypeToAiName)
	// model.AiType = 12
	// model.AiName = "duongcc"
	// o.Using("ai_config")
	_, err := o.QueryTable(MySqlTableAdminUser).All(&res)
	// fmt.Println(num)
	// fmt.Println(o.Insert(model))
	// fmt.Println("num")
	if err!=nil {
		return nil,err
	}
	return res, nil
}

func ExistAdminUserByAccountAndPassword (account string, password string) (bool,error){
	o := orm.NewOrm()
	qs := o.QueryTable(MySqlTableAdminUser).Filter("Account",account).Filter("Password",password)
	res := AdminUser{}
	err := qs.One(&res)
	if err != nil {
		return false, err
	}

	if (res != AdminUser{}) {
		return true, nil
	}
	return false, nil
}

func AddAdminUser(admin_user *AdminUser) (int64, error){
	o := orm.NewOrm()
	return o.Insert(admin_user)
}

func ExistAdminUserbyAccount(account string) (bool,error){
	o := orm.NewOrm()
	qs := o.QueryTable(MySqlTableAdminUser).Filter("Account",account)
	res := AdminUser{}
	err := qs.One(&res)
	if err != nil {
		return false, err
	}
	if (res != AdminUser{}) {
		return true, nil
	}
	return false, nil
}

func FindAdminUserByAccount(account string) (*AdminUser,error){
	o := orm.NewOrm()
	qs := o.QueryTable(MySqlTableAdminUser).Filter("Account",account)
	res := AdminUser{}
	err := qs.One(&res)
	if err != nil {
		return nil, err
	}
	if (res != AdminUser{}) {
		return &res, nil
	}
	return nil, nil
}