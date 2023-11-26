package main

import "fmt"

func main() {
	var name string
	print("请输入你的名字:")
	fmt.Scan(&name)
	if name == "shawnzhao19" {
		fmt.Println("你真棒")
	} else {
		fmt.Println("你有点小失误")
	}
}
