package main

import "fmt"

//为什么需要接口

type dog struct {
}

func (d dog) say() { fmt.Println("汪汪汪") }

type cat struct {
}

func (c cat) say() { fmt.Println("喵喵喵") }

type person struct {
	name string
}

func (p person) say() { fmt.Println("啊啊啊") }

//定义一个类型，一个抽象的类型。 只要实现了say 这个方法的类型，都可以成为sayer这个类型
type sayer interface {
	say()
}

//打的函数
func da(arg sayer) {
	arg.say()
}

func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)
	p1 := person{"娜扎"}
	da(p1)
}
