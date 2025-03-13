package main

import "fmt"

func fLoat() {
	// float32, float64
	// 字面量：十进制表示，科学技术表示

	var height float64 = 1.11
	fmt.Printf("%T %f\n", height, height)

	var weight float64 = 13.5e1
	fmt.Println(weight)

	// 操作
	// 算数运算(+-*/,++,--)
	fmt.Println(1.1 + 2.2)
	fmt.Println(1.1 - 2.2)
	fmt.Println(1.1 * 2.2)
	fmt.Println(1.1 / 2.2)

	// 计算机无法精确计算浮点数，结果有误差
	height++
	fmt.Println(height)
	height--
	fmt.Println(height)

	// 关系运算(>=, <=, >, <)
	fmt.Println(1.1 > 1.2)
	fmt.Println(1.1 < 1.2)
	fmt.Println(1.1 >= 1.2)
	fmt.Println(1.1 <= 1.2)

	// 赋值(=, +=, -=, *=, /=)
	height += 0.25
	fmt.Println(height)

	fmt.Printf("%T %T\n", 1.1, float32(1.1))

	// 保留5位，其中2位小数
	fmt.Printf("%5.2f\n", height) // _1.36

}
