// package main

// import "fmt"

// func f1() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 		if i == 5 {
// 			continue
// 		}
// 		if i == 7 {
// 			break
// 		}
// 	}
// }

// func f2() {
// 	i := 1
// 	for ; i < 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Printf("i: %v\n", i)
// 		} else {

// 			// continue 关键字跳出本次循环， 进入下一次循环
// 			continue

// 			// 实质上这条语句是执行不到的
// 			fmt.Println("hello go")
// 		}
// 	}
// }

// func f3() {
// 	for i := 0; i < 10; i++ {
// 	MYLABEL:
// 		for j := 0; j < 10; j++ {
// 			if i%2 == 0 && j%2 == 0 {
// 				continue MYLABEL
// 			}
// 			fmt.Printf("i:%v, j:%v\n", i, j)
// 		}
// 	}
// }

// func main() {
// 	// f1()
// 	// f2()
// 	f3()
// }
