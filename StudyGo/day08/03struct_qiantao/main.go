package main

import "fmt"

//嵌套结构体
type address struct {
	province string
	city     string
}
type person struct {
	name    string
	age     int
	address address
}
type company struct {
	name    string
	address address
}

func main() {
	p1 := person{
		name: "qzq",
		age:  123,
		address: address{
			province: "危害",
			city:     "dad"},
	}
	fmt.Println(p1)
	fmt.Println()
}
