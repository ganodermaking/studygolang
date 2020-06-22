# day22

## 下面的代码有几处语法问题，各是什么

```go
func main() {
    var x string = nil
    if x == nil {
        x = "default"
    }
    fmt.Println(x)
}

var a bool = true
```

## return 之后的 defer 语句会执行吗，下面这段代码输出什么

```go
func main() {
    defer func() {
        fmt.Println("1")
    }()

    if a == true {
        fmt.Println("2")
        return
    }

    defer func() {
        fmt.Println("3")
    }()
}
```
