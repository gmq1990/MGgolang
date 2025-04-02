package main

import (
	"fmt"
	"net"
)

func nNet() {
	// joinHostPort,连接一个IP和一个端口
	fmt.Println(net.JoinHostPort("0.0.0.0", "8888"))
	// SplitHostPort, 拆分一个IP:Port为一个IP和一个端口
	fmt.Println(net.SplitHostPort("127.0.0.1:9999"))
	// 获得主机名
	fmt.Println(net.LookupAddr("23.218.109.123")) // [a23-218-109-123.deploy.static.akamaitechnologies.com.] <nil>
	// 获取主机IP地址
	fmt.Println(net.LookupHost("www.baidu.com"))
	// CIDR带掩码格式。ParseCIDR获取地址、掩码
	fmt.Println(net.ParseCIDR("192.168.1.1/24")) // 192.168.1.1 192.168.1.0/24 <nil>
	// 转换为IP。错误返回nil
	ip := net.ParseIP("240e:ff:e020:98c:0:ff:b061:c306")
	fmt.Println(ip)
	// 查找IP集合（与lookuphost类似）
	ips, _ := net.LookupIP("www.baidu.com")
	fmt.Println(ips)
	// 掩码格式获取ip和掩码
	ip, ipnet, _ := net.ParseCIDR("192.168.1.1/24")
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.1.40"))) // true
	fmt.Println(ipnet.Contains(net.ParseIP("192.168.2.24"))) // false
	fmt.Println(ipnet.Network())                             // ip+net
	// 获取本地主机的所有IP
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		fmt.Println(addr.Network(), addr.String())
	}
	// 获取网络接口
	inters, _ := net.Interfaces()
	for _, inter := range inters {
		fmt.Println(inter)
		// 单播地址列表
		fmt.Println(inter.Addrs())
		// 多播地址列表
		fmt.Println(inter.MulticastAddrs())
	}
}
