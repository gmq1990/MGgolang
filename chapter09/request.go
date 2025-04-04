package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

func request() {
	// 目标url
	// 这里使用一个过期的证书的url
	url := "https://expired.badssl.com/"
	// request对象
	request, _ := http.NewRequest("GET", url, nil)
	// trasport对象
	transport := &http.Transport{
		// 跳过不安全的证书验证
		// false不跳过，true跳过。
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// client对象
	client := &http.Client{Transport: transport}

	response, err := client.Do(request)
	if err == nil {
		fmt.Println(response.Proto)
		fmt.Println(response.Status)
		fmt.Println(response.Header)
		io.Copy(os.Stdout, response.Body)
	}

}
