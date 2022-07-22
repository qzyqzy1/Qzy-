package main

import "fmt"

//使用值接收者实现接口 :类型的值和类型的指针都能够保存到接口变量中。
type mover interface {
	move()
}
type person struct {
	name string
	age  int
}

//func (p person) move() {
//	fmt.Printf("%s在跑", p.name)
//}
//使用指针接收者实现接口
func (p *person) move() {
	fmt.Printf("%s在跑", p.name)
}
func main() {
	var m mover
	//p1 := person{"小王子", 18} //值类型
	p2 := &person{"其中国", 20}
	//m = p1 //？ 值
	m = p2
	m.move()
	fmt.Println(m)
}
