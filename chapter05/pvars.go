package main

import "fmt"

type Address5 struct {
	Region string
	Street string
	No     string
}

type User8 struct {
	ID   int
	Name string
	Addr *Address5 // 指针型嵌入
}

// go需要定义一个函数，来模拟构造函数
func NewUser(id int, name string, region, street, num string) *User8 {
	return &User8{
		ID:   id,
		Name: name,
		Addr: &Address5{region, street, num},
	}
}

func pvars() {
	// me := User8{}
	// me2 := me
	// me2.Name = "kk"
	// me2.Addr.Street = "nanjinglu" // panic：Addr为nil，初始化前不能赋值
	// fmt.Printf("%#v\n", me)
	// fmt.Printf("%#v\n", me2)

	me3 := User8{
		ID:   1,
		Name: "tom",
		Addr: &Address5{"beijing", "changanjie", "123"},
	}
	me4 := me3
	me4.Name = "kk"
	me4.Addr.Street = "nanjinglu"
	fmt.Printf("%#v\n", me3) // Name:"", &Address5{"beijing", "changanjie", "123"}
	fmt.Printf("%#v\n", me4) // Name:"kk", &Address5{"beijing", "changanjie", "123"}

	tom := NewUser(2, "tom", "beijing", "haidianlu", "123")
	fmt.Printf("%#v\n", tom)

}
