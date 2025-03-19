package main

import (
	"fmt"
)

type Address6 struct {
	Region string
	Street string
	No     string
}

// 方法的命名嵌入
// Address()命名成String()
func (addr Address6) String() string {
	return fmt.Sprintf("%s %s %s\n", addr.Region, addr.Street, addr.No)
}

type User9 struct {
	ID   int
	Name string
	Addr Address6
}

// 方法的命名嵌入
// User()命名成String()
func (user User9) String() string {
	return fmt.Sprintf("[%d]%s: %s\n", user.ID, user.Name, user.Addr)
}

func mcombin() {
	var u User9 = User9{
		ID:   1,
		Name: "tom",
		Addr: Address6{"beijing", "haidianlu", "001"},
	}

	// 自动调用String()方法
	// 没有String方法就打印属性值
	fmt.Println(u)
	fmt.Println(u.Addr)
}
