package main

import "fmt"

func calc1(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

// 命名返回值
// 返回值中直接声明
func calc2(a, b int) (sum, sub, mult, div int) {
	sum = a + b
	sub = a - b
	mult = a * b
	div = a / b
	return
}

func returns() {
	a, b, _, _ := calc1(9, 3)
	fmt.Println(a, b)

	fmt.Println(calc2(5, 2))
}
