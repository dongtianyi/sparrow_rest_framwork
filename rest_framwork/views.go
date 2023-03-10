package rest_framwork

import (
	"errors"
	"fmt"
	"reflect"
)

type ListAPIView struct {
	model interface{} // 设置model
}

func (v *ListAPIView) SetModel(mi interface{}) {
	v.model = mi
}

func (v *ListAPIView) List() (interface{}, error) {
	if v.model == nil {
		return nil, errors.New("model 未设置")
	}

	xxx := reflect.TypeOf(v.model)
	fmt.Println(xxx)
	// yyy := reflect.ValueOf(v.model)
	// fmt.Println(yyy)
	// 查询数据库
	// db := database.GetDb()
	// // 反射,获取到真正的类型
	res, ok := v.model.(xxx)
	x := make(xxx, 0)
	// err := db.Find(x).Error
	// if err != nil {
	// 	return nil, err
	// }
	// return &v.model, nil
	return nil, nil
}
