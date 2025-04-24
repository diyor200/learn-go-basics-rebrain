package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Class Class  `json:"class"`
}

type Class struct {
	Name string `json:"name"`
}

func main() {
	user := User{
		Name: "Alex",
		Age:  15,
		Class: Class{
			Name: "9e",
		},
	}

	m := map[string]interface{}{
		"name": "Alex",
		"age":  15,
		"class": map[string]interface{}{
			"name": "9e",
		},
	}

	var r User
	_ = MapToStruct(m, &r)
	fmt.Println(r)
	fmt.Println(StructToMap(user))
}

func MapToStruct(mp map[string]interface{}, item interface{}) error {
	v := reflect.ValueOf(item).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("json")
		mapVal, ok := mp[tag]
		if !ok {
			continue
		}

		val := reflect.ValueOf(mapVal)

		if val.Kind() == reflect.Struct {
			nest := map[string]interface{}{}
			MapToStruct(nest)
		}

		if val.Type().AssignableTo(v.Field(i).Type()) {
			v.Field(i).Set(val)
		}
	}

	return nil
}

func StructToMap(item interface{}) map[string]interface{} {
	v := reflect.TypeOf(item)
	val := reflect.ValueOf(item)

	if val.Kind() != reflect.Struct {
		return nil
	}

	var result = map[string]interface{}{}
	for i := 0; i < val.NumField(); i++ {
		fmt.Println(v.Field(i).Tag.Get("json"))
		if v.Field(i).Type.Kind() == reflect.Struct {
			result[v.Field(i).Tag.Get("json")] = StructToMap(val.Field(i).Interface())
		} else {
			result[v.Field(i).Tag.Get("json")] = val.Field(i).Interface()
		}
	}

	return result
}
