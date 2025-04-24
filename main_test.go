package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type StructWithMoreNested struct {
	StructWithStructField `keyname:"struct_with_struct_field"`
	Message               string `keyname:"message"`
}

type StructWithStructField struct {
	Message string       `keyname:"message"`
	Simple  SimpleStruct `keyname:"simple"`
}

type SimpleStruct struct {
	Name string `keyname:"name"`
	ID   int    `keyname:"id"`
}

func TestStructToMap(t *testing.T) {
	req := require.New(t)

	simple := SimpleStruct{
		Name: "Superman",
		ID:   3388,
	}

	withStructFiled := StructWithStructField{
		Simple:  simple,
		Message: "message",
	}

	withMoreNested := StructWithMoreNested{
		StructWithStructField: withStructFiled,
		Message:               "message",
	}

	t.Run("simple test", func(t *testing.T) {
		res := StructToMap(simple)
		req.NotNil(res)
		req.Equal(simple.Name, res["name"])
		req.Equal(simple.ID, res["id"])
	})

	t.Run("with struct field", func(t *testing.T) {
		res := StructToMap(withStructFiled)
		req.NotNil(res)
		req.Equal(withStructFiled.Message, res["message"])
		req.Equal(simple, res["simple"])
	})

	t.Run("more nesting", func(t *testing.T) {
		res := StructToMap(withMoreNested)
		req.NotNil(res)
		req.Equal(withMoreNested.Message, res["message"])
		req.Equal(withStructFiled, res["struct_with_struct_field"])
	})

	t.Run("not struct", func(t *testing.T) {
		res := StructToMap("is not struct")
		req.Nil(res)
	})
}

func TestMapToStruct(t *testing.T) {
	req := require.New(t)

	simple := map[string]interface{}{
		"name": "Superman",
		"id":   3388,
	}

	withStructField := map[string]interface{}{
		"simple":  SimpleStruct{Name: "Batman", ID: 123},
		"message": "hello",
	}

	withAnotherStrunct := map[string]interface{}{
		"struct_with_struct_field": StructWithStructField{
			Simple:  SimpleStruct{Name: "Batman", ID: 123},
			Message: "i am",
		},
		"message": "the night",
	}

	t.Run("simple test", func(t *testing.T) {
		var res SimpleStruct
		err := MapToStruct(simple, &res)
		req.NoError(err)
		req.Equal(simple["name"], res.Name)
		req.Equal(simple["id"], res.ID)
	})

	t.Run("with struct field", func(t *testing.T) {
		var res StructWithStructField
		err := MapToStruct(withStructField, &res)
		req.NoError(err)
		req.Equal(withStructField["message"], res.Message)
		req.Equal("Batman", res.Simple.Name)
		req.Equal(123, res.Simple.ID)
	})

	t.Run("more nesting", func(t *testing.T) {
		var res StructWithMoreNested
		err := MapToStruct(withAnotherStrunct, &res)
		req.NoError(err)
		req.Equal(withAnotherStrunct["message"], res.Message)
		req.Equal("i am", res.StructWithStructField.Message)
		req.Equal("Batman", res.Simple.Name)
		req.Equal(123, res.Simple.ID)
	})
}

func TestStructToMapAndBack(t *testing.T) {
	req := require.New(t)

	simple := SimpleStruct{
		Name: "Superman",
		ID:   3388,
	}

	withStructField := StructWithStructField{
		Simple:  simple,
		Message: "message",
	}

	withMoreNested := StructWithMoreNested{
		StructWithStructField: withStructField,
		Message:               "message",
	}

	t.Run("simple test", func(t *testing.T) {
		mp := StructToMap(simple)
		req.NotNil(mp)

		var resStruct SimpleStruct
		err := MapToStruct(mp, &resStruct)
		req.NoError(err)
		req.Equal(resStruct, simple)

		resMap := StructToMap(resStruct)
		req.NotNil(resMap)
		req.Equal(resMap, mp)
	})

	t.Run("with pointer", func(t *testing.T) {
		mp := StructToMap(withStructField)
		req.NotNil(mp)

		var resStruct StructWithStructField
		err := MapToStruct(mp, &resStruct)
		req.NoError(err)
		req.Equal(resStruct, withStructField)

		resMap := StructToMap(resStruct)
		req.NotNil(resMap)
		req.Equal(resMap, mp)
	})

	t.Run("more nesting", func(t *testing.T) {
		mp := StructToMap(withMoreNested)
		req.NotNil(mp)

		var resStruct StructWithMoreNested
		err := MapToStruct(mp, &resStruct)
		req.NoError(err)
		req.Equal(resStruct, withMoreNested)

		resMap := StructToMap(resStruct)
		req.NotNil(resMap)
		req.Equal(resMap, mp)
	})
}
