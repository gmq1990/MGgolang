package main

import (
	"errors"
	"fmt"
)

// 返回值，如何定义错误类型 error
// 如何创建错误类型的值 errors.New(), fmt.Errorf()
// 无错误 nil

func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("divided by zero")
	}

	return a / b, nil
}

func err() {
	fmt.Println(division(1, 3))

	if v, err := division(1, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	e := fmt.Errorf("Error: %s", "divided by zero")
	// e的类型、e的值
	fmt.Printf("%T, %v\n", e, e) // *errors.errorString, Error: divided by zero
}
