package main

import (
	"fmt"
	"sort"
)

type User11 struct {
	ID    int
	Name  string
	Score float64
}

// 使用sort库排序
func sSort() {
	list := [][2]int{{1, 3}, {5, 9}, {4, 5}, {6, 2}, {5, 8}}
	// 排序: 使用数组第二个元素(索引为1)大小排序
	sort.Slice(list, func(i, j int) bool {
		return list[i][1] < list[j][1]
	})
	fmt.Println(list)

	users := []User11{{1, "tom", 1}, {2, "jerry", 5}, {3, "mangge", 3}, {4, "xiaohei", 2}}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Score < users[j].Score
	})
	fmt.Println(users)
}
