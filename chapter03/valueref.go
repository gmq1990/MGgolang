package main

import "fmt"

func valueref() {
	array := [3]string{"A", "B", "C"}
	slice := []string{"A", "B", "C"}

	arrayA := array
	sliceA := slice

	arrayA[0] = "Z"
	sliceA[0] = "Z"
	// 值类型： 被赋值的新变量 => 改变 => 不影响 原变量
	// 引用类型：被赋值的新变量 => 改变 => 影响 原变量
	// 值类型：int系列、float系列、bool、string、数组、结构体
	// 引用类型：slice切片、管道channel、接口interface、map、函数
	// 指针类型：不属于引用类型，是一种特殊的值类型
	// 值类型的特点是：变量直接存储值，内存通常在栈中分配
	// 引用类型的特点是：变量存储的是一个地址，这个地址对应的空间里才是真正存储的值，内存通常在堆中分配
	fmt.Println(arrayA, array) // [Z B C] [A B C]
	fmt.Println(sliceA, slice) // [Z B C] [Z B C]

	// %p 取地址
	fmt.Printf("%p %p\n", &arrayA, &array) // 不同
	fmt.Printf("%p %p\n", sliceA, slice)   // 相同
}
