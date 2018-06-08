package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string
}

func (this *MyStruct) GetName(str string) string {
	this.name = str
	return this.name
}

func main()  {
	// 对象
	s := "this is string"

	// 获取对象类型 (string)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))
}