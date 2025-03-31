package main

import (
	"fmt"
	"sync"
)

func objpool() {
	// 对象池
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("new")
			return 1
		},
	}
	// 获取1个对象
	x := pool.Get() // new
	fmt.Println(x)  // 1
	// 放回对象池
	pool.Put(x)
	// 放回后再次调用
	// 对象池中有对象，不用再创建
	x = pool.Get() // no "new"
	fmt.Println(x)
}
