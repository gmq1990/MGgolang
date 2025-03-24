package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
)

func md5str(txt string) {
	fmt.Printf("%x\n", md5.Sum([]byte(txt)))
}

func md5File(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)
		// 分切片读取
		bytes := make([]byte, 1024*1024*100)
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
			io.WriteString(hasher, string(bytes[:n]))
			// hasher.Write(bytes[:n])
		}
		fmt.Printf("%x\n", hasher.Sum(nil))
	}
}

func md5Sum() {
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

	if *help || (*path == "" && *txt == "") {
		flag.Usage()
	} else {
		if *txt != "" {
			md5str(*txt)
		} else if *path != "" {
			md5File(*path)
		}
	}
}
