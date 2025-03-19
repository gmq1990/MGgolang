package main

import "fmt"

type Dog3 struct {
	name string
}

func (dog *Dog3) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

func (dog *Dog3) SetName(name string) {
	dog.name = name
}

func methodtypeV2() {
	dog := Dog3{"豆豆"}
	// 方法值:　绑定了变量dog（接收者）
	// 指针接收者方法，赋值后可改变结果
	m1 := dog.Call         // dog 取引用。赋值dog的指针
	fmt.Printf("%T\n", m1) // func()
	m1()                   // 豆豆: 汪汪

	dog.SetName("小黑")
	dog.Call() // 小黑
	m1()       // 小黑

	pdog := &Dog3{"豆豆"}
	m2 := pdog.Call        // 赋值pdog的指针
	fmt.Printf("%T\n", m2) // func()
	m2()

	pdog.SetName("小黑")
	pdog.Call() // 小黑
	m2()        // 小黑
}
