package main

import "fmt"

type Dog2 struct {
	name string
}

func (dog Dog2) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

func (dog *Dog2) SetName(name string) {
	dog.name = name
}

func methodtype() {
	dog := Dog2{"豆豆"}
	// 方法值:　绑定了变量dog（接收者）
	// 值接收者方法，赋值后不会再改变结果
	m1 := dog.Call         // 赋值 dog给m1。拷贝
	fmt.Printf("%T\n", m1) // func()
	m1()                   // 豆豆: 汪汪

	dog.SetName("小黑")
	dog.Call() // 小黑
	m1()       // 豆豆

	pdog := &Dog2{"豆豆"}    // 指针
	m2 := pdog.Call        // pdog会自动解引用，拷贝
	fmt.Printf("%T\n", m2) // func()
	m2()

	pdog.SetName("小黑")
	pdog.Call() // 小黑
	m2()        // 豆豆

	// 如果call2是指针接受者，将拷贝地址。那么m1、m2会随着dog、pdog改变而改变

}
