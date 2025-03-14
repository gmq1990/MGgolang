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
	// 返回的函数与上面的addN函数一样
	addBase := func(base int) func(int) int {
		// 下面的函数中，base+n引用了closure()里的变量base，形成闭包
		// 这个base变量，被匿名函数引用后，就一直存在，直到closures()结束
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

	// addBase函数仍然在，base就一直在
	// END: closures()函数结束时，才会销毁base
}
