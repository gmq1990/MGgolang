package users

import (
	"bufio"
	"crypto/md5"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/howeyc/gopass"
)

const (
	MaxAuth      = 3
	passwordFile = "password"
	userFile     = "users.csv"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
	Desc     string
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d\n名字: %s\n出生日期: %s\n联系方式: %s\n联系地址: %s\n备注: %s\n", u.ID, u.Name, u.Birthday.Format("2006-01-02"), u.Tel, u.Addr, u.Desc)
}

// 从文件读取
func loadUser() map[int]User {
	users := map[int]User{}
	if file, err := os.Open(userFile); err == nil {
		defer file.Close()
		reader := csv.NewReader(file)
		for {
			line, err := reader.Read()
			if err != nil {
				if err != io.EOF {
					fmt.Println("[-]发生错误:", err)
				}
				break
			}

			id, _ := strconv.Atoi(line[0])
			birthday, _ := time.Parse("2006-01-02", line[2])
			users[id] = User{
				ID:       id,
				Name:     line[1],
				Birthday: birthday,
				Tel:      line[3],
				Addr:     line[4],
				Desc:     line[5],
			}
		}
	} else {
		if !os.IsNotExist(err) {
			fmt.Println("[-]发生错误:", err)
		}
	}
	return users
}

// 持久化到文件
func storeUser(users map[int]User) {
	// 重命名文件
	if _, err := os.Stat(userFile); err == nil {
		// os.Rename(userFile, fmt.Sprintf("%d", time.Now().Unix()))
		// os.Rename(userFile,time.Now().Format("2006-01-02 15:04:05"))
		os.Rename(userFile, strconv.FormatInt(time.Now().Unix(), 10)+".user.csv")
	}

	if names, err := filepath.Glob("*.user.csv"); err == nil {
		fmt.Println(names)
		// 按降序排序
		sort.Sort(sort.Reverse(sort.StringSlice(names)))
		fmt.Println(names)
		// 删除旧文件，保留3个最近的
		if len(names) > 3 {
			for _, name := range names[3:] {
				if err := os.Remove(name); err != nil {
					fmt.Printf("删除文件 %s 时出错: %v\n", name, err)
				}
			}
		}
	}

	if file, err := os.Create(userFile); err == nil {
		defer file.Close()
		writer := csv.NewWriter(file)
		for _, user := range users {
			writer.Write([]string{
				strconv.Itoa(user.ID),
				user.Name,
				user.Birthday.Format("2006-01-02"),
				user.Tel,
				user.Addr,
				user.Desc,
			})
		}
		writer.Flush()
	}
}

// 从命令行输入密码，并验证
// 用返回值告知验证结果
func Auth() bool {
	password, err := os.ReadFile(passwordFile)
	if err == nil && len(password) > 0 {
		// 验证密码
		for i := 0; i < MaxAuth; i++ {
			fmt.Print("请输入密码：")
			// fmt.Scan(&input)
			bytes, _ := gopass.GetPasswd()
			// md5是转换成16个数字，要将数字转换成16进制
			if string(password) == fmt.Sprintf("%x", md5.Sum(bytes)) {
				return true
			} else {
				fmt.Println("[-]密码错误")
			}
		}
		return false
	} else {
		if os.IsNotExist(err) || len(password) == 0 {
			// 初始化密码文件
			fmt.Print("请输入初始化密码:")
			bytes, _ := gopass.GetPasswd()
			// 将密码转换成md5密文，写入文件
			os.WriteFile(passwordFile, fmt.Appendf(nil, "%x", md5.Sum(bytes)), os.ModePerm)
			return true
		} else {
			// 其他原因
			fmt.Println("[-]发生错误: ", err)
			return false
		}
	}
}

