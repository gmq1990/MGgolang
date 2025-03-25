package main

import (
	"fmt"
	"os"
	"time"
	upkg "usermanager/users"
)

func main() {

	if !upkg.Auth() {
		fmt.Printf("[-]密码%d次错误，程序退出", upkg.MaxAuth)
		fmt.Println()
		return
	}

	menu := `********************
1. 新建用户
2. 修改用户
3. 删除用户
4. 查询用户
5. 退出
********************`

	// id := 0
	callbacks := map[string]func(){
		"1": upkg.AddUser,
		"2": upkg.Modify,
		"3": upkg.DeleteUser,
		"4": upkg.Query,
		"5": func() {
			// time.Sleep(500 * time.Millisecond)
			// fmt.Println("即将退出")
			// time.Sleep(2 * time.Second)
			os.Exit(0)
		},
	}

	fmt.Println("欢迎使用用户管理系统")

	// END:
	for {
		fmt.Println(menu)
		op := upkg.InputStr("请输入指令: ")
		if callback, ok := callbacks[op]; ok {
			callback()
		} else {
			fmt.Println("[-]指令错误")
			time.Sleep(time.Second)
		}
	}
}
