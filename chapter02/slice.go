package main

import "fmt"

func slice() {
	// 声明nums
	var nums []int
	fmt.Printf("%T\n", nums) // []int
	// 三个要素：指针、长度、容量。
	// 未初始化时，指针默认为nil，长度、容量默认为0
	fmt.Println(nums == nil) // true
	// 获取值、长度，容量
	// #的作用：类型+值
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // []int(nil) 0 0

	// 字面量
	// 长度可任意
	nums = []int{1, 2, 3}
	fmt.Println(nums)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // []int{1, 2, 3} 3 3

	nums = []int{1, 2, 3, 4}
	fmt.Println(nums)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // []int{1, 2, 3, 4} 4 4

	// 数组切片赋值
	var arrays [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums = arrays[1:7]
	fmt.Println(nums)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // []int{2, 3, 4, 5, 6, 7} 6 9

	// make函数: 初始化
	// var nums []int => []int(nil) 0 0
	nums = make([]int, 3)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // []int{0, 0, 0} 3 3

	nums = make([]int, 3, 5)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // []int{0, 0, 0} 3 5

	// 元素操作（增、删、改、查）
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	nums[2] = 10
	fmt.Println(nums)

	nums = append(nums, 1)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // []int{0, 0, 10, 1} 4 5

	nums = append(nums, 1, 1)
	// 扩展旧容量的1~1.5倍
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums)) // []int{0, 0, 10, 1, 1, 1} 6 10

	for _, v := range nums {
		fmt.Print(v)
	}
	fmt.Println()
	for i := 0; i <= len(nums)-1; i++ {
		fmt.Print(nums[i])
	}
	fmt.Println()

	// 切片操作
	fmt.Printf("%T\n", nums[1:5])

	n := nums[1:3:10]
	// n_cap - start
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n)) // []int []int{0, 10} 2 9

	n = nums[2:3]
	// nums_cap - start
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n)) // []int []int{10} 1 8

	// 切片副作用
	nums = make([]int, 3, 5)
	nums2 := nums[1:3]
	fmt.Println(nums, nums2) // [0 0 0] [0 0]
	nums2[0] = 1
	fmt.Println(nums, nums2) // [0 1 0] [1 0]

	nums2 = append(nums2, 3)
	fmt.Println(nums, nums2) // [0 1 0] [1 0 3]
	nums = append(nums, 5)
	fmt.Println(nums, nums2) // [0 1 0 5] [1 0 5]

	// 删除
	// copy的方式
	nums3 := []int{1, 2, 3}
	nums4 := []int{10, 20, 30, 40}
	copy(nums4, nums3)
	fmt.Println(nums3, nums4) // [1 2 3] [1 2 3 40]

	nums4 = []int{10, 20, 30, 40}
	copy(nums3, nums4)
	fmt.Println(nums3, nums4) // [10 20 30] [10 20 30 40]

	// 索引为0或最后一个的删除
	nums5 := []int{1, 2, 3, 4, 5}
	fmt.Println(nums5[1:])
	fmt.Println(nums5[:len(nums5)-1])

	// 索引为中间的删除
	copy(nums5[2:], nums5[3:])
	fmt.Println(nums5)
	fmt.Println(nums5[:len(nums5)-1])

	// 堆栈，队列
	// 堆栈：每次添加到队尾，移除元素在队尾（先进后出）
	// 队列：每次添加到队尾，移除元素在队头（先进先出）
	queue := []int{}
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 4)

	queue = queue[1:]
	fmt.Println(queue) // [2 3 4]
	queue = queue[1:]
	fmt.Println(queue) // [3 4]

	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println(stack)

	stack = stack[:len(stack)-1]
	fmt.Println(stack) // [1 2]
	stack = stack[:len(stack)-1]
	fmt.Println(stack) // [1]

	points := [][]int{}
	points2 := make([][]int, 0)
	fmt.Println(points2)

	points = append(points, []int{1, 2, 3})
	points = append(points, []int{3, 4, 0})
	points = append(points, []int{3, 4, 0, 2, 4, 5})
	fmt.Println(points)
	fmt.Println(points[0])
	fmt.Println(points[0][1])

}
