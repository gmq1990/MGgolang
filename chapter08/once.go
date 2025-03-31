package main

import (
	"fmt"
	"sync"
)

func Once() {
	// sync.Once让函数只执行1次
	once := sync.Once{}
	for i := 0; i < 10; i++ {
		once.Do(func() {
			fmt.Println(i)
		})
	}
}
