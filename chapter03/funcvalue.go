package main

import "fmt"

func changeInt(a int) {
	a = 100
}

func changeIntPointer(p *int) {
	*p = 100
}

func changeSlice(s []int) {
	s[0] = 100
}

func funcvalue() {
	num := 1
	changeInt(num)
	// 值类型的传递，是复制，不是引用。
	// num没有变
	fmt.Println(num) // 1

	nums := []int{1, 2, 3}
	changeSlice(nums)
	// 引用类型的传递，是指向同一个内存空间。
	// nums改变了
	fmt.Println(nums) // [100 2 3]

	// 传递指针，指向同一个内存空间。
	// num改变了
	changeIntPointer(&num)
	fmt.Println(num) // 100
}
