package main

import "fmt"

type Dog struct {
	name string
}

// 结构体的方法
func (dog Dog) Call() {
	fmt.Printf("%s: 汪汪\n", dog.name)
}

// 为内部属性设置访问方法
func (dog Dog) SetName(name string) {
	dog.name = name
}

func (dog *Dog) PsetName(name string) {
	dog.name = name
}

func method() {
	dog := Dog{"豆豆"}

	dog.Call()

	// 修改属性从新调用方法
	// dog.name = "小黑"
	// dog.Call()

	dog.SetName("小黑") //拷贝了一份dog给setname方法
	dog.Call()        // 豆豆: 汪汪

	(&dog).PsetName("小白")   // 取引用
	dog.PsetName("xiaohei") // 自动取引用，语法糖
	dog.Call()

	pdog := &Dog{"豆豆"} // 指针型
	(*pdog).Call()     // 解引用
	pdog.PsetName("小黑")
	pdog.Call() // 自动解引用，语法糖

	//var pdog2 *Dog // 未初始化，不能调用方法
	//fmt.Println(pdgod2.Psetname("test"))

}
