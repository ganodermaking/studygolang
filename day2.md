# day2

## 如下代码输出什么？说明原因

```go
func main() {
    slice := []int{0, 1, 2, 3}
    m := make(map[int]*int)

    for key, val := range slice {
        m[key] = &val
    }

    for k, v := range m {
        fmt.Println(k, "->", *v)
    }
}
```
