// package main

// import "fmt"

// func f1() {

// 	// 声明一个 map1 变量，它的类型是 string:string
// 	var map1 map[string]string
// 	// 这里的 map 只是声明，但并没有初始化，也没有开辟空间
// 	fmt.Printf("map1: %v\n", map1)
// 	map1 = make(map[string]string, 4)
// 	map1["name"] = "zs"
// 	map1["age"] = "18"

// 	map2 := map[string]string{
// 		"name": "zss",
// 		"age":  "17",
// 	}
// 	fmt.Printf("map1: %v\n", map1)
// 	fmt.Printf("map2: %v\n", map2)
// }

// func f2() {

// 	// 如何根据键来取对应的值
// 	m1 := make(map[int]int)
// 	m1[1] = 1
// 	m1[2] = 2
// 	m1[3] = 3
// 	index := 2
// 	value := m1[index]
// 	fmt.Printf("value: %v\n", value)
// }

// func f3() {
// 	m1 := make(map[int]int)
// 	m1[1] = 1
// 	m1[2] = 2
// 	m1[3] = 3

// 	// 这里的键值对如果根据 键取值 的话，有一个 ok 这个固定的变量，是在判断是否能取到值，它算是个 bool 类型
// 	v1, ok := m1[0]
// 	fmt.Printf("v1: %v\n", v1)
// 	fmt.Printf("ok: %v\n", ok)
// 	// fmt.Println("-----------------------")
// 	v2, ok := m1[1]
// 	fmt.Printf("v2: %v\n", v2)
// 	fmt.Printf("ok: %v\n", ok)
// }

// func main() {
// 	f1()
// 	f2()
// 	f3()
// }
