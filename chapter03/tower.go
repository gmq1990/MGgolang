package main

import "fmt"

// 汉诺塔递归演示
func hanoi(a, b, c string, n int) {
	// a为起点，b为中转点，c为终点
	if n == 1 {
		fmt.Println(a, "->", c)
		return
	}

	// a: n-1个盘子 借助c 挪到b
	hanoi(a, c, b, n-1)
	fmt.Println(a, "->", c)

	// b: n-1个盘子 借助a 挪到c
	hanoi(b, a, c, n-1)
}

func tower_hanoi() {
	fmt.Println("1层：")
	hanoi("A", "B", "C", 1)
	fmt.Println("2层：")
	hanoi("A", "B", "C", 2)
	fmt.Println("3层：")
	hanoi("A", "B", "C", 3)
	fmt.Println("4层：")
	hanoi("A", "B", "C", 4)
	fmt.Println("5层：")
	hanoi("A", "B", "C", 5)

}
