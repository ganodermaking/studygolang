# day5

## 下面这段代码是否通过编译，不能的话，原因是什么？如果通过，输出什么

```go
func main() {
    sn1 := struct {
        age  int
        name string
    }{
        age:  11,
        name: "qq",
    }

    sn2 := struct {
        age  int
        name string
    }{
        age:  11,
        name: "qq",
    }

    if sn1 == sn2 {
       fmt.Println("sn1 == sn2")
    }

    sm1 := struct {
        age  int
        name map[string]string
    }{
        age:  11,
        name: map[string]string{"a": "1"},
    }

    sm2 := struct {
        age  int
        name map[string]string
    }{
        age:  11,
        name: map[string]string{"a": "1"},
    }

    if sm1 == sm2 {
        fmt.Println("sm1 == sm2")
    }
}
```
