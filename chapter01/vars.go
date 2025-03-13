package main

import "fmt"

// 包级别的变量
var version string = "1.0"

func vars() {
	// 定义string变量
	/*
		变量名需要满足标识符命名规则
		1. 必须由非空unicode字符串、数字、_组成
		2. 不能数字开头
		3. 不能用go关键字（25个）

		4. 避免与go预定义标识符冲突，如”true“
		5. 驼峰式命名
		6. 标识符区分大小写
	*/
	var me string
	// 零值打印
	fmt.Println(me)

	me = "kk"
	fmt.Println(me)

	// 相同类型变量
	var name, user string = "kk", "hah"
	fmt.Println(name, user)

	// 不同类型变量(推荐使用)
	var (
		age    int     = 35
		height float64 = 1.80
	)
	fmt.Println(age, height)

	// 以附值方式来自动推导类型
	var (
		ss = "kk"
		aa = 33
	)
	fmt.Println(ss, aa)

	// 简短声明（只能函数内）
	isBoy := true
	fmt.Println(isBoy)

	ss, aa, isBoy = "1", 23, false
	fmt.Println(ss, aa, isBoy)

	// 交换变量值
	var s string
	s, ss = ss, s
}
