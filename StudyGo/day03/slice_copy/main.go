package main

import "fmt"

//	copy（）函数复制切片
func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)
	//将a1中的索引为1的3这个元素删掉
	a1 = append(a1[:1], a2[2:]...)
	fmt.Println(a1)
	fmt.Println(cap(a1))
}
