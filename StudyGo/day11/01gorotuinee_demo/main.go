package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//gorotine demo

func hello() {
	fmt.Println("hello 娜扎")
	wg.Done()
}
func main() { //开启一个主goroutine去执行main函数

	wg.Add(1)  //计数牌+1
	go hello() //开启了一个独立的goroutine去执行这个hello这个函数
	fmt.Println("hello main")
	//time.Sleep(time.Second)
	wg.Wait() //等所有小弟都干完活。
}
