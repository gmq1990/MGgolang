package main

import "fmt"

type Address3 struct {
	Region string
	Street string
	No     string
}

type User6 struct {
	ID   int
	Name string
	Addr Address3
}

type Company1 struct {
	ID     int
	Name   string
	Addr   Addr
	Salary float64
}

type Employee1 struct {
	// 匿名嵌入：结构体指针
	*User6
	Salary float64
	Name   string
}

func pacombine() {
	var me01 Employee1
	fmt.Printf("%#v\n", me01) // User6:(*main.User)(nil)

	// 指针型匿名结构体，初始化
	me02 := Employee1{
		User6: &User6{
			ID:   1,
			Name: "tom",
			Addr: Address3{"beijing", "changanjie", "123"},
		},
		Salary: 1000,
		Name:   "xxtomxx",
	}
	fmt.Printf("%#v\n", me02)    // User6:(*main.User6)(0x1400012a000)
	fmt.Println(me02.Name)       // xxtomxx。默认先访问外面一层的Name
	fmt.Println(me02.Addr)       // {beijing changanjie 123}
	fmt.Println(me02.User6.Name) // tom
	fmt.Println(me02.User6.Addr) // {beijing changanjie 123}

}
