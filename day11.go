package main

import "fmt"

// 1.关于 cap() 函数的适用类型，下面说法正确的是()
// A. array
// B. slice
// C. map
// D. channel

// 2.下面这段代码输出什么？
// A. nil
// B. not nil
// C. compilation error
func main() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

// 3.下面这段代码输出什么？
// A. runtime panic
// B. 0
// C. compilation error
func main() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}
