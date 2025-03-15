package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxAuth  = 3
	password = "123!@#"
)

// 从命令行输入密码，并验证
// 用返回值告知验证结果
func auth() bool {
	var input string
	for i := 0; i < maxAuth; i++ {
		fmt.Print("请输入密码：")
		fmt.Scan(&input)
		if password == input {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return false
}

func inputStr(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

func getId(users map[int]map[string]string) int {
	var id int
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

func inputUser() map[string]string {
	var user = map[string]string{}
	user["name"] = inputStr("请输入姓名：")
	user["birthday"] = inputStr("请输入出生日期（2000-01-01)：")
	user["tel"] = inputStr("请输入联系方式：")
	user["addr"] = inputStr("请输入家庭地址：")
	user["desc"] = inputStr("请输入备注：")

	return user
}

// 添加用户
func addUser(users map[int]map[string]string) {
	var id = getId(users)
	user := inputUser()
	users[id] = user
	fmt.Println("[+]添加成功")
}

func modify(users map[int]map[string]string) {
	if id, err := strconv.Atoi(inputStr("请输入修改用户ID：")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息：")
			fmt.Println(user)
			input := inputStr("确定修改(Y/N)? ")
			if input == "y" || input == "Y" {
				user := inputUser()
				users[id] = user
				fmt.Println("[+]修改成功")
			} else {
				fmt.Println("[-]用户ID不存在")
			}
		}
	} else {
		fmt.Println("[-]输入ID不正确")
	}
}

func deleteUser(users map[int]map[string]string) {
	if id, err := strconv.Atoi(inputStr("请输入删除用户ID：")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息：")
			fmt.Println(user)
			input := inputStr("确定删除(Y/N)? ")
			if input == "y" || input == "Y" {
				delete(users, id)
				fmt.Println("[+]删除成功")
			} else {
				fmt.Println("[-]用户ID不存在")
			}
		}
	} else {
		fmt.Println("输入ID不正确")
	}
}

// 查询用户
// keyword =="" 查找全部
// 非空，名称，电话，住址中包含keyword内容的，显示
func query(users map[int]map[string]string) {
	keyword := inputStr("请输入关键字：")
	title := fmt.Sprintf("%5s|%10s|%12s|%20s|%50s|%50s\n", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))

	for k, v := range users {
		if strings.Contains(v["name"], keyword) ||
			strings.Contains(v["tel"], keyword) ||
			strings.Contains(v["addr"], keyword) {
			fmt.Printf("%5d|%10s|%12s|%20s|%50s|%50s", k, v["name"], v["birthday"], v["tel"], v["addr"], v["desc"])
			fmt.Println()
		}
	}
}

func users() {

	if !auth() {
		fmt.Printf("密码%d次错误，程序退出", maxAuth)
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

	users := make(map[int]map[string]string)
	// id := 0
	callbacks := map[string]func(map[int]map[string]string){
		"1": addUser,
		"2": modify,
		"3": deleteUser,
		"4": query,
		"5": func(users map[int]map[string]string) {
			os.Exit(0)
		},
	}

	fmt.Println("欢迎使用用户管理系统")

	// END:
	for {
		fmt.Println(menu)
		op := inputStr("请输入指令: ")
		if callback, ok := callbacks[op]; ok {
			callback(users)
		} else {
			fmt.Println("[-]指令错误")
			time.Sleep(time.Second)
		}
		// switch op := inputStr("请输入指令: "); op {
		// case "1":
		// 	id++
		// 	addUser(users)
		// case "2":
		// 	modify(users)
		// case "3":
		// 	deleteUser(users)
		// case "4":
		// 	query(users)
		// case "5":
		// 	time.Sleep(500 * time.Millisecond)
		// 	fmt.Println("[-]即将退出")
		// 	time.Sleep(2 * time.Second)
		// 	break END
		// default:
		// 	fmt.Println("[-]指令错误")
		// 	time.Sleep(time.Second)
		// }

		// if op == "1" {
		// 	id++
		// 	adduser(id, users)
		// } else if op == "2" {

		// } else if op == "3" {

		// } else if op == "4" {
		// 	query(users)
		// } else if op == "5" {
		// 	time.Sleep(500 * time.Millisecond)
		// 	fmt.Println("即将退出")
		// 	time.Sleep(2 * time.Second)
		// 	break
		// } else {
		// 	fmt.Println("指令错误")
		// 	time.Sleep(time.Second)
		// }
	}
}
