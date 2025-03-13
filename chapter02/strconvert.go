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

	// => float (parsefloat,指定目标int位数)
	if v, err := strconv.ParseFloat("1.1", 64); err == nil {
		fmt.Printf("%T %#v\n", v, v) // float64 1.1
	}

	//  其他类型 => string
	sf := fmt.Sprintf("%d", 12)
	fmt.Println(sf)

	sf = fmt.Sprintf("%f", 12.01)
	fmt.Printf("%q\n", sf)

	fmt.Printf("%q\n", strconv.FormatBool(false))
	fmt.Printf("%q\n", strconv.Itoa(12))
	// 指定进制数
	fmt.Printf("%q\n", strconv.FormatInt(12, 16))
	// 指定科学计数、原有精度、float64位
	fmt.Printf("%q\n", strconv.FormatFloat(10.1, 'E', -1, 64))
}
