package main

import "fmt"

func forCtrl() {
	// 索引 => 记录已经加到的n
	// 记录结果
	result := 0

	/*
		result += 1
		result += 2
		...
		result += 100
	*/
	/*
		i => 1,2,...,100
		result += i
	*/

	// 初始化子语句；条件子语句；后置子语句
	for i := 1; i <= 100; i++ {
		result += i
	}
	fmt.Println(result)

	result = 0
	i := 1
	// while
	for i <= 100 {
		result += i
		i++
	}
	fmt.Println(result)

	i = 0
	// while true
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 字符串，数组，切片，映射，管道
	desc := "还记得我吗"
	for i, ch := range desc {
		fmt.Printf("%d %T %q\n", i, ch, ch)
	}

}
