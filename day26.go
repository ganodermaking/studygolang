package main

import "fmt"

// 1.下面这段代码输出什么？
const (
	a = iota
	b = iota
)
const (
	name = "name"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// 2.下面这段代码输出什么？为什么？
type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func main() {

	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
