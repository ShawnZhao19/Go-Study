package main

import "fmt"

func fun() {}

func main() {

	// 反向输出某个变量的类型
	name := "zs"
	age := 3.1
	vender := true

	// %T 表示获取变量类型
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", vender)

	// 指针类型 & 用这个符号
	a := 1
	b := &a
	// b 的值是 a 的指针类型，*int
	fmt.Printf("%T\n", b)

	// 数组类型 array
	// 定义： = 号右边首先是数组长度 [...] 三个点代表长度可以省略， 然后跟上数组类型 int ，最后是 {数组的值}
	arr := [...]int{1, 2, 3}
	fmt.Printf("%T\n", arr)

	// 切片类型 = 动态数组
	// 定义： = 号右边首先是 [] ，不需要写长度，因为可以动态添加， 然后跟上类型 int，最后是 {值}
	sl := []int{1, 2}
	fmt.Printf("%T\n", sl)

	// 函数类型
	fmt.Printf("%T\n", fun)
}
