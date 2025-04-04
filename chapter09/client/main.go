package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	urllib "net/url"
	"os"
)

func main() {
	url := "http://localhost:9999/request/"
	response, err := http.Get(url)
	if err == nil {
		fmt.Println(response.Proto)
		fmt.Println(response.Status)
		fmt.Println(response.Header)
		io.Copy(os.Stdout, response.Body)
	}

	json := bytes.NewReader([]byte(`{"name": "tom", "password": "123456"}`))
	response, err = http.Post(url, "application/json", json)
	if err == nil {
		fmt.Println(response.Proto)
		fmt.Println(response.Status)
		fmt.Println(response.Header)
		io.Copy(os.Stdout, response.Body)
	}

	// 发送表单数据
	params := make(urllib.Values)
	params.Add("name", "tom")
	params.Add("password", "123456")

	response, err = http.PostForm(url, params)
	if err == nil {
		fmt.Println(response.Proto)
		fmt.Println(response.Status)
		fmt.Println(response.Header)
		io.Copy(os.Stdout, response.Body)
	}
}
