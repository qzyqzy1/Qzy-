package main

import (
	"fmt"
	"sync"
)

//sync.Map 并发安全的map
var m = make(map[int]int)
var wg sync.WaitGroup

func get(key int) int {
	return m[key]
}
func set(key int, value int) {
	m[key] = value
}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			set(i, i+100)                            //map中设置键值对
			fmt.Printf("key:%v value:%v", i, get(i)) //打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}
