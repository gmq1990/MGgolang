package main

import "fmt"

// 递归
func addN1(n int) int {
	if n == 1 {
		return 1
	}
	return n + addN1(n-1)
}

func factorial(n int) uint64 {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	return uint64(n) * factorial(n-1)
}

func recursion() {
	fmt.Println(addN1(5))

	fmt.Println(factorial(20))
}

// addN1(5) => 5 + addN1(4)
// addN1(4) => 4 + addN1(3)
// addN1(3) => 3 + addN1(2)
// addN1(2) => 2 + addN1(1)
// addN1(1) => 1 + addN1(0)
// addN1(0) => 0 + addN1(-1)
// ... 死循环
// 设置一个终止点。令n=1时，结束死循环
