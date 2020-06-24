# day9

## 关于channel，下面语法正确的是()

```text
A. var ch chan int
B. ch := make(chan int)
C. <- ch
D. ch <-
```

## 下面这段代码输出什么

```text
A.0
B.1
C.Compilation error
```

```go
type person struct {
    name string
}

func main() {
    var m map[person]int
    p := person{"mike"}
    fmt.Println(m[p])
}
```

## 下面这段代码输出什么呢

```text
A.18
B.5
C.Compilation error
```

```go
func hello(num ...int) {
    num[0] = 18
}

func main() {
    i := []int{5, 6, 7}
    hello(i...)
    fmt.Println(i[0])
}
```
