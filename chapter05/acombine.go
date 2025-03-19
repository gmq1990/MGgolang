package main

import "fmt"

type Address2 struct {
	Region string
	Street string
	No     string
}

type User5 struct {
	ID   int
	Name string
	Addr *Address2 // 指针型嵌入
}

func acombine() {
	var me01 User5
	fmt.Printf("%#v\n", me01) //  Addr:(*main.Address2)(nil)

	me02 := User5{
		ID:   1,
		Name: "tom",
		Addr: &Address2{
			"beijing",
			"changanjie",
			"123",
		},
	}
	fmt.Printf("%#v\n", me02) // Addr:(*main.Address2)(0x140001020c0)

	fmt.Printf("%#v\n", me02.Addr.Street) // "changanjie"
}
