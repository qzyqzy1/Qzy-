package main

import "fmt"

//panic和 recover
func funcA() {
	fmt.Println("a")
}
func funcB() {
	panic("出现了严重的错误")
	fmt.Println("b")
}
func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
