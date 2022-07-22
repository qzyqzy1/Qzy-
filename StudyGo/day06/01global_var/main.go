package main

import "fmt"

//变量的作用域
var x = 100 //定义一个全局变量

//定义一个函数
func f1() {
	fmt.Println(x)
}
func main() {
	f1()
}
