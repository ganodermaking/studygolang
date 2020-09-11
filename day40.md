# day40

## 关于select机制，下面说法正确的是?​

A. select机制用来处理异步IO问题；
B. select机制最大的一条限制就是每个case语句里必须是一个IO操作；
C. golang在语言级别支持select关键字；
D. select关键字的用法与switch语句非常类似，后面要带判断条件；

## 下面的代码有什么问题？​

```golang
func Stop(stop <-chan bool) {
    close(stop)
}
```

## 下面这段代码存在什么问题？​

```golang
type Param map[string]interface{}

type Show struct {
    *Param
}

func main() {
    s := new(Show)
    s.Param["day"] = 2
}
```
