// package main

// import "fmt"

// func guess(gender int, age int) {
// 	if gender == 0 {
// 		if age >= 18 {
// 			fmt.Println("你的性别是女，而且你已经十八岁啦，成年了")
// 		} else if age <= 0 {
// 			fmt.Println("你的年龄怎么可能是 0 或是比 0 还小呢，请重新输入")
// 		} else {
// 			fmt.Println("你的性别是女，而且你还没有十八岁，未成年人哦")
// 		}
// 	} else if gender == 1 {
// 		if age >= 18 {
// 			fmt.Println("你的性别是男，而且你已经十八岁啦，成年了")
// 		} else if age <= 0 {
// 			fmt.Println("你的年龄怎么可能是 0 或是比 0 还小呢，请重新输入")
// 		} else {
// 			fmt.Println("你的性别是男，而且你还没有十八岁，未成年人哦")
// 		}
// 	} else {
// 		fmt.Println("你的性别输入有误， 0 代表的是女性， 1 代表的是男性， 请重新输入")
// 	}
// }
// func main() {

// 	// if 嵌套，也就是 if 里还有 if 了
// 	var gender int
// 	var age int
// 	fmt.Println("请输入您的性别和年龄: ")
// 	fmt.Scan(&gender, &age)
// 	guess(gender, age)
// }
