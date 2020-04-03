package main

import "fmt"

// 1.定义一个包内全局字符串变量，下面语法正确的是（）
// A. var str string
// B. str := ""
// C. str = ""
// D. var str = ""

// 2.下面这段代码输出什么?
func hello(i int) {
	fmt.Println(i)
}

func main() {
	i := 5
	defer hello(i)
	i = i + 10
}

// 3.下面这段代码输出什么？
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
