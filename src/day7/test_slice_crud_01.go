// package main

// import "fmt"

// func f1() {
// 	// 这里使用 make 方法来初始化一个 string 类型的切片，并开辟了空间，两个空的字符串
// 	scl1 := make([]string, 2)
// 	fmt.Printf("scl1: %v\n", scl1)
// 	fmt.Printf("len(scl1): %v\n", len(scl1))
// }

// func f2() {

// 	var scl1 []string
// 	// 往切片中增加数据
// 	scl1 = append(scl1, "a", "b", "c", "d")
// 	fmt.Printf("scl1: %v\n", scl1)
// }

// func f3() {

// 	// 从切片中删除某个数据，自我感觉像是切割，不要某个元素，那就把那个元素前后切断然后再重新拼接起来
// 	// 公式：从切片中删除索引为 index 的元素时， xxx = append(xxx[:index], xxx[index+1:]...)
// 	var scl1 []int
// 	scl1 = append(scl1, 1, 2, 3, 4, 5)
// 	scl1 = append(scl1[:2], scl1[3:]...)
// 	fmt.Printf("scl1: %v\n", scl1)
// }

// func f4() {

// 	// update 这个改，就是拿出来某个元素给其重新赋值即可
// 	var scl1 []int
// 	scl1 = append(scl1, 1, 2, 3, 4, 5)
// 	scl1[1] = 100
// 	fmt.Printf("scl1: %v\n", scl1)
// }

// func f5() {

// 	// 查，就是遍历元素
// 	var scl1 []int
// 	scl1 = append(scl1, 1, 2, 3, 4, 5)
// 	var key = 3
// 	for index, v := range scl1 {
// 		if v == key {
// 			fmt.Printf("index: %v\n", index)
// 		}
// 	}
// }

// func f6() {

// 	// 切片的拷贝，如果是直接用 = 赋值的方式，相当于是把地址赋值了，如果修改了值，那么另外的也会一起跟着改变
// 	var scl1 = []string{}
// 	scl1 = append(scl1, "a", "c", "d")
// 	scl2 := scl1
// 	scl2[0] = "c"
// 	fmt.Printf("scl1: %v\n", scl1)
// 	fmt.Printf("scl2: %v\n", scl2)

// 	// 系统内置的 copy 方法可以避免这个问题，前提是要为切片分配空间，分配空间如果是0，则拷贝是失败的
// 	scl3 := make([]string, 0)
// 	copy(scl3, scl2)
// 	fmt.Printf("scl3: %v\n", scl3)
// }

// func main() {
// 	// f1()
// 	// f2()
// 	// f3()
// 	// f4()
// 	// f5()
// 	f6()
// }
