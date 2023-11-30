// package main

// import "fmt"

// func main() {

// 	a := 100
// 	b := 10

// 	if a > 0 {
// 		a -= 10
// 		b--
// 		fmt.Printf("a: %v\n", a)
// 		fmt.Printf("b: %v\n", b)
// 	}

// 	switch a {
// 	case 100:
// 		a -= 10
// 		fmt.Printf("a: %v\n", a)
// 		fmt.Printf("b: %v\n", b)
// 	case 90:
// 		a -= 20
// 		fmt.Printf("b: %v\n", b)
// 	}

// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 	}

// 	arr := [...]string{"a", "b", "c"}

// 	// _ 下划线的值指的是索引的值，默认为不用，不输出，也可以输出
// 	for s, v := range arr {
// 		fmt.Printf("s: %v\n", s)
// 		fmt.Printf("v: %v\n", v)
// 	}

// }
