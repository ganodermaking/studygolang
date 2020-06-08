// 1.关于channel的特性，下面说法正确的是？

// A. 给一个 nil channel 发送数据，造成永远阻塞
// B. 从一个 nil channel 接收数据，造成永远阻塞
// C. 给一个已经关闭的 channel 发送数据，引起 panic
// D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值

// 2.下面代码有什么问题？

// const i = 100
// var j = 123

// func main() {
//     fmt.Println(&j, j)
//     fmt.Println(&i, i)
// }

// 3.下面代码能否编译通过？如果通过，输出什么？

// func GetValue(m map[int]string, id int) (string, bool) {
//      if _, exist := m[id]; exist {
//          return "exist", true
//      }
//      return nil, false
//  }

// func main() {
//      intmap := map[int]string{
//         1: "a",
//         2: "b",
//         3: "c",
//     }

//     v, err := GetValue(intmap, 3)
//     fmt.Println(v, err)
// }
