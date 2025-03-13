package main

import "fmt"

func pointer() {
	A := 2
	B := A
	B = 3
	fmt.Println(A, B)

	// 指针
	var cc *int = &A
	C := &A
	fmt.Printf("%T %T %p\n", C, cc, C)

	fmt.Println(*cc)
	*cc = 4
	fmt.Println(A)
}
