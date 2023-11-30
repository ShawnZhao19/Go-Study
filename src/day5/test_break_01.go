// package main

// import "fmt"

// // Go 的 switch 默认相当于每个 case 最后都有 break
// // 匹配成功后不会自动向下执行其他 case，而是跳出整个 switch , 但是可以使用 fallthrough 强制执行后面的 case 代码
// func f2() {
// 	i := 1
// 	switch {
// 	case i <= 1:
// 		fmt.Println("1")
// 		break
// 		fallthrough // fallthrough 强制执行下一个 call 的代码，但如果 fallthrough 上有 break ，就会执行完当前 case 退出
// 	case i > 1:
// 		fmt.Println("2")
// 	}
// }

// func f1() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 		if i == 5 {
// 			break
// 		}
// 	}
// }

// func f3() {

// 	// 可以在方法中定义一个标签， break 的时候可以使其调到某一个标签的位置，并跳出某个循环
// MYLABEL:
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 		if i == 4 {
// 			break MYLABEL
// 		}

// 	}
// 	fmt.Println("ending....")
// }
// func main() {
// 	// f1()
// 	// f2()
// 	f3()
// }
