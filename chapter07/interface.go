package main

import "fmt"

type singleSender interface {
	send(to, msg string) error
}

// 定义interface
type Sender interface {
	// 抽象出行为（func）
	// func的名字
	// func参数的数量、类型
	// func返回值的数量、类型
	send(to string, msg string) error
	sendAll(tos []string, meg string) error
}

type EmailSender struct {
	SmptAddr string
}

func (es EmailSender) send(to string, msg string) error {
	fmt.Println("发送邮件给：", to, ", 消息内容是：", msg)
	return nil
}
func (es EmailSender) sendAll(tos []string, msg string) error {
	for _, to := range tos {
		es.send(to, msg)
	}
	return nil
}

type SMSSender struct {
	SmsApiAddr string
}

// 把接收者定义成指针类型，看看有什么变化
func (ss *SMSSender) send(to string, msg string) error {
	fmt.Println("发送短消息给：", to, ", 短消息内容是：", msg)
	return nil
}
func (ss *SMSSender) sendAll(tos []string, msg string) error {
	for _, to := range tos {
		ss.send(to, msg)
	}
	return nil
}

type WeChatSender struct {
	ID string
}

// 两个方法，一个值接收者，一个指针接收者。有什么影响？
func (ws WeChatSender) send(to string, msg string) error {
	fmt.Println("发送短消息给：", to, ", 短消息内容是：", msg)
	return nil
}
func (ws *WeChatSender) sendAll(tos []string, msg string) error {
	for _, to := range tos {
		ws.send(to, msg)
	}
	return nil
}

func do(sender Sender) {
	sender.send("mom", "会晚点回家")
}

func iInterface() {
	// 定义变量
	// 将有相同行为的结构体赋值给同一个interface
	// sms、email是两个不同的结构体
	// var sender Sender = SMSSender{}
	var sender Sender = EmailSender{}

	// fmt.Printf("%T %#v\n", sender, sender) // main.EmailSender main.EmailSender{}
	// sender.send("tom", "早上好")
	// sender.sendAll([]string{"tom", "jerry"}, "早上好")

	do(sender)

	// 值？不能调用指针接收者的方法
	// sender = SMSSender{} ❌
	sender = &SMSSender{"Aggregate"}
	//  sender.SmsApiAddr ❌
	sender.send("tom", "早上好")

	// 指针？能调用值接收者的方法
	sender = &EmailSender{"xiaobai@example.com"}
	// sender.SmptAddr ❌
	// 值接收者能识别指针，从而使用方法
	sender.sendAll([]string{"tom", "jerry"}, "早上好")
	do(sender)

	// WeChatSender只能声明成指针变量
	sender = &WeChatSender{"imID"}
	// sender.ID ❌
	// interface 赋值给 interface
	// 后者继承前者
	// 前者的方法必须在后者中全部实现
	var ssender singleSender = sender

	ssender.send("xiaobai", "你好")

}
