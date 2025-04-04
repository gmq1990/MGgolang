package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// boss直聘。不要重复请求。以防IP被封
	url := "https://www.zhipin.com/web/geek/job?query=go&city=100010000"

	document, err := goquery.NewDocument(url)
	fmt.Println(document, err)
}
