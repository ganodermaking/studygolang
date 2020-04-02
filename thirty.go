package main

import (
	"fmt"
)

// 下面这段代码输出什么？
func f300(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()
	defer f()

	f = func() {
		r += 2
	}

	return n + 1
}

// 下面这段代码输出什么？
func f301() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r=", r)
	fmt.Println("a=", a)
}
