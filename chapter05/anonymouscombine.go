package main

import "fmt"

type Addr struct {
	Region string
	Street string
	No     string
}

type User4 struct {
	ID   int
	Name string
	Addr Addr
}

type Company struct {
	ID     int
	Name   string
	Addr   Addr
	Salary float64
}

type Employee struct {
	// 匿名嵌入
	User4
	Company
	Salary float64
	// Name   string
}

func anonymouscombine() {
	var me Employee
	fmt.Printf("%T %#v\n", me, me)

	me02 := Employee{
		User4: User4{
			ID:   1,
			Name: "tom",
			Addr: Addr{ // 指针型嵌入
				Region: "beijing",
				Street: "haidian",
				No:     "123",
			},
		},
		Salary: 10000,
	}
	fmt.Printf("%T %#v\n", me02, me02)

	fmt.Println(me02.User4.Name)
	me02.User4.Addr.Street = "changanjie"
	fmt.Println(me02.User4.Addr.Street)

	// 匿名结构体：在访问其属性时，可以省略结构名
	// 属性名冲突：以me02为优先，没有Name时，才访问User4里面的Name
	// 多个匿名结构体的内部元素冲突时，不能简写
	// fmt.Println(me02.Name)
	// me02.Addr.No = "321"      // User4被省略
	// fmt.Println(me02.Addr.No) // User4被省略
	fmt.Println(me02.User4.Name)
	me02.Company.Addr.No = "321"
	fmt.Println(me02.Company.Addr.No)
}
