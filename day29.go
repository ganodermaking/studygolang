package main

import (
	"fmt"
	"time"
)

// 1.下面这段代码能否正常结束？
func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
}

// 2.下面这段代码输出什么？为什么？
func main() {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}
