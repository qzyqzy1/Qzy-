package main

import (
	"fmt"
	"sync"
	"time"
)

//读写互斥锁
var (
	x  int64
	wg sync.WaitGroup
	//
	rwLock sync.RWMutex
)

func read() {
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	wg.Done()
}
func write() {
	rwLock.Lock()
	x = x + 1
	rwLock.Unlock()
	time.Sleep(time.Millisecond * 20)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
