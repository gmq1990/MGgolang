package main

import "fmt"

func randomInt() {
	channel := make(chan int, 1)
	slice := make([]int, 10)
	// 随机生成10个数
	for i := 0; i < 10; i++ {
		select {
		case channel <- 0:
		case channel <- 1:
		case channel <- 2:
		case channel <- 3:
		case channel <- 4:
		case channel <- 5:
		}
		slice[i] = <-channel
	}

	fmt.Println(slice)
}
