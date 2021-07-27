package main

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

var refCache sync.Map

type OrmReflect struct {
	val       reflect.Value
	typ       reflect.Type
	insertSQL string
}

func (o *OrmReflect) generateInsertSQL() {
	// insert into %s(f...) values (?...)
	numField := o.typ.NumField()
	for i := 0; i < numField; i++ {
		field := o.typ.Field(i)
		// t.Field(i).Name
		fmt.Println(field.Name)
	}
}

func GetReflect(i interface{}) *OrmReflect {
	if val, ok := refCache.Load(i); ok {
		return val.(*OrmReflect)
	}
	result := new(OrmReflect)
	result.val = reflect.ValueOf(i)
	result.typ = reflect.TypeOf(i)

	for result.val.Kind() == reflect.Ptr {
		result.val = result.val.Elem()
	}
	for result.typ.Kind() == reflect.Ptr {
		result.typ = result.typ.Elem()
	}

	result.generateInsertSQL()
	return result
}

func GetInsertSQL(i interface{}) string {
	ref := GetReflect(i)
	return ref.insertSQL
}

func TestRef(t *testing.T) {
	type Student struct {
		Age  int
		Name string
	}
	GetInsertSQL(new(Student))
}
