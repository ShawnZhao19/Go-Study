// package main

// import "fmt"

// func f1() {
// 	// 第一步，定义一个切片
// 	// 1、最直观的定义
// 	var sl1 []string
// 	sl1 = append(sl1, "bbcc", "dasdas")
// 	// 2、使用 make 方法来定义切片的容量及类型和初始值，这里的 2 代表容量2，并赋值了！初始值为0
// 	var sl2 = make([]int, 2)
// 	sl2 = append(sl2, 2, 3)
// 	// 3、 使用短语句来声明
// 	sl3 := []int{1, 2, 3}
// 	sl3 = append(sl3, 124124, 1000)
// 	// 4、 短语句 + make 方法来定义
// 	sl4 := make([]string, 4)
// 	sl4 = append(sl4, "bbcc", "dasdassda")

// 	sl5 := sl4[:]
// 	fmt.Printf("sl5: %v\n", sl5)

// 	sl6 := sl2[:1]
// 	fmt.Println(sl2)
// 	fmt.Println(sl3)
// 	fmt.Printf("sl6: %v\n", sl6)
// }

// func main() {
// 	f1()
// }