func InputStr(prompt string) string {
	fmt.Print(prompt)
	// 带缓冲io，防止一行输入被多个命令行读取
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func getId() int {
	var id int
	users := loadUser()
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

func inputUser(id int) User {
	var user User = User{}

	user.ID = id
	user.Name = InputStr("请输入姓名：")
	// 字符串转换为时间格式
	birthday, _ := time.Parse("2006-01-02", InputStr("请输入出生日期（2000-01-01)："))
	user.Birthday = birthday
	user.Tel = InputStr("请输入联系方式：")
	user.Addr = InputStr("请输入家庭地址：")
	user.Desc = InputStr("请输入备注：")

	return user
}

func validUser(user User) error {
	if user.Name == "" {
		return errors.New("输入的用户名为空")
	}
	users := loadUser()
	for _, tuser := range users {
		if tuser.Name == user.Name && tuser.ID != user.ID {
			return errors.New("输入的名字已存在")
		}
	}
	return nil
}

// 添加用户
func AddUser() {
	var id = getId()
	user := inputUser(id)

	if err := validUser(user); err == nil {
		users := loadUser()
		users[id] = user
		storeUser(users)
		fmt.Println("[+]添加成功")
	} else {
		fmt.Println("[-]添加失败：", err)
	}

}

// 修改用户
func Modify() {
	if id, err := strconv.Atoi(InputStr("请输入修改用户ID：")); err == nil {
		users := loadUser()
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息：")
			fmt.Println(user)
			input := InputStr("确定修改(Y/N)? ")
			if input == "y" || input == "Y" {
				user := inputUser(id)
				if err := validUser(user); err == nil {
					users[id] = user
					storeUser(users)
					fmt.Println("[+]修改成功")
				} else {
					fmt.Println("[-]修改失败：", err)
				}
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
		users := loadUser()
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息：")
			fmt.Println(user)
			input := InputStr("确定删除(Y/N)? ")
			if input == "y" || input == "Y" {
				delete(users, id)
				storeUser(users)
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
	list := make([]User, 0)
	users := loadUser()
	fmt.Println("==============================")
	// title := fmt.Sprintf("%5s|%10s|%12s|%20s|%20s|%20s\n", "ID", "Name", "Birthday", "Tel", "Addr", "Desc")
	// fmt.Println(title)
	// fmt.Println(strings.Repeat("-", len(title)))
	for _, v := range users {
		if strings.Contains(v.Name, keyword) ||
			strings.Contains(v.Tel, keyword) ||
			strings.Contains(v.Addr, keyword) ||
			strings.Contains(v.Desc, keyword) {
			list = append(list, v)
		}
	}
	sortKey := InputStr("请输入排序字段(id/name/tel/addr/desc):")

	sort.Slice(list, func(i, j int) bool {
		switch sortKey {
		case "id":
			return list[i].ID > list[j].ID
		case "name":
			return list[i].Name > list[j].Name
		case "tel":
			return list[i].Tel > list[j].Tel
		case "addr":
			return list[i].Addr > list[j].Addr
		case "desc":
			return list[i].Desc > list[j].Desc
		default:
			return list[i].ID > list[j].ID
		}
	})
	for _, user := range list {
		fmt.Println(user)
	}
	fmt.Println("==============================")
}

// 修改密码
func ModifyPasswd() {
	password, err := os.ReadFile(passwordFile)
	if err == nil {
		// 验证密码
		fmt.Print("请输入旧密码:")
		// fmt.Scan(&input)
		bytes, _ := gopass.GetPasswd()
		// md5是转换成16个数字，要将数字转换成16进制
		if string(password) == fmt.Sprintf("%x", md5.Sum(bytes)) {
			for {
				fmt.Print("请输入新密码:")
				bytes, _ = gopass.GetPasswd()
				if len(bytes) > 0 {
					break
				} else {
					fmt.Println("不能是空密码！")
				}
			}
			os.WriteFile(passwordFile, fmt.Appendf(nil, "%x", md5.Sum(bytes)), os.ModePerm)
			fmt.Println("[+]密码修改成功")
		} else {
			fmt.Println("[-]密码错误")
		}
	}
}
