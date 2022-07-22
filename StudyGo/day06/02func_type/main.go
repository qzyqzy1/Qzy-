package main

import "fmt"

//函数类型

func f1() {
	fmt.Println("helloqzy")
}
func f2() int {
	return 10
}
func f3(x func()) {

}
func main() {
	a := f1
	b := f2
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Printf("%T\n", b)
}
