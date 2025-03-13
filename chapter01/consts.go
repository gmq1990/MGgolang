package main

import "fmt"

func consts() {
	// 使用const定义常量，常量不能被修改
	// const类型，必须附值
	// 约定，常量使用大写字母
	const NAME string = "kk"

	fmt.Println(NAME)

	// 省略类型
	const aa = 11
	fmt.Println(aa)

	// 定义多个常量，类型相同
	const bb, cc string = "bb", "cc"
	fmt.Println(bb, cc)

	// 定义多个常量，类型不同
	const (
		dd int    = 44
		ee string = "ee"
	)
	fmt.Println(dd, ee)

	// 定义多个常量，省略类型。将会默认附值为第一个常量的类型、值
	const (
		FF int = 1
		GG     // 1
		HH     // 1
	)
	fmt.Println(FF, GG, HH)

	// 枚举, const+iota
	const (
		E1 int = iota // 0
		E2            // 1
		E3            // 2
	)
	const (
		E4 int = iota + 1 // 1
		E5                // 2
		E6                // 3
	)
	fmt.Println(E1, E2, E3)
	fmt.Println(E4, E5, E6)

}
