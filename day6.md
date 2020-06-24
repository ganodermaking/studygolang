# day6

## 通过指针变量p访问其成员变量name，有哪几种方式

```text
A p.name
B (&p).name
C (*p).name
D p->name
```

## 下面这段代码是否能通过编译？如果通过，输出什么

```go
type MyInt1 int
type MyInt2 = int

func main() {
    var i int = 0
    var i1 MyInt1 = i
    var i2 MyInt2 = i
    fmt.Println(i1, i2)
}
