package main

import (
	"fmt"
	"sort"
)

func sSort() {
	// 排序
	nums := []int{1, 3, 5, 9, 8, 7, 4}
	sort.Ints(nums)
	fmt.Println(nums)

	// 二分查找
	// 条件：有序数组中查找
	fmt.Println(sort.SearchInts(nums, 5)) // 3
	// 插入新值的位置
	fmt.Println(sort.SearchInts(nums, 6)) // 4

	// 判断值是否在序列中
	num := 6
	idx := sort.SearchInts(nums, num)
	if nums[idx] == num {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}
