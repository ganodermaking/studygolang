package main

import (
	"fmt"
)

// 1.下面两段代码输出什么
func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s1 := make([]int, 0)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)
}

// 2.下面这段代码有什么缺陷
func funcMui(x, y int) (sum int, error) {
	return x + y, nil
}

// 3.new() 和 make() 的区别
