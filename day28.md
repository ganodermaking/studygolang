# day28

## 下面的代码有什么问题

```golang
func main() {
    fmt.Println([...]int{1} == [2]int{1})
    fmt.Println([]int{1} == []int{1})
}
```

## 一道很有代表性的题目，很多老司机都因此翻车！下面这段代码输出什么？如果编译错误的话，为什么

```text
A. 5 5
B. runtime error
```

```golang
var p *int

func foo() (*int, error) {
    var i int = 5
    return &i, nil
}

func bar() {
    fmt.Println(*p)
}

func main() {
    p, err := foo()
    if err != nil {
        fmt.Println(err)
        return
    }
    bar()
    fmt.Println(*p)
}
```
