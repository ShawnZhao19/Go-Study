package main

import "fmt"

// 定义空数组
func f1() {
	var a1 [2]int
	var a2 [3]string
	fmt.Printf("a1: %T\n", a1)
	fmt.Printf("a2: %T\n", a2)
}

// 定义数组，并且初始化数组。这些数组的长度是固定的
func f2() {
	var a3 = [3]int{1, 2, 33}
	var a4 = [3]string{"a", "b", "c"}
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
}

// 定义数组，数组的长度可以使用 ... 来代替，自动推断数组的大小
func f3() {
	a5 := [...]int{5, 6, 77, 888}
	a6 := [...]string{"k1", "k2"}
	fmt.Printf("a5: %v\n", a5)
	fmt.Printf("a6: %v\n", a6)
}

// 定义数组，数组的值可以用 索引+值 来初始化
func f4() {
	a7 := [...]int{1: 2, 2: 3, 3: 4}
	a8 := [...]bool{2: false, 3: true}
	fmt.Printf("a7: %v\n", a7)
	fmt.Printf("a8: %v\n", a8)
}

func main() {
	f1()
	f2()
	f3()
	f4()
}
