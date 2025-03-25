package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func funcCopyFile(src, dest string) {
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	} else {
		defer srcfile.Close()
		destfile, err := os.Create(dest)
		if err != nil {
			fmt.Println(err)
		} else {
			defer destfile.Close()

			reader := bufio.NewReader(srcfile)
			writer := bufio.NewWriter(destfile)

			bytes := make([]byte, 1024*1024*10)

			for {
				n, err := reader.Read(bytes)
				if err != nil {
					if err != io.EOF {
						fmt.Println(err)
					}
					break
				}
				fmt.Println(n)
				writer.Write(bytes[:n])
				writer.Flush()
			}

		}

	}
}

// 拷贝文件夹
func copyDir(src, dest string) {
	// 获取文件切片（路径名称）
	files, err := os.ReadDir(src)
	if err == nil {
		os.Mkdir(dest, os.ModePerm)
		for _, file := range files {
			// 递归调用，非文件就继续copyDir。若是文件就funCopyFile
			if file.IsDir() {
				// 上一层路径和本层路径进行拼接
				copyDir(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
			} else {
				// 末端文件直接复制
				funcCopyFile(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
			}
		}
	}
}

func copyfile() {
	src := flag.String("s", "", "src file")
	dest := flag.String("d", "", "dest file")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println(`
Usage: copyfile -s srcfile -d destfile
Options: 
		`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help || *src == "" || *dest == "" {
		flag.Usage()
	} else {
		// 判断src是否存在，若不存在就退出
		// src 目录 copyDir
		// src 文件 funcCopyFile
		// dest 若存在，就退出

		if _, err := os.Stat(*dest); err == nil {
			fmt.Println("目的文件已存在")
			return
		} else {
			if !os.IsNotExist(err) {
				// 非“不存在”引起
				fmt.Println("目的文件获取错误", err)
			}
		}

		if info, err := os.Stat(*src); err != nil {
			if os.IsNotExist(err) {
				fmt.Println("源文件不存在")
			} else {
				fmt.Println("源文件获取错误: ", err)
			}
		} else {
			if info.IsDir() {
				copyDir(*src, *dest)
			} else {
				funcCopyFile(*src, *dest)
			}
		}
	}

	// fmt.Printf("%T,%T,%v,%v\n", src, dest, *src, *dest) // *string,*string,,

}

// $ ./bin/chapter06 -s user.txt -d user2.txt
// $  md5 user.txt
// $  md5 user2.txt
