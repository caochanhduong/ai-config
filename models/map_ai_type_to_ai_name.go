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

func ExistMapByAIType(ai_type int64) (bool,error) {
	o := orm.NewOrm()
	qs := o.QueryTable(MySqlMapAiTypeToAiName).Filter("ai_type",ai_type)
	res := MapAiTypeToAiName{}
	err := qs.One(&res)
	if err != nil {
		return false, err
	}
	if (res != MapAiTypeToAiName{}) {
		return true, nil
	}
	return false, nil
}

func AddMap(map_ai *MapAiTypeToAiName) (int64, error){
	o := orm.NewOrm()
	return o.Insert(map_ai)
}

func UpdateMapByAiType(map_ai *MapAiTypeToAiName) (int64, error) {
	o := orm.NewOrm()
	v := MapAiTypeToAiName{AiType: map_ai.AiType}
	err := o.Read(&v)
	if err == nil {
		return o.Update(map_ai)
	}
	return 0, err
}