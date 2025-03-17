package gopkg

import "fmt"

const VERSION = 1.0

const name = "local gopkg" // 包外不可见。首字母小写

func PrintName() { // 包外可见。首字母大写
	fmt.Println(name)
}

func init() {
	fmt.Println("init")
}
