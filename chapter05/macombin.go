package main

import "fmt"

type User10 struct {
	id   int
	name string
}

func (user User10) GetID() int {
	return user.id
}

func (user User10) GetName() string {
	return user.name
}

func (user *User10) SetID(id int) {
	user.id = id
}

func (user *User10) SetName(name string) {
	user.name = name
}

type Employee2 struct {
	User10
	Salary float64
	name   string
}

func (employee Employee2) GetName() string {
	return employee.name
}

func (employee Employee2) SetName(name string) {
	employee.name = name
}

func macombin() {
	var me = Employee2{
		User10: User10{1, "tom"},
		Salary: 1000,
		name:   "xk",
	}

	// 匿名嵌入，可省略结构体名。
	// 方法如果有重复，则调用第一层的方法
	fmt.Println(me.User10.GetName()) // tom
	fmt.Println(me.GetName())        // xk
	me.SetName("jerry")
	fmt.Println(me.GetName()) // xk，SetName()不是指针引入
}
