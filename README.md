# studygolang-daily-interview
【go语言中文网】每日面试题及其解析

## day1
```
打印后
打印中
打印前
```
参考解析：defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic。

### day2
```
0 -> 3
1 -> 3
2 -> 3
3 -> 3
```
参考解析：这是新手常会犯的错误写法，for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.

正确的写法：
```
func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "===>", *v)
	}
}
```

### day3
1. 两段代码分别输出：
```
1[0 0 0 0 0 1 2 3]
2[1 2 3 4]
```
参考解析：这道题考的是使用 append 向 slice 添加元素，第一段代码常见的错误是 [1 2 3]，需要注意。

2. 参考答案：第二个返回值没有命名。
参考解析：
在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。这里的第一个返回值有命名 sum，第二个没有命名，所以错误。

3. 参考答案：
new(T) 和 make(T, args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。

make(T, args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.


### day4
1. 参考答案及解析：不能通过编译，new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作。可以使用 make() 初始化之后再用。同样的，map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new() 。

2. 参考答案及解析：不能通过编译。append() 的第二个参数不能直接使用 slice，需使用 … 操作符，将一个切片追加到另一个切片上：append(s1,s2…)。或者直接跟上元素，形如：append(s1,1,2,3)。

3. 参考答案及解析：不能通过编译。这道题的主要知识点是变量声明的简短模式，形如：x := 100。但这种声明方式有限制：

必须使用显示初始化；
不能提供数据类型，编译器会自动推导；
只能在函数内部使用简短模式；

### day5
编译不通过 invalid operation: sm1 == sm2
这道题目考的是结构体的比较，有几个需要注意的地方：
结构体只能比较是否相等，但是不能比较大小。
相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关，sn3 与 sn1 就是不同的结构体；
```
sn3 := struct {
	name string
	age  int
}{age: 11, name: "qq"}
```

如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
那什么是可比较的呢，常见的有 bool、数值型、字符、指针、数组等，像切片、map、函数等是不能比较的。

### day6
Go面试题# 第6天答案解析，题目见截图
1. 参考答案及解析：AC。& 取址运算符，* 指针解引用。
2. 参考答案及解析：编译不通过，cannot use i (type int) as type MyInt1 in assignment。这道题考的是类型别名与类型定义的区别。第 5 行代码是基于类型 int 创建了新类型 MyInt1，第 6 行代码是创建了 int 的类型别名 MyInt2，注意类型别名的定义时 = 。所以，第 10 行代码相当于是将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过；而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值。第 10 行代码的赋值可以使用强制类型转化 var i1 MyInt1 = MyInt1(i).

### day7
1. 参考答案及解析：BD。知识点：字符串连接。除了以上两种连接方式，还有 strings.Join()、buffer.WriteString()等。

2. 参考答案及解析：编译通过，输出：0 2 zz zz 5。知识点：iota 的使用。给大家贴篇文章，讲的很详细
icongolang 使用 iota - 迪克猪 - 博客园

3. 参考答案及解析：BD。知识点：nil 值。nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。强调下 D 选项的 error 类型，它是一种内置接口类型，看下方贴出的源码就知道，所以 D 是对的。
```
type error interface {
	Error() string
}
```

### day8 
1. 参考答案及解析：AB。关于 init() 函数有几个需要注意的地方：
1）init() 函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等;
2）一个包可以出线多个 init() 函数,一个源文件也可以包含多个 init() 函数；
3）同一个包中多个 init() 函数的执行顺序没有明确定义，但是不同包的init函数是根据包导入的依赖关系决定的（看下图）;
4）init() 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误;
5）一个包被引用多次，如 A import B,C import B,A import C，B 被引用多次，但 B 包只会初始化一次；
6）引入包，不可出现死循坏。即 A import B,B import A，这种情况编译失败；

2. 答案及解析：B。这道题目里面，是将 hello() 赋值给变量 h，而不是函数的返回值，所以输出 not nil。

3. 参考答案及解析：编译失败。考点：类型选择，类型选择的语法形如：i.(type)，其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型选择。

### day9
1. 参考答案及解析：ABC。A、B都是声明 channel；C 读取 channel；写 channel 是必须带上值，所以 D 错误。

2. 参考答案及解析：A。打印一个 map 中不存在的值时，返回元素类型的零值。这个例子中，m 的类型是 map[person]int，因为 m 中不存在 p，所以打印 int 类型的零值，即 0。

3. 参考答案及解析：18。知识点：可变函数。

### day10
1. 参考答案及解析：C。a 的类型是 int，b 的类型是 float，两个不同类型的数值不能相加，编译报错。

2. 参考答案及解析：B。知识点：操作符 [i,j]。基于数组（切片）可以使用操作符 [i,j] 创建新的切片，从索引 i，到索引 j 结束，截取已有数组（切片）的任意部分，返回新的切片，新切片的值包含原数组（切片）的 i 索引的值，但是不包含 j 索引的值。i、j 都是可选的，i 如果省略，默认是 0，j 如果省略，默认是原数组（切片）的长度。i、j 都不能超过这个长度值。假如底层数组的大小为 k，截取之后获得的切片的长度和容量的计算方法：长度：j-i，容量：k-i。截取操作符还可以有第三个参数，形如 [i,j,k]，第三个参数 k 用来限制新切片的容量，但不能超过原数组（切片）的底层数组大小。截取获得的切片的长度和容量分别是：j-i、k-i。所以例子中，切片t为[4]，长度和容量都是1。

3. 参考答案及解析：A。Go 中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的组成部分，所以a和b是不同的类型，是不能比较的，所以编译错误。

### day11
1. 参考答案及解析：ABD。知识点：cap()，cap() 函数不适用 map。

2. 参考答案及解析：A。当且仅当接口的动态值和动态类型都为nil时，接口类型值才为nil。

3. 参考答案及解析：B。删除 map 不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的减值对时，返回值类型对应的零值，所以返回 0。

### day12
1. 参考答案及解析：ABD。知识点：Go 语言的关键字。Go语言有25个关键字。
2. 参考答案及解析：A。%d表示输出十进制数字，+表示输出数值的符号。这里不表示取反。
3. 参考答案及解析：teacher showB。知识点：结构体嵌套。在嵌套结构体中，People 称为内部类型，Teacher 称为外部类型；通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。此外，外部类型还可以定义自己的属性和方法，甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。这个例子中的 ShowB() 就是同名方法。

### day13
1. 参考答案及解析：AD。B 只支持局部变量声明；C 是赋值，str 必须在这之前已经声明；
2. 参考答案及解析：5。这个例子中，hello() 函数的参数在执行defer语句的时候会保存一份副本，在实际调用hello()函数时用，所以是5。
3. 参考答案及解析：
```
showA
showB
```
知识点：结构体嵌套。这道题可以结合第12天的第三题一起看，Teacher 没有自己ShowA()，所以调用内部类型 People的同名方法，需要注意的是第5行代码调用的是People自己的ShowB方法。

### day14
1. 参考代码及解析：C。知识点：常量，Go 语言中的字符串是只读的。
2. 参考答案及解析：B。知识点：指针，incr() 函数里的 p 是 *int 类型的指针，指向的是 main() 函数的变量 p 的地址。第 2 行代码是将该地址的值执行一个自增操作，incr() 返回自增后的结果。
3. 参考答案及解析：ABD。知识点：可变函数。

### day15
1. 参考答案及解析：A。知识点：nil 切片和空切片。nil 切片和 nil 相等，一般用来表示一个不存在的切片；空切片和 nil 不相等，表示一个空的集合。
2. 参考答案及解析：A。UTF-8 编码中，十进制数字 65 对应的符号是 A。
3. 参考答案及解析：13 23。知识点：接口。一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。