// package main

// import "fmt"

// func main() {

// 	// 5、定义常量，关键字 const，定义必须赋初始值
// 	const PI float32 = 3.1415926

// 	const (
// 		a = 100
// 		b = 200
// 	)
// 	const w, y = 1, 2

// 	// iota 用法，默认值是从 0 开始，每调用一次，自动加 1，类似于 i++，只不过它是常量
// 	// 在 iota 的中间使用 _ 下划线，表示跳过 1，如果中间插入了别的常量，也是 iota 也是会自动加 1
// 	const (
// 		a1 = iota
// 		a2 = iota
// 		a3 = iota
// 	)
// 	fmt.Println(a1, a2, a3)
// }
