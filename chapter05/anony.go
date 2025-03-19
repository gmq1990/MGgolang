package main

import "fmt"

func anony() {
	// 匿名结构体
	var me struct {
		ID   int
		Name string
	}

	fmt.Printf("%T\n", me)  // struct { ID int; Name string }
	fmt.Printf("%#v\n", me) // struct { ID int; Name string }{ID:0, Name:""}

	// 访问匿名结构体的属性
	me.ID = 3
	me.Name = "tom"
	// me.Addr无法访问 ，因为未定义
	fmt.Printf("%#v\n", me) // struct { ID int; Name string }{ID:3, Name:"tom"}

	// 使用字面量来初始化
	var me2 = struct {
		ID   int
		Name string
	}{1, "tom"}
	fmt.Printf("%#v\n", me2)

	// 匿名结构体使用场景
	// 加载配置；传递数据
}
