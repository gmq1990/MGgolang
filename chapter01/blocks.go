package main

import "fmt"

func blocks() {
	// 作用域：定义标识符可以使用的范围
	// 在GO中，用{}来定义作用域的范围
	// 使用原则：子语句块可以使用父语句块的变量，父语句块不能使用子语句块的变量
	outer := 1
	{
		inner := 2
		fmt.Println(inner)
		fmt.Println(outer)
		outer := 21
		{
			inner2 := 3
			fmt.Println(outer, inner, inner2)
		}
	}
	fmt.Println(outer)
	// fmt.Println(inner) ❌❌❌
}
