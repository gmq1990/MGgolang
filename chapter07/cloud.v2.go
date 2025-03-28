package main

import "fmt"

type Host2 struct {
	ID   string
	Name string
}

type Cloud interface {
	GetList2() []Host2
	Start2(Id string) error
	Stop2(Id string) error
	Detail2(Id string) Host2
}

type Tencent2 struct {
	API    string
	key    string
	secret string
}

func NewTencent2(api, key, secret string) Tencent2 {
	return Tencent2{api, key, secret}
}

func (t Tencent2) GetList2() []Host2 {
	fmt.Println("Tencent.Getlist")
	return []Host2{}
}

func (t Tencent2) Start2(id string) error {
	fmt.Println("Tencent.Start")
	return nil
}

func (t Tencent2) Stop2(id string) error {
	fmt.Println("Tencent.Stop")
	return nil
}

func (t Tencent2) Detail2(id string) Host2 {
	fmt.Println("Tencent.Detail")
	return Host2{}
}

type Aliyun2 struct {
	API    string
	key    string
	secret string
}

func NewAliyun2(api, key, secret string) Aliyun2 {
	return Aliyun2{api, key, secret}
}

func (t Aliyun2) GetList2() []Host2 {
	fmt.Println("Aliyun.Getlist")
	return []Host2{}
}

func (t Aliyun2) Start2(id string) error {
	fmt.Println("Aliyun.Start")
	return nil
}

func (t Aliyun2) Stop2(id string) error {
	fmt.Println("Aliyun.Stop")
	return nil
}

func (t Aliyun2) Detail2(id string) Host2 {
	fmt.Println("Aliyun.Detail")
	return Host2{}
}

func NewCloud(cloudType string) Cloud {
	var cloud Cloud
	if cloudType == "tencent" {
		cloud = NewTencent2("tencent", "tencent", "tencent")
	} else if cloudType == "aliyun" {
		cloud = NewAliyun2("aliyun", "aliyun", "aliyun")
	} else {
		fmt.Println("error")
	}
	return cloud
}

func cloudV2() {
	var cloudType string // tencent/aliyun
	cloudType = "aliyun"

	var cloud = NewCloud(cloudType)
	if cloud == nil {
		return
	}

	var operate string
EOF:
	for {
		fmt.Scan(&operate)

		switch operate {
		case "1":
			cloud.GetList2()
			fmt.Println("同步主机列表")
			fmt.Println("插入数据库")
			fmt.Println("通知添加新主机")
		case "2":
			cloud.Start2("1")
			fmt.Println("启动主机")
		case "3":
			cloud.Stop2("2")
			fmt.Println("停止主机")
		case "4":
			cloud.Detail2("3")
			fmt.Println("获取主机详情")
		case "5":
			break EOF
		}
	}
}
