package main

import "fmt"

// 1.下面这段代码输出什么？为什么？
// A. 1
// B. compilation error
func (i int) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i int = 1
	i.PrintInt()
}

// 2.下面这段代码输出什么？为什么？
// A. speak
// B. compilation error
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}
