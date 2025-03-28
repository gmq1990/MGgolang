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
	SmsAPI string
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
	//  sender.SmsAPI ❌
	sender.send("tom", "早上好")

	// 指针？能调用值接收者的方法
	sender = &EmailSender{"xiaobai@example.com"}
	// sender.SmptAddr ❌
	// 值接收者能识别指针，从而指针能使用方法
	sender.sendAll([]string{"tom", "jerry"}, "早上好")
	do(sender)

	// WeChatSender只能声明成指针变量
	sender = &WeChatSender{"I Am ID"}
	// sender.ID ❌
	// interface 赋值给 interface
	// var 左边 = 右边
	// 右边 继承 左边
	// 左边的范围小，右边的范围大
	var ssender singleSender = sender

	ssender.send("xiaobai", "你好")

	// interface的类型断言：将interface转为特定类型
	sender01, ok := ssender.(Sender)
	fmt.Printf("%T %v\n", sender01, ok) // *main.WeChatSender true
	sender01.sendAll([]string{"xiaobai", "xiaohei"}, "你好")

	wesender01, ok := ssender.(*WeChatSender)
	fmt.Printf("%T %v\n", wesender01, ok) // *main.WeChatSender true

	esender01, ok := ssender.(EmailSender)
	// 此时sender已不再是emailsender
	fmt.Printf("%T %v\n", esender01, ok) // main.EmailSender false
	esender02, ok := sender.(EmailSender)
	fmt.Printf("%T %v\n", esender02, ok) // main.EmailSender false

	// 断言：类型查询
	switch ssender.(type) {
	case EmailSender:
		fmt.Println("emailsender")
	case *SMSSender:
		fmt.Println("*smssender")
	case *WeChatSender:
		fmt.Println("*wechatsender")
	}

}
