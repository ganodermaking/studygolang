package main

import (
	"fmt"
)

type MyInt1 int
type MyInt2 = int

// 1.通过指针变量p访问其成员变量name，有哪几种方式？
// A p.name
// B (&p).name
// C (*p).name
// D p->name

// 2.下面这段代码是否能通过编译？如果通过，输出什么？
func main() {
	var i int = 0
	var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}
