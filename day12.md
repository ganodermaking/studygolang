# day12

## 下面属于关键字的是（）

```text
A.func
B.struct
C.class
D.defer
```

## 下面这段代码输出什么

```text
A. -5 +5
B. +5 +5
C.  0  0
```

```go
func main() {
    i := -5
    j := +5
    fmt.Printf("%+d %+d", i, j)
}
```

## 下面这段代码输出什么呢

```go
type People struct{}

func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func (p*People) ShowB() {
    fmt.Println("showB")
}

type Teacher struct {
    People
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}

func main() {
    t := Teacher{}
    t.ShowB()
}
```
