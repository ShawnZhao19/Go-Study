package main

import (
	"fmt"
)

func sum(a int, b int) (sum int) {
	sum = a + b
	return sum
}

// 这里即使 res 没用到，也不会报错
func add(a string, b string) (res string) {
	return a + b
}

func main() {

	// 1、在 go 语言中有三种函数：普通函数、匿名函数(没有名称的函数)、方法(定义在 struct 上的函数)
	// 2、在 go 语言中是不允许函数的重载的，也就是说不允许函数同名
	// 3、在 go 语言中的函数不能嵌套函数，但是可以嵌套匿名函数
	// 4、在 go 语言中函数是一个值，可以将函数赋值给变量，使得这个变量也是一个函数
	// 5、在 go 语言中一个函数的返回值可以当成参数传递给另外一个函数
	// 6、在 go 语言中函数调用时，如果有参数传递给函数，则会先拷贝参数的副本，再把副本传递给参数
	// 7、在 go 语言中函数参数可以没有名称 _

	i := sum(1, 2)
	fmt.Printf("i: %v\n", i)

	s := add("a", "b")
	fmt.Printf("s: %v\n", s)
}
