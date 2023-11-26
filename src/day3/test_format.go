// package main

// import "fmt"

// // 结构体
// type WebSite struct {
// 	Name string
// 	Age  int8
// }

// func main() {

// 	ws := WebSite{
// 		Name: "zs",
// 		Age:  18,
// 	}

// 	// fmt 本意是格式化包，会将值格式化输出
// 	// %v 取的是某个变量的值
// 	// %#v 会将值包括哪个函数调用一起输出出来
// 	// ws: main.WebSite{Name:"zs", Age:18}
// 	fmt.Printf("ws: %v\n", ws)
// 	fmt.Printf("ws: %#v\n", ws)
// 	// %T 代表输出的 type 类型
// 	fmt.Printf("ws: %T\n", ws)

// 	// %p 代表输出的是指针地址
// 	a := 100
// 	p := &a
// 	fmt.Printf("p: %p\n", p)
// }
