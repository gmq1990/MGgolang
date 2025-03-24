package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
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
		funcCopyFile(*src, *dest)
	}

	// fmt.Printf("%T,%T,%v,%v\n", src, dest, *src, *dest) // *string,*string,,

}

// $ ./bin/chapter06 -s user.txt -d user2.txt
// $  md5 user.txt
// $  md5 user2.txt
