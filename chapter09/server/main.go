package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func test02(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time is: %d", time.Now().Unix())
}

type Test03 struct {
}

func (t Test03) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time is: %s", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {

	// 定义处理器函数
	http.HandleFunc("/test01/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	http.HandleFunc("/test02/", test02)

	http.Handle("/test03/", Test03{})
	http.Handle("/test04/", &Test03{})

	http.HandleFunc("/request/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.UserAgent())
		fmt.Println(r.Referer())
		// 请求行
		fmt.Println(r.Method, r.URL, r.Proto)
		// 请求头
		fmt.Println(r.Header)

		// 请求体
		fmt.Println("体: ")
		// bytes := make([]byte, 1024)
		// n, err := r.Body.Read(bytes)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(string(bytes[:n]))
		io.Copy(os.Stdout, r.Body)
		w.Write([]byte("Request"))

	})

	// http.Handle => 类型转换
	http.Handle("/bin/", http.FileServer(http.Dir("."))) // 网络访问路径映射到本地路径

	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}
