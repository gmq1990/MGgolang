package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func mD5(reader *bufio.Reader) string {
	// 分切片读取
	bytes := make([]byte, 1024*1024*10)
	// hash md5计算器
	hasher := md5.New()

	for {
		n, err := reader.Read(bytes)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		// io.WriteString(hasher, string(bytes[:n]))
		hasher.Write(bytes[:n])
	}
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func md5stV2(txt string) (string, error) {
	// 用bufio封装strings
	reader := bufio.NewReader(strings.NewReader(txt))
	return mD5(reader), nil
}

func md5FileV2(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)
		return mD5(reader), nil
	}
}

func md5SumV2() {
	txt := flag.String("s", "", " a string")
	path := flag.String("f", "", " a file")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println(`
Usage: output a MD5 sum of a string or a file
Options:
		`)
		flag.PrintDefaults()
	}

	flag.Parse()

	var md5res string
	var err error

	if *help || (*path == "" && *txt == "") {
		flag.Usage()
	} else {
		if *txt != "" {
			md5res, err = md5stV2(*txt)
		} else if *path != "" {
			md5res, err = md5FileV2(*path)
		}
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(md5res)
	}

}
