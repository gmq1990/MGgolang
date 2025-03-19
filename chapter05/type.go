package main

import "fmt"

// 自定义类型
// 定义成自己需要的格式
// 相当于在基础库的类型基础上构建新的类型
type Counter int

type User map[string]string

type Callback func(...string)

func tType() {
	// int格式，counter可执行int方式的计算
	var counter Counter = 20
	counter += 10
	fmt.Println(counter)

	// map[string]string格式，可定义成User
	me := make(User)
	me["name"] = "tom"
	me["addr"] = "beijing"

	fmt.Println(me)
	fmt.Printf("%T %T\n", counter, me)

	// func(...string)格式，可定义成Callback
	var list Callback = func(args ...string) {
		for i, v := range args {
			fmt.Println(i, ":", v)
		}
	}

	list("a", "b", "c")

	// 自定义类型和基础类型 不能相互运算，需要转换类型
	var counter2 int = 10
	fmt.Println(int(counter) > counter2)
}
