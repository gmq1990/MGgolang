package users

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
)

var users map[int]map[string]string

func init() {
	users = make(map[int]map[string]string)
}

const (
	MaxAuth = 3
	// 提前将密码转换成md5密文
	Password = "1fdb7184e697ab9355a3f1438ddc6ef9"
)

// 从命令行输入密码，并验证
// 用返回值告知验证结果
func Auth() bool {
	for i := 0; i < MaxAuth; i++ {
		fmt.Print("请输入密码：")
		// fmt.Scan(&input)
		bytes, _ := gopass.GetPasswd()
		// md5是转换成16个数字，要将数字转换成16进制
		if Password == fmt.Sprintf("%x", md5.Sum(bytes)) {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return false
}

func printUser(pk int, user map[string]string) {
	fmt.Println("ID:", pk)
	fmt.Println("名字:", user["name"])
	fmt.Println("出生日期:", user["birthday"])
	fmt.Println("联系方式:", user["tel"])
	fmt.Println("联系地址:", user["addr"])
	fmt.Println("备注:", user["desc"])
}

func InputStr(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

func getId() int {
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
	user["name"] = InputStr("请输入姓名：")
	user["birthday"] = InputStr("请输入出生日期（2000-01-01)：")
	user["tel"] = InputStr("请输入联系方式：")
	user["addr"] = InputStr("请输入家庭地址：")
	user["desc"] = InputStr("请输入备注：")

	return user
}

// 添加用户
func AddUser() {
	var id = getId()
	user := inputUser()
	users[id] = user
	fmt.Println("[+]添加成功")
}

// 修改用户
func Modify() {
	if id, err := strconv.Atoi(InputStr("请输入修改用户ID：")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息：")
			printUser(id, user)
			input := InputStr("确定修改(Y/N)? ")
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

// 删除用户
func DeleteUser() {
	if id, err := strconv.Atoi(InputStr("请输入删除用户ID：")); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息：")
			printUser(id, user)
			input := InputStr("确定删除(Y/N)? ")
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
func Query() {
	keyword := InputStr("请输入关键字：")
	title := fmt.Sprintf("%5s|%10s|%12s|%20s|%20s|%20s\n", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))

	for k, v := range users {
		if strings.Contains(v["name"], keyword) ||
			strings.Contains(v["tel"], keyword) ||
			strings.Contains(v["addr"], keyword) {
			fmt.Printf("%5d|%10s|%12s|%20s|%20s|%20s", k, v["name"], v["birthday"], v["tel"], v["addr"], v["desc"])
			fmt.Println()
		}
	}
}
