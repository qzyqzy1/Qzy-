package main

import "fmt"

//切片

func main() {
	var s1 []int    //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放string类型元素的切片
	fmt.Println(s1, s2)

	//切片的初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平沙村"}
	fmt.Println(s1, s2)
	//切片的长度和容量
	//len函数求长度，cap函数求容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))
	//2.由数组得到切片
	a1 := [...]int{1, 12, 3, 213, 1, 45, 32, 534}
	s3 := a1[0:4] //[1,3,5,7] 基于一个数组切割，左闭右开
	fmt.Println(s3)
	//切片的容量是从切的第一个到底层数组的最后一个容量

}
