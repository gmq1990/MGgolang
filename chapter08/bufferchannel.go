package main

import "fmt"

func bufferchannel() {
	channel := make(chan string, 2)

	fmt.Println(len(channel)) // 0
	channel <- "x"
	fmt.Println(len(channel)) // 1
	channel <- "y"
	fmt.Println(len(channel)) // 2
	// channel <- "z"

	fmt.Println(<-channel)
	fmt.Println(len(channel)) // 1
	fmt.Println(<-channel)
	fmt.Println(len(channel)) // 0

	channel <- "z"
	channel <- "a"
	close(channel)
	// fmt.Println(<-channel)
	// channel <- "z"

	for ch := range channel { // 需要在例程中关闭channel，否则死锁
		fmt.Println(ch)
	}
}
