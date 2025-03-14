package main

import "fmt"

func bubble_sort() {
	heights := []int{10, 6, 7, 9, 5}
	// 先把最高的人移到最后
	// 从第一个人开始，两两开始比较，如果前面的人高，这两人就交换位置
	// 1: 10, 6 => 交换 => 6,10,7,9,5
	// 2: 10, 7 => 交换 => 6,7,10,9,5
	// 3: 10, 9 => 交换 => 6,7,9,10,5
	// 4: 10, 5 => 交换 => 6,7,9,5,10
	// 第二轮
	// 1: 6, 7 => 不交换 => 6,7,9,5,10
	// 2: 7, 9 => 不交换 => 6,7,9,5,10
	// 3: 9, 5 => 交换 => 6,7,5,9,10
	// 4: 9, 10 => 不交换 => 6,7,5,9,10
	// ...
	// n 个人比较 n-1 次
	// 0 -> n-1

	for j := 0; j < len(heights)-1; j++ {
		fmt.Println("第", j, "轮：")
		// 第j轮的结果，就确定了末尾j+1个数已排好
		for i := 0; i < len(heights)-1-j; i++ {
			if heights[i] > heights[i+1] {
				println("交换: ", heights[i], heights[i+1])
				// tmp := a; a = b; b = tmp
				// a, b := b, a
				heights[i], heights[i+1] = heights[i+1], heights[i]
			}
			fmt.Println(i, "交换完毕", heights)
		}
		fmt.Println("第", j, "轮交换完毕")
		fmt.Println()
	}

}
