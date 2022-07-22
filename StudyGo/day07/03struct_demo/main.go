package main

import "fmt"

//结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个Person类型的变量P
	var qzy person
	//通过字段赋值
	qzy.name = "qizhongyi"
	qzy.age = 65
	qzy.gender = "男"
	qzy.hobby = []string{"篮球", "RAP"}
	fmt.Println(qzy)

	var s struct {
		name   string
		age    int
		gender string
	}
	s.age = 19
	A1 := s.age
	fmt.Println(A1)
}
