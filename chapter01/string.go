package main

import "fmt"

func sString() {
	// "" => 可解释字符串
	// `` => 原生字符串
	// 特殊字符：\r \n \t \f \b \v
	var name string = "k\tk"
	var desc string = `你\t好`
	fmt.Println(name)
	fmt.Println(desc)

	// 操作
	// 算数运算符：+ （连接字符串）
	fmt.Println("我叫" + "kk")
	// 关系运算（== != > >= < <=)
	fmt.Println("ab" == "bb")
	fmt.Println("ab" != "bb")
	fmt.Println("ab" < "bb")
	fmt.Println("ab" <= "bb")
	fmt.Println("ab" > "bb")
	fmt.Println("ab" >= "bb")
	fmt.Println("bb" >= "b")

	// 赋值
	s := "我叫"
	s += "kk"
	fmt.Println(s)

	// 字符串定义内容必须只能为ascii
	// 索引 0 ～ n-1 （n，字符串长度）
	desc = "abcdef"
	fmt.Printf("%T %c\n", desc[0], desc[0])
	// 切片[start:end] start ~ end-1
	fmt.Printf("%T\n", desc[0:2])
	// len()，ascii码的字符串
	fmt.Println(len(desc))
}
