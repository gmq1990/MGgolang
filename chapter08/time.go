package main

import (
	"fmt"
	"time"
)

func tTime() {
	//
	fmt.Println(time.Now())
	channel := time.After(3 * time.Second)
	// time.After只读1次
	// 用途：延迟执行
	fmt.Println(<-channel)
	// fmt.Println(<-channel) // panic

	// time.Tick可以一直读
	// 用途：每个间隔定时执行一次
	ticker := time.Tick(3 * time.Second)
	fmt.Println(<-ticker)
	fmt.Println(<-ticker)
	fmt.Println(<-ticker)
}
