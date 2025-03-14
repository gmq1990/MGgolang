package main

import "fmt"

// 函数参数的类型合并
// 连续同类型参数省略类型
func add1(a, b int) int {
	return a + b
}

// 可变行参，只能有一组，并且放在形参的最后
func addN(a, b int, args ...int) int {
	fmt.Println(a, b, args)
	fmt.Printf("%T\n", args) // []int
	var args_sum int
	for _, v := range args {
		args_sum += v
	}
	return a + b + args_sum
}

func calc(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		return addN(a, b, args...)
	}
	return -1
}

func arguments() {

	fmt.Println(addN(1, 4))       // 5
	fmt.Println(addN(1, 4, 5))    // 10
	fmt.Println(addN(1, 4, 5, 8)) // 18

	fmt.Println(calc("add", 1, 2))
	fmt.Println(calc("add", 1, 2, 5))
	fmt.Println(calc("add", 1, 2, 5, 8))

	args := []int{1, 2, 5, 6, 10}
	// 可变参数传递时需要解包。加上"..."
	fmt.Println(calc("add", 1, 2, args...))

	nums := []int{1, 2, 5, 8}
	// num[:1] + nums[2:]
	// nums[2:] 解包
	nums = append(nums[:1], nums[2:]...)
	fmt.Println(nums) // [1 5 8]

}
