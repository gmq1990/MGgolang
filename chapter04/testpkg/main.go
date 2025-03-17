package main

import (
	"fmt"
	// 导入的几种方法：
	// 相对导入
	// 绝对导入，常用
	// 空白导入（下划线），常用
	// 点导入
	// 别名导入（另外命名，解决包名冲突），常用
	"go-study/MGgolang/chapter04/gopkg"
	// 进行包初始化
	_ "go-study/MGgolang/chapter04/gopkg"

	"github.com/howeyc/gopass"
)

// main是用来声明本包为可执行程序，一个包下只能有一个main函数

func main() {
	// 调用gopkg包的VERSION常量
	fmt.Println(gopkg.VERSION)
	// fmt.Println(gopkg.name) // 无法调用name，包外不可见。首字母小写
	gopkg.PrintName() // 可调用PrintName，包外可见。首字母大写

	fmt.Print("请输入密码：")
	if bytes, err := gopass.GetPasswd(); err == nil {
		fmt.Println(string(bytes))
	}
}
