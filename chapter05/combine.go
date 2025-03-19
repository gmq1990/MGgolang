package main

import "fmt"

type Address struct {
	Region string
	Street string
	No     string
}

type User3 struct {
	ID   int
	Name string
	Addr *Address
}

func combine() {
	// 定义嵌套结构体
	// 方法1
	var me01 User
	fmt.Printf("%#v\n", me01) // main.User(nil)

	// 方法2
	addr := Address{"shanghai", "jingan", "231"}
	me02 := User3{
		ID:   1,
		Name: "tom",
		Addr: &addr,
	}
	fmt.Printf("%#v\n", me02) // main.User3{ID:1, Name:"tom", Addr:main.Address{Region:"shanghai", Street:"jingan", No:"231"}}

	// 方法3
	me03 := User3{
		ID:   1,
		Name: "tom",
		Addr: &Address{
			Region: "beijing",
			Street: "haidian",
			No:     "123",
		},
	}

	fmt.Printf("%T %#v\n", me03, me03) // main.User3 main.User3{ID:1, Name:"tom", Addr:main.Address{Region:"beijing", Street:"haidian", No:"123"}}

	// 访问嵌入结构体的属性
	me03.Addr.Region = "nanjing"
	me03.Addr.Street = "xxxxx"
	me03.Addr.No = "432"
	fmt.Printf("%#v\n", me03) // main.User3{ID:1, Name:"tom", Addr:main.Address{Region:"nanjing", Street:"xxxxx", No:"432"}}
}
