package main

import "fmt"

func marrays() {

	var marrays [3][2]int
	// 长度为2的int类型数组 [2]int
	fmt.Println(marrays)
	fmt.Println(marrays[0])
	fmt.Println(marrays[0][0])
	marrays[0] = [2]int{2, 2}
	fmt.Println(marrays)

	marrays = [3][2]int{{1, 2}, {3, 4}}
	fmt.Println(marrays)
}
