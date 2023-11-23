// package main

// import (
// 	"fmt"
// 	"math"
// 	"unsafe"
// )

// func main() {

// 	var i8 int8
// 	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), -math.MaxInt8, math.MaxInt8)

// 	a := 10

// 	// 二进制
// 	fmt.Printf("%b", a)

// 	// 十进制
// 	fmt.Printf("%d \n", a)

// 	// 八进制
// 	fmt.Printf("%o \n", a)

// 	// 十六进制
// 	fmt.Printf("%x \n", a)

// 	// 浮点类型，假如用 math 包下的 PI 来输出
// 	// %.2f 代表输出两位小数
// 	pi := math.Pi
// 	fmt.Printf("%.2f \n", pi)
// 	fmt.Printf("%f \n", pi)
// }
