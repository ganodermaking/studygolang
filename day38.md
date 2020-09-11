# day38

## 关于异常的触发，下面说法正确的是

```text
A. 空指针解析；
B. 下标越界；
C. 除数为0；
D. 调用panic函数；
```

## 下面代码输出什么

```golang
func main() {
     x := []string{"a", "b", "c"}
     for v := range x {
         fmt.Print(v)
     }
}

func main() {
    x := []string{"a", "b", "c"}
    for _, v := range x {
        fmt.Print(v)
    }
}
```

## 下面这段代码能否编译通过？如果通过，输出什么

```golang
type User struct{}
type User1 User
type User2 = User

func (i User1) m1() {
    fmt.Println("m1")
}

func (i User) m2() {
    fmt.Println("m2")
}

func main() {
    var i1 User1
    var i2 User2
    i1.m1()
    i2.m2()
}
```
