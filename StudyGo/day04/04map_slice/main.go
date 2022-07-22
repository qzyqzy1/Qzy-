package main

import "fmt"

//map和slice组合
func main() {

	//元素类型为map的切片
	var s1 = make([]map[int]string, 0, 10)

	s1[0][100] = "A"

	fmt.Println(s1)

}
