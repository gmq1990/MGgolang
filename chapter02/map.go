package main

import "fmt"

func mMap() {
	var scores map[string]int // nil 映射

	fmt.Printf("%T %#v\n", scores, scores) // map[string]int map[string]int(nil)
	fmt.Println(scores == nil)             // true

	// 字面量

	// 初始化 方法1
	// scores := map[string]int{}
	scores = map[string]int{"张三": 8, "李四": 9, "王五": 10}
	fmt.Println(scores)
	// 初始化 方法2
	scores = make(map[string]int)
	scores = map[string]int{"张三": 8, "李四": 9, "王五": 10}
	fmt.Println(scores)

	// 增，删，改，查
	// key
	fmt.Println(scores["张三"]) // 8
	fmt.Println(scores["kk"]) // 0

	// 查：判断key是否存在
	if _, ok := scores["kk"]; !ok {
		// 如果不存在
		fmt.Println(ok) // false
	}
	// 改
	scores["张三"] = 1
	fmt.Println(scores) // map[张三:1 李四:9 王五:10]
	// 增
	scores["kk"] = 3
	fmt.Println(scores) // map[kk:3 张三:1 李四:9 王五:10]
	// 删
	delete(scores, "kk")
	fmt.Println(scores) // map[张三:1 李四:9 王五:10]

	// 拿到元素数量
	fmt.Println(len(scores))
	for k, v := range scores {
		fmt.Println(k, v)
	}

	// key 至少要有==、!=运算，bool、整数、字符串、数组可作为key
	// value => 任意类型 （含slice map）
	var users map[string](map[string]string)
	users = map[string]map[string]string{"张三": {"地址": "北京", "分数": "8", "联系方式": "2123123"}}
	fmt.Printf("%T %#v\n", users, users)
	// 增
	users["张三"]["年龄"] = "23"
	fmt.Println(users)
	// 删
	delete(users["张三"], "年龄")
	fmt.Println(users)
	// 查
	if v, ok := users["张三"]["年龄"]; ok {
		fmt.Println(v)
	}
	// 改
	users["张三"]["分数"] = "9"
	fmt.Println(users)
}
