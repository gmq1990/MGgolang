package main

import "fmt"

func debug() {
	// 使用fmt.println来打印运行结果，以排查错误
	var age = 30

	age = age + 2
	fmt.Println("明年：", age)
	age = age + 1
	fmt.Println("后年：", age)

	// 不换行
	fmt.Print("打印第一行")
	fmt.Print("打印第二行")

	// 换行
	fmt.Println("打印第一行")
	fmt.Println("打印第二行")

	// 格式化打印
	fmt.Printf("%T,%s,%T,%d", "KK", "KK", 1, 1)
}
