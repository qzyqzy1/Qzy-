package main

import "fmt"

//结构体占用一块连续的内存
type x struct {
	a int8
	b int8
	c int8
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Println("%p\n", &(m.a))
	fmt.Println("%p\n", &(m.b))
	fmt.Println("%p\n", &(m.c))
}
