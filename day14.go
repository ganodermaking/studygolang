package main

import "fmt"

// 1.下面代码输出什么？
// A. hello
// B. xello
// C. compilation error
func main() {
	str := "hello"
	str[0] = 'x'
	fmt.Println(str)
}

// 2.下面代码输出什么？
// A. 1
// B. 2
// C. 3
func incr(p *int) int {
	*p++
	return *p
}

func main() {
	p := 1
	incr(&p)
	fmt.Println(p)
}

// 3.对 add() 函数调用正确的是（）
// A. add(1, 2)
// B. add(1, 3, 7)
// C. add([]int{1, 2})
// D. add([]int{1, 3, 7}…)
func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
