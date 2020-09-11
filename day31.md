# day31

## 下面这段代码输出什么

```golang
func change(s ...int) {
    s = append(s, 3)
}

func main() {
    slice := make([]int, 5, 5)
    slice[0] = 1
    slice[1] = 2

    change(slice...)
    fmt.Println(slice)

    change(slice[0:2]...)
    fmt.Println(slice)
}
```

## 下面这段代码输出什么呢

```golang
func main() {
    var a = []int{1, 2, 3, 4, 5}
    var r [5]int

    for i, v := range a {
        if i == 0 {
        a[1] = 12
        a[2] = 13
    }

    r[i] = v
    fmt.Println("r = ", r)
    fmt.Println("a = ", a)
}
```
