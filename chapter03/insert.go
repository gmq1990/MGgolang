package main

import "fmt"

func insert_sort() {
	// 插入排序
	// 先指定某数为基数 每次和后面数比较 比基数大，就指定基数为后面的数
	// 初始化，假定第一个数为最大数maxNum
	nums := []int{1, 3, 5, 20, 30, 8, 9, 11}
	// 第一轮：
	// 1: maxNum =1; 1, 3 比较 => 大 => maxNum =3
	// 2: maxNum =3; 3, 5 比较 => 大 => maxNum =5
	// 3: maxNum =5; 5, 20 比较 => 大 => maxNum =20
	// 4: maxNum =20; 20, 30 比较 => 大 => maxNum =30
	// 5: maxNum =30; 30, 8 比较 => 小 => {1, 3, 5, 20, 8, 8, 9, 11} ，maxNum 不更新
	// 6: maxNum =30; 30, 9 比较 => 小 => {1, 3, 5, 20, 8, 9, 9, 11}， maxNum 不更新
	// 7: maxNum =30; 30, 11 比较 => 小 => {1, 3, 5, 20, 8, 9, 11, 11}， maxNum 不更新
	// maxNum -> 末尾 => {1, 3, 5, 20, 8, 9, 11, 30}
	// 第二轮：
	// 1: maxNum =1;...
	// ...

	for j := 0; j < len(nums)-1; j++ {
		maxNum := nums[0]
		fmt.Println("第", j, "轮：")
		// 第j轮，倒数第j+1个数将被置为最大数，只需前面n-j个数比较
		for i := 1; i < len(nums)-j; i++ {
			if nums[i] > maxNum {
				maxNum = nums[i]
			} else {
				nums[i-1] = nums[i]
			}
			fmt.Println(i, "maxNum=", maxNum, nums)
		}
		// 第j轮的结果，倒数第j+1个数被置为最大数
		nums[len(nums)-1-j] = maxNum
		fmt.Println("第", j, "轮结果", nums)
		fmt.Println()
	}
}
