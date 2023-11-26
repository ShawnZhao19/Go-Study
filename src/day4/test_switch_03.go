// package main

// import "fmt"

// // go 语言中没有 break
// // go 的 switch 可以匹配多个条件，只要有其中一个满足即可

// func getLevel(score int) {

// 	switch score {
// 	case 90, 100:
// 		fmt.Println("优秀成绩")
// 	case 60:
// 		fmt.Println("一般般")
// 	default:
// 		fmt.Println("胡闹")
// 	}
// }

// // switch 后边可以什么也不跟，默认为 true 自动进入 switch 中， 可以在 case 中进行条件语句的判断
// func getLevel2(score int) {

// 	switch {
// 	case score >= 85 && score <= 100:
// 		fmt.Println("优秀成绩")
// 	case score >= 60 && score < 85:
// 		fmt.Println("一般般")
// 	default:
// 		fmt.Println("胡闹")
// 	}
// }
// func main() {
// 	var s int = 0
// 	getLevel(s)
// 	getLevel2(s)
// }
