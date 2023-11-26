// package main

// import "fmt"

// // 传统意义上的 for 循环
// func f1() {
// 	for i := 0; i <= 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 	}
// }

// // i 可以先定义，后 for 循环第一个赋值语句可以省略
// func f2() {
// 	i := 0
// 	for ; i <= 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 	}
// }

// // 第三种可以继续拆：先定义 for 循环第一条语句， for 可以只写第二条语句， 在最后执行 ++ 或 -- 这种自增自减操作
// func f3() {
// 	i := 0
// 	for i <= 10 {
// 		fmt.Printf("i: %v\n", i)
// 		i++
// 	}
// }

// func f4() {
// 	// 第四种 for ，意义上相当于 while 循环，永真循环
// 	for {
// 		fmt.Println("true")
// 		break
// 	}
// }

// func main() {
// 	f1()
// 	f2()
// 	f3()
// }
