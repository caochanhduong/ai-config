package models

import (
	

	"github.com/astaxie/beego/orm"
)

const (
	MySqlMapAiTypeToAiName = "map_ai_type_to_ai_name"
)

type MapAiTypeToAiName struct {
	AiType     int64  `orm:"column(ai_type);type(bigint);pk" json:"ai_type"`
	AiName string `orm:"column(ai_name);type(text)" json:"ai_name"`
}


// }
// func (t *AIType) TableUnique() [][]string {
// 	return [][]string{
// 		[]string{"ID", "AIName"},
// 	}
// }

func init() {
	orm.RegisterModel(new(MapAiTypeToAiName))
}

func GetAllMapAiTypeToAiName() ([] MapAiTypeToAiName,error){
	var res []MapAiTypeToAiName
	o := orm.NewOrm()
	// model := new(MapAiTypeToAiName)
	// model.AiType = 12
	// model.AiName = "duongcc"
	// o.Using("ai_config")
	_, err := o.QueryTable(MySqlMapAiTypeToAiName).All(&res)
	// fmt.Println(num)
	// fmt.Println(o.Insert(model))
	// fmt.Println("num")
	if err!=nil {
		return nil,err
	}
	return res, nil
}
