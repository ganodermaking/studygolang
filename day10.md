# day10

## 下面这段代码输出什么

```text
A.13.1
B.13
C.compilation error
```

```golang
func main() {
    a := 5
    b := 8.1
    fmt.Println(a + b)
}
```

## 下面这段代码输出什么呢

```text
A.3
B.4
C.compilation error
```

```golang
func main() {
    a := [5]int{1, 2, 3, 4, 5}
    t := a[3:4:4]
    fmt.Println(t[0])
}
```

## 下面这段代码输出什么呢呢

```text
A. compilation error
B. equal
C. not equal
```

```golang
func main() {
    a := [2]int{5, 6}
    b := [3]int{5, 6}
    if a == b {
        fmt.Println("equal")
    } else {
        fmt.Println("not equal")
    }
}
```
