package main

import (
	"fmt"
	"sync"
)

func mMap() {
	var users sync.Map

	users.Store(10, "kk")
	users.Store(20, "xiaobai")
	// 查询操作
	value, ok := users.Load(10)
	// value 为any类型
	fmt.Println(value.(string), ok) // kk true

	value, ok = users.Load(30)
	fmt.Println(value, ok) // <nil> false

	// 删除操作
	users.Delete(10)
	value, ok = users.Load(10)
	fmt.Println(value, ok) // <nil> false
}
