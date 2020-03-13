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
	AiType string `orm:"column(ai_type);type(text)" json:"ai_type"`
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
