package main

import "fmt"

// 1.下面这段代码输出什么？
// A.13.1
// B.13
// C.compilation error
func main() {
	a := 5
	b := 8.1
	fmt.Println(a + b)
}

// 2.下面这段代码输出什么？
// A.3
// B.4
// C.compilation error
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
}

// 3.下面这段代码输出什么？
// A. compilation error
// B. equal
// C. not equal
func main() {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
