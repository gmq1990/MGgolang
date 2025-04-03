package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr := "127.0.0.1:9999"
	url := "/500.html"
	client, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Fprintf(client, "GET %s HTTP/1.0/\r\n", url)

	defer client.Close()
	reader := bufio.NewReader(client)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Print(line)
	}
}
