# day25

## 下面这段代码输出什么？为什么

```text
A. 1
B. compilation error
```

```go
func (i int) PrintInt() {
 fmt.Println(i)
}

func main() {
    var i int = 1
    i.PrintInt()
}
```

## 下面这段代码输出什么呢？为什么

```text
A. speak
B. compilation error
```

```go
type People interface {
    Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
    if think == "speak" {
        talk = "speak"
    } else {
        talk = "hi"
    }
    return
}

func main() {
    var peo People = Student{}
    think := "speak"
    fmt.Println(peo.Speak(think))
}
```
