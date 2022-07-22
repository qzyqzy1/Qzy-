package main

import "fmt"

// 常量 定义了常量之后就不能修改 在程序运行期间不会改变的量

const pi = 3.1415926

//批量声明常量
const (
	status0  = 200
	notfound = 404
)

// 批量时下面没有写值，则默认和上方一样
const (
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota //0
	a2
	a3
	a4
)
const (
	b1 = iota
	b2

	_
)

func main() {
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}
