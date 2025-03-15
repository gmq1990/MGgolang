package main

import "fmt"

// 闭包
func closures() {

	add2 := func(n int) int {
		return n + 2
	}
	add10 := func(n int) int {
		return n + 10
	}
	add100 := func(n int) int {
		return n + 100
	}

	// 闭包：
	// 匿名函数中return的函数，引用了外部作用域里的变量

	// 定义一个函数，传入一个基本值，返回一个函数。
	// addBase为创造函数的函数（外部函数）
	// 返回的函数与上面的addN函数一样
	addBase := func(base int) func(int) int {
		// 被return的函数为内部函数
		// 内部函数调用了addBase里的形参base，形成闭包
		// 这个base变量，就和内部函数赋值给某个变量后的生命周期一致了（后面的add2，add10，add100）
		return func(n int) int {
			return base + n
		}
	}
	// fmt.Println(addBase(2)(3))
	// fmt.Println(addBase(10)(3))
	// fmt.Println(addBase(100)(3))
	add2 = addBase(2)
	add10 = addBase(10)
	add100 = addBase(100)
	fmt.Println(add2(3))
	fmt.Println(add10(3))
	fmt.Println(add100(3))

	// 用addBase生成一个add8函数
	// base的值为最后被修改时的值
	add8 := addBase(8)
	fmt.Printf("%T\n", add8) // func(int) int
	fmt.Println(add8(10))    // 18

	// add2，add10，add100存在，addBase的参数就一直在
}
