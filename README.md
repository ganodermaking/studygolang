# go-question-bank
go question bank

### f290
考答案及解析：
不会出现死循环，能正常结束。
循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。

### f291
参考答案及解析：
```
2 3
2 3
2 3
```

for range使用短变量声明(:=)的形式迭代变量，需要注意的是，变量i、v在每次循环体中都会被重用，而不是重新声明。各个goroutine中输出的i、v值都是for range循环结束后的i、v最终值，而不是各个goroutine启动时的i, v值。可以理解为闭包引用，使用的是上下文环境的值。
两种可行的 fix 方法:

1.使用函数传递
```go
for i, v := range m {
    go func(i,v int) {
        fmt.Println(i, v)
    }(i,v)
}
```

2.使用临时变量保留当前值
```go
for i, v := range m {
    i := i           // 这里的 := 会重新声明变量，而不是重用
    v := v
    go func() {
        fmt.Println(i, v)
    }()
}
```

### f300
参考答案及解析：7。根据 5 年 Gopher 都不知道的 defer 细节，你别再掉进坑里！ 提到的“三步拆解法”，第一步执行r = n +1，接着执行第二个 defer，由于此时 f() 未定义，引发异常，随即执行第一个 defer，异常被 recover()，程序正常执行，最后 return。

### f301
参考答案及解析：
```
r =  [1 2 3 4 5]
a =  [1 12 13 4 5]
```

range 表达式是副本参与循环，就是说例子中参与循环的是 a 的副本，而不是真正的 a。就这个例子来说，假设 b 是 a 的副本，则 range 循环代码是这样的：
```go
for i, v := range b {
    if i == 0 {
        a[1] = 12
        a[2] = 13
    }
    r[i] = v
}
```

因此无论 a 被如何修改，其副本 b 依旧保持原值，并且参与循环的是 b，因此 v 从 b 中取出的仍旧是 a 的原值，而非修改后的值。
如果想要 r 和 a 一样输出，修复办法：
```go
func main() {
    var a = [5]int{1, 2, 3, 4, 5}
    var r [5]int

    for i, v := range &a {
        if i == 0 {
            a[1] = 12
            a[2] = 13
        }
        r[i] = v
    }
    fmt.Println("r = ", r)
    fmt.Println("a = ", a)
}
```

输出：
```
r =  [1 12 13 4 5]
a =  [1 12 13 4 5]
```
修复代码中，使用 *[5]int 作为 range 表达式，其副本依旧是一个指向原数组 a 的指针，因此后续所有循环中均是 &a 指向的原数组亲自参与的，因此 v 能从 &a 指向的原数组中取出 a 修改后的值。
