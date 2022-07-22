package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

//go语言中函数传参永远是拷贝，即为副本。
func f(x person) {
	x.gender = "女"
}
func f2(x *person) {
	(*x).gender = "女"
}
func main() {
	var p person
	p.name = "qzy"
	p.gender = "男"
	f(p)
	fmt.Println(p)
}
