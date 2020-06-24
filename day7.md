# day7

## 关于字符串连接，下面语法正确的是

```text
A str := 'abc' + '123'
B str := "abc" + "123"
C str := 'abc' + "abc"
D fmt.Sprintf("abc%d", 123)
```

## 下面这段代码是否编译通过？如果可以，输出什么

```go
const (
    x = iota
    _
    y
    z = "zz"
    k
    p = iota
)

func main() {
    fmt.Println(x, y, z, k, p)
}
```

## 下面赋值正确的是

```text
A var x = nil
B var x interface{} = nil
C var x string = nil
D var x error = nil
```
