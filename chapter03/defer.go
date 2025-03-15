package main

import "fmt"

func defers() {
	// 延迟执行：
	// 当函数退出的时候执行defer
	defer func() {
		fmt.Println("defer01")
	}()

	defer func() {
		fmt.Println("defer02")
	}()

	fmt.Println("main over")

	// 执行结果:
	// main over
	// defer02
	// defer01
	// defer修饰符，符合“栈”的设定。后声明的先执行。
}
