package main

import "fmt"

func arrays() {
	var nums [10]int
	var t2 [5]bool
	var t3 [3]string
	fmt.Printf("%T\n", nums)
	fmt.Println(t2)
	fmt.Printf("%q\n", t3)

	// 字面量
	nums = [10]int{10, 20, 30} // 前三个
	fmt.Println(nums)

	// 使用索引，给指定位置赋值
	nums = [10]int{0: 10, 9: 20} // 第0个，第9个
	fmt.Println(nums)

	// 不指定长度，以元素个数推导出长度
	nums = [...]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1}
	fmt.Println(nums)

	nums2 := [10]int{10, 20, 30}
	fmt.Printf("%T %#v\n", nums2, nums2)

	nums3 := [...]int{1, 2}
	fmt.Printf("%T %#v\n", nums3, nums3)

	nums4 := [...]int{0: 1, 13: 13}
	fmt.Printf("%T %#v\n", nums4, nums4)

	// 操作
	nums5 := [3]int{1, 3, 4}
	nums6 := [3]int{2, 3, 4}
	// 关系运算
	fmt.Println(nums5 == nums6)

	// 获取数组的长度
	fmt.Println(len(nums5))

	// 索引 0,1,2,...,len(n)-1
	fmt.Println(nums4[0], nums4[3])
	nums4[0] = 1000
	fmt.Println(nums4)

	for index, value := range nums4 {
		fmt.Printf("%d\t%d\n", index, value)
	}

	// 切片
	fmt.Printf("%T\n", nums4[0:10]) // []int
	fmt.Printf("%v\n", nums4[1:3:14])

}
