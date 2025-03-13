package main

import "fmt"

func boolen() {
	// 布尔类型，表示真假
	// 标识符bool
	// 字面量：true/false
	// 零值false
	var zero bool
	isBoy := true
	isGirl := false
	fmt.Println(zero, isBoy, isGirl)

	// 操作
	// 逻辑运算(与&&，或||，非!)

	// &&: 前后变量都为true，才会是true
	fmt.Println("&&")
	fmt.Println(true && true)   // true
	fmt.Println(true && false)  // false
	fmt.Println(false && true)  // false
	fmt.Println(false && false) // false

	// ||: 只要一个变量为true，就是true
	fmt.Println("||")
	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(false || true)  // true
	fmt.Println(false || false) // false

	// ! 取反，单目运算
	fmt.Println("!")
	fmt.Println(!true)  // false
	fmt.Println(!false) // true

	// 关系运算(==，!=)
	fmt.Println("---")
	fmt.Println(isBoy == isGirl) // false
	fmt.Println(isBoy != zero)   // true

	fmt.Printf("%T %t\n", isBoy, isBoy)
}
