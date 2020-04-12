package models

import (
	

	"github.com/astaxie/beego/orm"
)

const (
	MySqlTableUser = "user"
)

type User struct {
	Id     int64  `orm:"column(id);type(bigint);pk" json:"id"`
	Account string `orm:"column(account);type(text)" json:"account"`
	AiType int64 `orm:"column(ai_type);type(bigint)" json:"ai_type"`
	StartTime int64 `orm:"column(start_time);type(bigint)" json:"start_time"`
	Duration int64 `orm:"column(duration);type(bigint)" json:"duration"`
	EndTime int64 `orm:"column(end_time);type(bigint)" json:"end_time"`
}


// }
// func (t *AIType) TableUnique() [][]string {
// 	return [][]string{
// 		[]string{"ID", "AIName"},
// 	}
// }

func init() {
	orm.RegisterModel(new(User))
}

func GetAllUser() ([] User,error){
	var res []User
	o := orm.NewOrm()
	// model := new(MapAiTypeToAiName)
	// model.AiType = 12
	// model.AiName = "duongcc"
	// o.Using("ai_config")
	_, err := o.QueryTable(MySqlTableUser).All(&res)
	// fmt.Println(num)
	// fmt.Println(o.Insert(model))
	// fmt.Println("num")
	if err!=nil {
		return nil,err
	}
	return res, nil
}

func ExistUserById(id int64) (bool,error) {
	o := orm.NewOrm()
	qs := o.QueryTable(MySqlTableUser).Filter("id",id)
	res := User{}
	err := qs.One(&res)
	if err != nil {
		return false, err
	}
	if (res != User{}) {
		return true, nil
	}
	return false, nil
}

func AddUser(user *User) (int64, error){
	o := orm.NewOrm()
	return o.Insert(user)
}

func UpdateUserById(user *User) (int64, error) {
	o := orm.NewOrm()
	v := User{Id: user.Id}
	err := o.Read(&v)
	if err == nil {
		return o.Update(user)
	}
	return 0, err
}

func FindUser(id int64, account string, ai_type int64, start_time int64, duration int64, end_time int64) ([]User, error) {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	
	if id != 0 {
		cond = cond.And("id",id)
	}
	if ai_type != 0 {
		cond = cond.And("ai_type",ai_type)
	}
	if start_time != 0 {
		cond = cond.And("start_time",start_time)
	}
	if duration != 0 {
		cond = cond.And("duration",duration)
	}
	if end_time != 0 {
		cond = cond.And("end_time",end_time)
	}
	if account != " " {
		cond = cond.And("account",account)
	}
	var v []User
	qs := o.QueryTable(MySqlTableUser)
	qs = qs.SetCond(cond)
	_, err := qs.All(&v)
	if err != nil{
		return nil, err
	}
	return v, nil
}

func DeleteUserById(id int64) (error) {
	o := orm.NewOrm()
	_, err := o.QueryTable(MySqlTableUser).Filter("id",id).Delete()
	return err
}
func DeleteUserByIds(ids []int64) error {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	for _, id := range ids {
		if id != 0 {
			cond = cond.Or("id",id)
		}
	}

	_, err := o.QueryTable(MySqlTableUser).SetCond(cond).Delete()

	return err
}