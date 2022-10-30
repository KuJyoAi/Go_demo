package main

import (
	"fmt"
	"reflect"
)

type sorder struct {
	name   string
	gender string
	age    int
}

func main() {
	s := sorder{
		name:   "苍井春人",
		gender: "男",
		age:    18,
	}
	fmt.Println(EnumStruct(s))

}

func EnumStruct(src interface{}) string {
	_type := reflect.TypeOf(src)
	_value := reflect.ValueOf(src)
	// 不是结构体
	if _type.Kind() != reflect.Struct {
		return ""
	}
	var str string
	for i := 0; i < _type.NumField(); i++ {
		typeField := _type.Field(i)
		valueField := _value.Field(i)

		switch typeField.Type.Kind() {
		case reflect.String:
			str += fmt.Sprintf("%s:%s ", typeField.Name, valueField.String())
		case reflect.Int:
			str += fmt.Sprintf("%s:%d ", typeField.Name, valueField.Int())
		case reflect.Bool:
			if valueField.Bool() {
				str += fmt.Sprintf("%s:true ", typeField.Name)
			} else {
				str += fmt.Sprintf("%s:false ", typeField.Name)
			}
		}
	}
	return str
}
