package main

import (
	"fmt"
	"strings"
)

func main() {
	//	how do you do 单词出现次数
	s1 := "how do you do"
	s2 := strings.Split(s1, " ")
	//创造一个map存储切片
	m1 := make(map[string]int, 10)
	//遍历切片存储到一个map内
	for _, v := range s2 {
		if _, ok := m1[v]; !ok {
			m1[v] = 1
		} else {
			m1[v]++
		}
	}
	//累加出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)

	}
}
