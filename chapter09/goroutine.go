package main

import (
	"fmt"
	"time"
)

func goroutine() {
	names := make([]string, 0)
	dnames := make(map[string]string)

	go func([]string, map[string]string) {
		time.Sleep(time.Second * 2)
		fmt.Println(dnames)
		fmt.Println(names)
	}(names, dnames)
	// 在goroutine外部赋值，被传到了goroutine内部
	names = append(names, "xiaobai")
	dnames["1"] = "xiaohei"

	time.Sleep(time.Second * 4)
}
