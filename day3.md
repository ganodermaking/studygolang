# day3

## 下面两段代码输出什么

```go
func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)

    s1 := make([]int, 0)
    s1 = append(s1, 1, 2, 3, 4)
    fmt.Println(s1)
}
```

## 下面这段代码有什么缺陷

```go
func funcMui(x, y int) (sum int, error) {
    return x + y, nil
}
```

## new() 和 make() 的区别
