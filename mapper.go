package mapper

import (
	"errors"
	"reflect"
)

//根据字段名称将src中的值自动映射到dest中
func StructMapByFieldName(src interface{}, dest interface{}) error {
	if reflect.TypeOf(src).Kind() != reflect.Ptr || reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("src and dst must be addressable.")
	}

	dic := make(map[string]reflect.Value)
	srcPtr := reflect.ValueOf(src).Elem()
	destPtr := reflect.ValueOf(dest).Elem()

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

func StructMapByTag(src interface{}, dest interface{}) error {
	//not addressable
	if reflect.TypeOf(src).Kind() != reflect.Ptr || reflect.TypeOf(dest).Kind() != reflect.Ptr {
		return errors.New("src and dst must be addressable.")
	}

	tSrc, vSrc, tDst, vDst := reflect.TypeOf(src).Elem(), reflect.ValueOf(src).Elem(),
		reflect.TypeOf(dest).Elem(), reflect.ValueOf(dest).Elem()

	//建立一个map存储字段与我们的tag的映射关系
	tagMap := make(map[string]reflect.Value)

	//首先遍历我们的dst，将所有的tag与对应的字段映射起来
	for i := 0; i < vDst.NumField(); i++ {
		if val, ok := tDst.Field(i).Tag.Lookup("mapper"); ok {
			tagMap[val] = vDst.Field(i)
		}
	}
	//然后遍历我们的request，遍历所有的field，每次获取到tag，然后填充对应的内容到我们的resp的field中
	for i := 0; i < vSrc.NumField(); i++ {
		if val, ok := tSrc.Field(i).Tag.Lookup("mapper"); ok {
			//通过val与tagMap找到对应于vDst的字段
			//有可能map中的值不存在
			if value, ok := tagMap[val]; ok && value.IsValid() && value.CanSet() && vSrc.Field(i).Kind() == value.Kind() {
				value.Set(vSrc.Field(i))
			}
		}
	}
	return nil
}
