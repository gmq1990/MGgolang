package main

import (
	"fmt"
	"strconv"
)

func strConvert() {
	// 字符串 => 其他类型
	// => bool
	v, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(v) // true
	} else {
		fmt.Println(err)
	}

	// => int
	if v, err := strconv.Atoi("1023"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	// => int （parseint,指定进制数、目标int位数）
	if v, err := strconv.ParseInt("64", 16, 64); err == nil {
		fmt.Printf("%T %#v\n", v, v) // int64 100
	}

	// => float
	if v, err := strconv.ParseFloat("1.1", 64); err == nil {
		fmt.Printf("%T %#v\n", v, v) // float64 1.1
	}
}
