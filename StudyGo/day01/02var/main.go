package main

import "fmt"

//Go语言中推荐使用驼峰式命名
//声明变量

// var name string
// var age int
// var isOK bool

//批量声明变量

var (
	name string //""
	age  int    //0
	isOK bool   // false
)

func main() {

	name = "qzy"
	age = 17
	isOK = true

	//Go语言变量声明必须使用，不使用就编译不过去
	fmt.Printf("name:%s", name) //%s：占位符  使用name这个变量的值去替换占位符
	fmt.Println(age)            //打印完指定的内容之后会加一个换行符

	// 声明变量同时赋值
	var s1 string = "qzy"
	fmt.Println(s1)

	//类型推导

	var s2 = "qzy"
	fmt.Println(s2)

	//简短变量声明，只能在函数内使用 简短声明。
	s3 := "hhhhh"
	fmt.Println(s3)

	//匿名变量，_ 使用下划线使用

}
