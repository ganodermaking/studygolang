# day8

## 关于init函数，下面说法正确的是()

```text
A. 一个包中，可以包含多个 init 函数；
B. 程序编译时，先执行依赖包的 init 函数，再执行 main 包内的 init 函数；
C. main 包中，不能有 init 函数；
D. init 函数可以被其他函数调用；
```

## 下面这段代码输出什么以及原因

```text
A. nil
B. not nil
C. compilation error
```

```golang
func hello() []string {
 return nil
}

func main() {
    h := hello
    if h == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("not nil")
    }
}
```

## 下面这段代码能否编译通过？如果可以，输出什么

```golang
func GetValue() int {
    return 1
}

func main() {
    i := GetValue()
    switch i.(type) {
    case int:
        println("int")
    case string:
        println("string")
    case interface{}:
        println("interface")
    default:
        println("unknown")
    }
}
```
