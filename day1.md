# day1

## 下面这段代码输出什么

```go
func main() {
    defer func() {
        fmt.Println("打印前")
    }()
    defer func() {
        fmt.Println("打印中")
    }()
    defer func() {
        fmt.Println("打印后")
    }()
}
```
