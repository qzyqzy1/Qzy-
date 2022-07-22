package main

import (
	"fmt"
	"sync"
)

//多个goroutine并发操作全局变量
var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex //互斥锁
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
