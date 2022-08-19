// 常量是一个简单的标识符，在程序运行时，不会被修改的量。
// 常量中的数据类型可以是布尔型、数字型（整数、浮点和复数型）和字符串型
// 定义格式：  const identifier [type] = value

// 常量
// package main

// import "fmt"

// func main(){
// 	const LENGTH int = 10
// 	const WIDTH int = 5
// 	var area int
// 	const a, b, c  = 1, false, "str"

// 	area = LENGTH * WIDTH
// 	fmt.Println("面积为 : %d",area)
// 	println()
// 	println(a,b,c)
// }


// 枚举
/*
package main

import "unsafe"
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	println(a,b,c)
}
*/

// iota
/*iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，
const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。*/ 

package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}