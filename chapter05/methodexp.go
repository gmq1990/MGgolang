package main

import "fmt"

type Dog4 struct {
	name string
}

func (dog Dog4) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

/*
// go解释器隐式生成一个指针接受者的方法
func (dog *Dog4) Call(){
	fmt.Printf("%s: 汪汪\n", dog.name)
}
*/

func (dog *Dog4) SetName(name string) {
	dog.name = name
}

func methodexp() {
	// 方法表达式
	m1 := Dog4.Call
	// m1 := (*Dog4).Call // 可以。go解释器生成。结果与Dog4.Call的结果一致
	m2 := (*Dog4).SetName
	// m2 := Dog4.SetName  // 不可以

	fmt.Printf("%T %T\n", m1, m2) // func(main.Dog4) func(*main.Dog4, string)
	// 方法表达式，使用时，要传入对象
	dog := Dog4{"豆豆"}
	m1(dog) // 豆豆

	m2(&dog, "小黑")
	m1(dog) // 小黑

	dog.SetName("小白")
	m1(dog) // 小白

	pdog := &Dog4{"doudou"}
	m1(*pdog) // 自动解引用
	m2(pdog, "xiaohei")
	m1(*pdog)
	pdog.SetName("xiaobai")
	m1(*pdog)
}
