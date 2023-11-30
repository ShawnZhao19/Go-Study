package main

import "fmt"

func getArr() {
	arr := [...]string{"a", "b", "c", "d"}
	for index, value := range arr {
		fmt.Printf("数组 arr 的键是 : %v, 键所对应的值是 : %v\n", index, value)
	}
}

// go 语言的 '切片' 概念？ 会慢慢铺开，这里自认为是对不定长数组的一个裁剪操作吧
func getArr2() {
	arr2 := []string{"a", "b", "c", "d"}
	for index, value := range arr2 {
		fmt.Printf("数组 arr 的键是 : %v, 键所对应的值是 : %v\n", index, value)
	}
}

// go 的 map 定义和存取值，自认为有点难理解，可能需要多花点时间去理解理解ing...
func getMap() {
	m := make(map[string]string, 0)
	m["name"] = "zs"
	m["age"] = "18"
	for k, v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
}

func main() {
	// getArr()
	// getArr2()
	getMap()
}
