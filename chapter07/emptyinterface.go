package main

import "fmt"

type EStruct struct {
}

// 空接口
type Empty interface {
}

func fargs(args ...interface{}) {
	for _, arg := range args {
		// fmt.Println(arg)
		switch arg.(type) {
		case int:
			fmt.Printf("int: %T %v\n", arg, arg)
		case string:
			fmt.Printf("string: %T %v\n", arg, arg)
		default:
			fmt.Printf("other: %T %v\n", arg, arg)
		}
	}
}

func emptyInterface() {
	es := EStruct{}
	// 空接口可接收任意类型
	var e Empty = 1

	fmt.Println(es) // {}
	fmt.Println(e)

	e = "test"
	fmt.Println(e)

	e = es
	fmt.Println(e)

	fargs(1, "test", true, es)

}
