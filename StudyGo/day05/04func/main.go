package main

import "fmt"

//函数：一段代码的封装
func f1() {
	fmt.Println("hello QZY")
}
func f2(name string) {
	fmt.Println("hello", name)
}
func f3(x int, y int) int {
	sum := x + y
	return sum
}

func main() {
	f1()
	f2("QZYYY")
	f3(100, 300) //调用函数
	ret := f3(100, 200)
	fmt.Println(ret)
}
