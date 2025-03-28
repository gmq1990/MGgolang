package main

import (
	"encoding/json"
	"fmt"
)

func jJson() {
	/*
		json.Marshal() // 序列化，内存 => 字符串/字节切片
		json.Unmarshal()  // 反序列化，字符串/字节切片 => 内存
	*/

	names := []string{"张三", "李四", "王五", "小白"}

	users := []map[string]string{{"name": "张三", "addr": "beijing"}, {"name": "李四", "addr": "shanghai"}, {"name": "王五", "addr": "shenzhen"}}

	// marshal
	// 传入某个类型的值
	// 返回一个字节切片，和error
	bytes, err := json.Marshal(names)
	if err == nil {
		fmt.Println(bytes)
		fmt.Println(string(bytes))
	}
	// pretty 序列化
	bytes, err = json.MarshalIndent(names, "", "\t")
	if err == nil {
		fmt.Println(string(bytes))
	}

	// unmarshal
	// 传入一个字节切片，和一个指针。
	// 返回一个error
	// 接收者的类型必须和bytes中的json数据结构体相同
	var names02 []string
	err = json.Unmarshal(bytes, &names02)
	if err == nil {
		fmt.Println(names02)
	}
	// pretty 序列化
	bytes, err = json.MarshalIndent(users, "", "\t")
	if err == nil {
		// fmt.Println(bytes)
		fmt.Println(string(bytes))
	}

	var users02 []map[string]string

	err = json.Unmarshal(bytes, &users02)
	if err == nil {
		fmt.Println(users02)
	}

	// Valid()，验证格式是否符合json
	fmt.Println(json.Valid([]byte("[]+"))) // false
	fmt.Println(json.Valid([]byte("[]")))  // true
}
