package main

import (
	"fmt"
	"time"
)

func tTime() {
	now := time.Now()
	fmt.Printf("%T\n", now)
	fmt.Printf("%v\n", now)

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("15:04:05"))

	// 转成字符串
	fmt.Println(now.Unix())
	// 纳秒，转换成字符串
	fmt.Println(now.UnixNano())

	// 前后格式需要一致，否则报错。前为格式，后为时间的字符串
	t, err := time.Parse("2006/01/02 15:04:05", "2006/01/02 03:04:05")
	fmt.Println(err, t)

	// 第一个0为秒，第二个0为纳秒
	t = time.Unix(0, 0)
	fmt.Println(t)

	d := now.Sub(t)
	fmt.Printf("%T %v\n", d, d)

	// time.Second
	// time.Minute
	// time.Hour

	fmt.Println(time.Now())
	time.Sleep(time.Second * 4)
	fmt.Println(time.Now())

	t = now.Add(3 * time.Hour)
	fmt.Println(t)

	// 字符串 => 解析 => 时间区间
	d, err = time.ParseDuration("3h2m4s")
	fmt.Println(err, d)
	// => 小时
	fmt.Println(d.Hours())
	// => 分钟
	fmt.Println(d.Minutes())
	// => 秒
	fmt.Println(d.Seconds())
}
