package main

import "fmt"

//函数
//函数的定义
//函数是一段代码的封装，把一段代码（逻辑）抽象出来，封装到一段函数。每次用到它的时候直接用函数名调用
//使用函数能够让代码结构更清晰，简洁。

func sum(x int, y int) (ret int) {
	return x + y
}

//没有返回值的函数
func sum1(x int, y int) {
	fmt.Println(x + y)
}
func f2() {
	fmt.Println(f2)
}

//命名返回值相当于在函数中声明了一个变量。
func f3() int {
	return 3
}

func f5() (int, int) {
	return 1, 2
}

//参数的类型简写

func f6(x, y int) int {
	return x + y
}
func main() {
	r := sum(10, 2)
	fmt.Println(r)
}
