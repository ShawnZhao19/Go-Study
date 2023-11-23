package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	// go 语言中的字符串是一个 任意字节的常量序列 [] byte
	var s1 string = "hello go"
	var s2 = "hello go"
	s3 := "hello go"
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)

	// 多行字符串的表示的话，用反引号 `
	s4 := `
		line 1
		line 2
		line 3
	`
	fmt.Printf("s4: %v\n", s4)

	// 字符串之间的连接方式1，用 + 号
	st := "go"
	st2 := "lang"
	st3 := st + st2
	fmt.Printf("st3: %v\n", st3)

	// 字符串之间的连接方式2，用 format.Sprintf() 方法来格式化
	st4 := fmt.Sprintf("%s%s", st, st2)
	fmt.Printf("st4: %v\n", st4)

	// 字符串之间的连接方式3，用 Strings.join() 方法来连接
	st5 := strings.Join([]string{st, st2}, "")
	fmt.Printf("st5: %v\n", st5)

	// 字符串之间的连接方式4，用 buffer 来连接，效率高
	var buffer bytes.Buffer
	buffer.WriteString(st)
	buffer.WriteString(st2)
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	// 转义字符 \' \r \n \t

	// 字符串的切片 []
	ss := "hello go"
	aa := 1
	bb := 6
	fmt.Printf("ss[aa]: %c\n", ss[aa])
	fmt.Printf("ss[aa:bb]: %c\n", ss[aa:bb])
	fmt.Printf("ss[aa:]: %c\n", ss[aa:])
	fmt.Printf("ss[:bb]: %c\n", ss[:bb])

	// strings 包里的一些方法 (经常用到)

	// 1、含有什么...
	fmt.Printf("strings.Contains(ss, \"lo\"): %v\n", strings.Contains(ss, "lo"))
}
