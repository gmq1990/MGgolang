package main

import "fmt"

func add2(a, b int) int {
	return a + b
}

// 格式化，将传递的数据按照每行打印，或按照一行按 | 分割
// 函数作为参数
func print(callback func(...string), args ...string) {
	fmt.Println("print函数输出：")

	callback(args...)
}

func list(args ...string) {
	for i, v := range args {
		fmt.Println(i, ":", v)
	}
}

func tType() {
	// 函数其实可以认为也是一种变量
	// 参数数量、类型，返回值的数量、类型，统称为函数的签名
	fmt.Printf("%T\n", add2) // func(int, int) int

	f := add
	// 变量可以传递给别的函数作为参数
	// var f func(int, int) int = add
	fmt.Printf("%T\n", f) // func(int, int) int

	fmt.Println(f(1, 4)) // 5

	// 函数作为参数，传递给另一个函数
	print(list, "A", "C", "E")

	// 匿名函数
	// 函数内部定义一个只在函数内使用的函数，赋值给一个变量
	sayHello := func(name string) {
		fmt.Println("Hello", name)
	}

	sayHello("KK")
	sayHello("tom")

	// 匿名函数：只使用一次的函数
	func(name string) {
		fmt.Println("Hi", name)
	}("KK")

	values := func(args ...string) {
		for _, v := range args {
			fmt.Println(v)
		}
	}

	print(values, "A", "B", "C")

	print(func(args ...string) {
		for _, v := range args {
			fmt.Print(v, "\t")
		}
		fmt.Println()
	}, "A", "B", "E")
}
