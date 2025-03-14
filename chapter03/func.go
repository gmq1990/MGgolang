package main

import (
	"fmt"
)

// 定义(无参，无返回值)
func sayHelloWorld() {
	fmt.Println("Hello World!")
}

// 定义有参函数
func sayHi(name string) { // name: 形参
	fmt.Println("你好：", name)
}

func add(a int, b int) int {
	return a + b
}

func function() {
	// 对某个代码块起一个名字 => 函数名
	// 定义一个函数
	// 参数：内部要使用的变量
	// 返回值：外部要使用的变量
	// 目的：代码复用

	// 打印标识符，sayHelloWorld类型
	fmt.Printf("%T\n", sayHelloWorld) // func()

	// 调用
	sayHelloWorld()

	sayHi("kk") // "kk": 实参

	name := "tom"
	sayHi(name) // 实参name的值传给了行参name

	rt := add(1, 5)
	fmt.Println(rt)

}
