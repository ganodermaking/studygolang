package main

import "fmt"

// 1.下面的两个切片声明中有什么区别？哪个更可取？
// A. var a []int
// B. a := []int{}

// 2. A、B、C、D 哪些选项有语法错误？
type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D
}

// 3.下面 A、B 两处应该填入什么代码，才能确保顺利打印出结果？
type S struct {
	m string
}

func f() *S {
	return __ // A
}

func main() {
	p := __          // B
	fmt.Println(p.m) // print "foo"
}
