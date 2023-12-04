// package main

// import "fmt"

// func f1() {
// 	arr1 := [...]int{77, 555, 34, 124}
// 	fmt.Printf("arr1: %v\n", arr1)
// 	arr1[3] = 11
// 	arr1[0] = 11
// 	fmt.Printf("arr1: %v\n", arr1)
// 	// 这里如果我们设置超于数组长度的下标的值的话，编译期就直接报错了
// 	// arr[4] = 1
// }

// // 获取数组的长度，系统内置方法， len()
// func f2() {
// 	arr2 := [...]string{"a", "bbc", "djkln"}
// 	fmt.Printf("len(arr2): %v\n", len(arr2))
// }

// // 如何遍历数组？
// // 1、根据数组长度遍历
// func f3() {
// 	var arr = [4]string{"daskj, kaslv, j2l, jkxlhc"}
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] == "j2l" {
// 			fmt.Println(i)
// 		}
// 	}
// }

// // 2、通过 for range 来遍历
// func f4() {
// 	var arr = [4]string{"daskj, kaslv, j2l, jkxlhc"}
// 	for _, v := range arr {
// 		fmt.Printf("v: %v\n", v)
// 	}
// }

// func main() {
// 	// f1()
// 	// f2()
// 	// f3()
// 	f4()
// }
