package main

import (
	"encoding/csv"
	"os"
)

func csvwriter() {
	// 将内容写入到csv文件
	// 每个元素为字符串，“,”分割，没有其他格式
	file, err := os.Create("user.csv")
	if err == nil {
		defer file.Close()

		writer := csv.NewWriter(file)

		writer.Write([]string{"编号", "名字", "性别"})
		writer.Write([]string{"1", "张三", "男"})
		writer.Write([]string{"2", "李四", "男"})
		writer.Write([]string{"3", "王五", "男"})
		writer.Flush()
	}
}
