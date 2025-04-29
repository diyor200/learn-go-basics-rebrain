package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Class   Class   `json:"class"`
	Account Account `json:"account"`
}

type Account struct {
	Balance int `json:"balance"`
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
		Account: Account{
			Balance: 15,
		},
	}

	m := map[string]interface{}{
		"name": "Alex",
		"age":  15,
		"class": map[string]interface{}{
			"name": "9e",
		},
		"account": map[string]interface{}{
			"balance": 15,
		},
	}

	var r User
	//_ = MapToStruct(m, &r)
	_ = MTS(m, &r)
	fmt.Println(r)
	//fmt.Println(StructToMap(user))
	fmt.Println("mp = ", STM(user))

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

		if val.Kind() == reflect.Map && v.Field(i).Kind() == reflect.Struct {
			ptr := reflect.New(v.Field(i).Type())
			err := MapToStruct(mapVal.(map[string]interface{}), ptr.Interface())
			if err != nil {
				return err
			}
			v.Field(i).Set(ptr.Elem())
		}

		if val.Type().AssignableTo(v.Field(i).Type()) {
			v.Field(i).Set(val)
		}
	}

	return nil
}

func MTS(mp map[string]interface{}, res interface{}) error {
	v := reflect.ValueOf(res)
	if v.Kind() == reflect.Ptr {
		v = reflect.ValueOf(res).Elem()
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		if key == "" {
			key = t.Field(i).Name
		}

		mpVal, ok := mp[key]
		if !ok {
			continue
		}

		val := reflect.ValueOf(mpVal)
		if val.Kind() == reflect.Map && v.Field(i).Kind() == reflect.Struct {
			childType := reflect.New(v.Field(i).Type())

			err := MTS(mpVal.(map[string]interface{}), childType.Interface())
			if err != nil {
				return err
			}

			v.Field(i).Set(childType.Elem())
		} else if v.Field(i).Kind() == val.Kind() {
			v.Field(i).Set(val)
		}
	}

	return nil
}

func STM(data interface{}) map[string]interface{} {
	v := reflect.ValueOf(data)
	t := v.Type()

	var mp = map[string]interface{}{}
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)
		key := f.Tag.Get("json")
		if key == "" {
			key = t.Field(i).Name
		}

		if v.Field(i).Kind() == reflect.Struct {
			mp[key] = STM(v.Field(i).Interface())
		} else {
			mp[key] = v.Field(i).Interface()
		}
	}

	return mp
}

func StructToMap(item interface{}) map[string]interface{} {
	v := reflect.TypeOf(item)
	val := reflect.ValueOf(item)

	if val.Kind() != reflect.Struct {
		return nil
	}

	var result = map[string]interface{}{}
	for i := 0; i < val.NumField(); i++ {
		if v.Field(i).Type.Kind() == reflect.Struct {
			result[v.Field(i).Tag.Get("json")] = StructToMap(val.Field(i).Interface())
		} else {
			result[v.Field(i).Tag.Get("json")] = val.Field(i).Interface()
		}
	}

	return result
}
