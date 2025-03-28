package main

import (
	"fmt"
)

type Sender2 interface {
	Send(msg string) error
	Test()
}

type Receiver interface {
	Receive() (string, error)
	Test()
}

type Client interface {
	// 匿名嵌入其他接口
	Sender2
	Receiver
	Open() error
	Close() error
}

type MSNClient struct {
}

func (c MSNClient) Send(msg string) error {
	fmt.Println("send:", msg)
	return nil
}

func (c MSNClient) Receive() (string, error) {
	fmt.Println("receive")
	return "", nil
}

func (c MSNClient) Open() error {
	fmt.Println("Open")
	return nil
}

func (c MSNClient) Close() error {
	fmt.Println("Close")
	return nil
}

func (c MSNClient) Test() {

}

func combine() {
	msn := MSNClient{}

	var s Sender2 = msn
	var r Receiver = msn
	var c Client = msn

	s.Send("abc")
	r.Receive()
	c.Open()
	defer c.Close()
	c.Send("client")
	c.Receive()

	// 匿名接口（直接声明一个interface于函数内）
	var closer interface {
		Close() error
	}
	closer = msn
	closer.Close()
}
