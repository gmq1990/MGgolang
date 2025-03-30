package main

import (
	"fmt"
	"time"
)

func rwchannel() {
	var channel chan int = make(chan int, 5)

	var rchannel <-chan int = channel // 只读管道

	var wchannel chan<- int = channel // 只写管道

	go func(channel <-chan int) {
		fmt.Println(<-channel)
	}(channel)

	go func(channel chan<- int) {
		channel <- 0
	}(channel)

	// 底层用同一个channel
	wchannel <- 1
	fmt.Println(<-rchannel)
	time.Sleep(time.Second)

}
