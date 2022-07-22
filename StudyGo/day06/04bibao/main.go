package main

import "fmt"

//闭包
func f1(f func()) {
	fmt.Printf("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

//底层的原理：
//1.函数可以作为返回值
//2.函数内部查找变量的顺序，现在自己内部找，找不往外层找。
//定义一个函数，对f2进行包装

//要求
//f1(f2)
func f3(x, y int) func() {
	tmp := func() {
		fmt.Println(x + y)
	}
	return tmp
}

func main() {
	f1(f3(100, 200))
}
