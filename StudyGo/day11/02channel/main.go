package main

import "fmt"

//channel
func main() {
	var ch1 chan int        //引用类型，需要初始化之后才能使用
	ch1 = make(chan int, 1) //1，缓冲区
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
	close(ch1)

}
