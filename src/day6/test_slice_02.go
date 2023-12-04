// package main

// import "fmt"

// func f1() {

// 	// 数组在定义时，长度是固定的，后续不可变
// 	var arr1 = [...]string{"a", "b", "c"}
// 	fmt.Printf("arr1: %v\n", arr1)

// 	// 切片在定义时，长度并不固定，后续是动态扩容的
// 	// ps: 这里只是声明了一个切片，并没有给其分配空间
// 	var slice1 []string
// 	fmt.Printf("slice1: %v\n", slice1)

// 	// 我们在定义切片时，可以使用 make 方法来初始化切片的容量，使用短语句创建一个切片
// 	// 这里是给切片开辟了一片空间，不过是空间为 0
// 	slice2 := make([]int, 0)
// 	slice2 = append(slice2, 1, 2, 3, 4)
// 	fmt.Printf("slice2: %v\n", slice2)

// 	// cap() 方法是看某个变量的容量是多少
// 	fmt.Printf("cap(slice2): %v\n", cap(slice2))

// }

// // 这里引入切片的概念，golang 的切片相对于数组而言，它的长度是动态的，因此可以实现自动扩容。自我感觉是若干数组的集合吧，很像是 List ？
// // var xxx []type
// func main() {
// 	f1()
// }
