package main

import (
	"fmt"
	"time"
)

// struct类型
type User1 struct {
	// User1的模式（属性）
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Remark   string
}

func sStruct() {
	// var counter Counter
	var me User1
	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me)

	// 实现初始化
	// 所有元素按顺序输入
	var me2 User1 = User1{
		1,
		"tom",
		time.Now(),
		"beijing",
		"15000000000",
		"tom",
	}
	fmt.Printf("%#v\n", me2)

	// 零值
	var me3 User1 = User1{} // var me3 User1
	fmt.Printf("%#v\n", me3)

	// 指定元素初始化
	// 没有顺序要求
	var me4 User1 = User1{ID: 1, Name: "tom", Addr: "beijing", Tel: "1223322443"}
	fmt.Printf("%#v\n", me4)

	var pointer *User1
	fmt.Printf("%T\n", pointer)  // *main.User1
	fmt.Printf("%#v\n", pointer) //	(*main.User1)(nil)

	// 指针赋值
	var pointer2 *User1 = &me2
	fmt.Println(pointer2)

	var pointer3 *User1 = &User1{}
	fmt.Printf("%#v\n", pointer3)
	// 使用new
	// 与&User1{}一个含义
	var pointer4 *User1 = new(User1)
	fmt.Printf("%#v\n", pointer4)

	// 指针new，可以用于其他类型
	// map[], []slice类型除外 => 使用make来初始化
	var pointer5 *int = new(int)
	fmt.Printf("%#v\n", pointer5)
}
