package main

import (
	"fmt"
	"reflect"
)

func rReflect() {
	var i string = "1"
	// 反射，"%T"
	fmt.Printf("%T\n", i)
	// 底层代码
	var typ reflect.Type = reflect.TypeOf(i)
	fmt.Println(typ)
	fmt.Println(typ.Kind() == reflect.String) // true  reflect.String = 24
}
