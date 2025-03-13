package main

import (
	"fmt"
	"strings"
)

func sStrings() {
	fmt.Println(strings.Compare("abc", "bcd"))                    // -1
	fmt.Println(strings.Contains("abc", "bc"))                    // true
	fmt.Println(strings.Count("abcabcd", "ab"))                   // 2
	fmt.Println(strings.Fields("abc def\neee\fxxx\vsddd"))        // 按空白字符分割(空格,\n,\r,\f,\t,\v)
	fmt.Println(strings.HasPrefix("abcd", "ab"))                  // true
	fmt.Println(strings.HasSuffix("abcdef", "def"))               // true
	fmt.Println(strings.Index("defabcdef", "abcf"))               // -1
	fmt.Println(strings.Index("defabcdef", "def"))                // 0
	fmt.Println(strings.LastIndex("defabcdef", "def"))            // 6
	fmt.Println(strings.Split("abdcdfe;efsd;dfs", ";"))           // 按指定符号分割
	fmt.Println(strings.Join([]string{"abd", "def", "ghi"}, ";")) // 按指定符号拼接
	fmt.Println(strings.Repeat("abc", 3))                         // 重复拼接
	fmt.Println(strings.Replace("abcabcabcab", "ab", "xx", -1))   // 替换ab为xx（-1全部，1一次）
	fmt.Println(strings.ToLower("ABC"))                           // abc
	fmt.Println(strings.ToUpper("abc"))                           // ABC
	fmt.Println(strings.Title("hi, kk"))                          // Hi, Kk
	fmt.Println(strings.Trim("xyzabcxyz", "xz"))                  // 删除首尾"x" "z"的字符
	fmt.Println(strings.TrimSpace("  acdb xxx \v \r \n"))         // 删除首位空白符
}
