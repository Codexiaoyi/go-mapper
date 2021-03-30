package mapper

import (
	"errors"
	"reflect"
)

//根据字段名称将src中的值自动映射到dest中
func StructMapByFieldName(src interface{}, dest interface{}) error {
	dic := make(map[string]reflect.Value)
	srcPtr := reflect.ValueOf(src).Elem()
	destPtr := reflect.ValueOf(dest).Elem()

	if srcPtr.Kind() != reflect.Struct || destPtr.Kind() != reflect.Struct {
		return errors.New("Only type of Ptr")
	}

	//存储src字段信息
	for i := 0; i < srcPtr.NumField(); i++ {
		field := srcPtr.Type().Field(i)                  //获取到字段
		dic[field.Name] = srcPtr.FieldByName(field.Name) //将字段保存
	}

	for i := 0; i < destPtr.NumField(); i++ {
		currentField := destPtr.Type().Field(i)
		name := currentField.Name
		//如果与src中字段名匹配并且类型相同则赋值
		if dic[name].IsValid() && dic[name].Kind() == currentField.Type.Kind() && dic[name].CanSet() {
			destPtr.FieldByName(name).Set(dic[name])
		}
	}

	return nil
}
