# day20

## 下面这段代码正确的输出是什么

```text
A. F M D
B. D F M
C. F D M
```

```golang
func f() {
 defer fmt.Println("D")
 fmt.Println("F")
}

func main() {
 f()
 fmt.Println("M")
}
```

## 下面代码输出什么

```golang
type Person struct {
    age int
}

func main() {
    person := &Person{28}

    // 1.
    defer fmt.Println(person.age)

    // 2.
    defer func(p *Person) {
        fmt.Println(p.age)
 }(person)

 // 3.
 defer func() {
  fmt.Println(person.age)
 }()

 person = &Person{29}
}
```
