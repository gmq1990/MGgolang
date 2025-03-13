package main

import "fmt"

func integer() {
	// 整型
	// 标识符：int/int*/uint/uint*/uintptr/byte/rune
	// 字面量：十进制，八进制（0开头），十六进制（0X开头）
	var age int = 31
	fmt.Printf("%T %d\n", age, age)

	// 八进制
	fmt.Println(0777)
	// 十六进制
	fmt.Println(0x10)

	//操作
	// 算术运算（+,-,*,/,%,++,--)
	fmt.Println(1 + 2)
	fmt.Println(3 - 10)
	fmt.Println(3 * 9)
	fmt.Println(9 / 2)
	fmt.Println(9 % 2)

	age++
	fmt.Println(age)

	age--
	fmt.Println(age)

	// 关系运算(<,>,==,!=,>=,<=)
	fmt.Println(2 == 3)
	fmt.Println(2 != 3)
	fmt.Println(2 >= 3)
	fmt.Println(2 <= 3)

	// 位运算 二进制的运算 10 => 2
	// & | ^ << >> &^
	// 十进制 => 二进制 7/2 =>1 3/2 =>1 1/2 =>1 0  -> 0111
	// 2 => 十进制  0010
	// 7 & 2 => 0111 & 0010 => 0010 => 2
	// 7 | 2 => 0111 | 0010 => 0111 => 7
	// 7 ^ 2 => 0111 ^ 0010 => 0101 => 5
	// 2 << 1 => 0010 << 1 => 0100 => 4
	// 2 >> 1 => 0010 >> 1 => 0001 => 1
	// 7 &^ 2 => 0111 &^ 0010 => 0101 => 5 (按位清空，都为1的位，置为0)
	fmt.Println(7 & 2)
	fmt.Println(7 | 2)
	fmt.Println(7 ^ 2)
	fmt.Println(2 << 1)
	fmt.Println(2 >> 1)
	fmt.Println(7 &^ 2)

	// 附值运算(=, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=, ß&^=)
	// a+=b  ==> a=a+b
	age = 1
	age += 3 // age = age + 3
	fmt.Println(age)

	// int/uint/byte/rune/int*
	var intA int = 10
	var uintB uint = 3
	fmt.Println(intA + int(uintB))
	fmt.Println(uint(intA) + uintB)

	// 大 -> 小转换，可能出现溢出
	var intC int = 0xFFFF
	fmt.Println(intC, uint8(intC), int8(intC))

	// fmt.Printf()
	// int/uint/int*/uint*
	// byte, rune
	var a byte = 'A' // uint8
	var w rune = '中' // uint32
	fmt.Println(a, w)

	fmt.Printf("%T %c \n", a, a)
	fmt.Printf("%T %q %U \n", w, w, w)
	// 对齐（5位对齐）
	fmt.Printf("%5d\n", age)
	// 补0对齐
	fmt.Printf("%05d\n", age)

}
