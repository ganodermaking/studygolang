# day11

## 关于 cap() 函数的适用类型，下面说法正确的是()

```text
A. array
B. slice
C. map
D. channel
```

## 下面这段代码输出什么

```text
A. nil
B. not nil
C. compilation error
```

```golang
func main() {
    var i interface{}
    if i == nil {
        fmt.Println("nil")
        return
    }
    fmt.Println("not nil")
}
```

## 下面这段代码输出什么呢

```text
A. runtime panic
B. 0
C. compilation error
```

```golang
func main() {
    s := make(map[string]int)
    delete(s, "h")
    fmt.Println(s["h"])
}
```
