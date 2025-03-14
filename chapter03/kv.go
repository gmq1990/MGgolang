package main

import "fmt"

func kv() {
	users := map[string]int{"张三": 9, "李四": 8, "王五": 9}
	keySlice := make([]string, len(users))

	valueSlice := make([]int, 0)

	i := 0
	for k, v := range users {
		// 指定了slice长度
		keySlice[i] = k
		i++
		// 未指定slice长度
		valueSlice = append(valueSlice, v)

	}
	fmt.Println(keySlice, valueSlice)

	for v := range users {
		fmt.Println(v) // 打印值：key
	}

	for _, v := range users {
		fmt.Println(v) // 打印值：value
	}
}
