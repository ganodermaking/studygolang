package main

import "fmt"

// 1.下面代码下划线处填入哪个选项输出 yes nil？
// A. s1
// B. s2
// C. s1、s2 都可以
func main() {
	var s1 []int
	var s2 = []int{}
	if __ == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
}

// 2.下面这段代码输出什么？
// A. A
// B. 65
// C. compilation error
func main() {
	i := 65
	fmt.Println(string(i))
}

// 3.下面这段代码输出什么？
type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}
