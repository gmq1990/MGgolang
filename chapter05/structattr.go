package main

import (
	"fmt"
	"time"
)

// struct类型
type User2 struct {
	// User1的模式（属性）
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Remark   string
}

func structattr() {
	var me User2 = User2{
		1,
		"tom",
		time.Now(),
		"beijing",
		"15000000000",
		"tom",
	}
	fmt.Println(me.ID, me.Name, me.Tel)

	// 修改属性的值
	me.Tel = "14200000000"
	fmt.Printf("%#v\n", me.Tel)

	//
	me2 := &User2{
		ID:   2,
		Name: "tom",
	}
	// 指针访问属性时，两种方式
	fmt.Printf("%#v\n", (*me2).ID) // 第一种
	fmt.Printf("%#v\n", me2.Name)  // 第二种

}
