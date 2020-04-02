package main

import (
	"fmt"
	"time"
)

// 下面这段代码是否正常结束
func f290() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}

// 下面这段代码输出什么？为什么？
func f291() {
	var m = [...]int{1, 2, 3}
	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 3)
}
