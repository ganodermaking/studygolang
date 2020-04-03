package main

// f1()、f2()、f3() 函数分别返回什么？
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func increaseB() (r int) {
	// 1.赋值
	r = 0
	// 2.闭包引用，r++
	defer func() {
		r++
	}()
	// 3.空 return
	return
}
