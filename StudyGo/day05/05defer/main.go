package main

import "fmt"

//defer多用于释放资源（文件句柄，数据库链接，socket链接）
func deferdemo() {
	fmt.Println("start")
	defer fmt.Println("hhhhh") //把它后面的语句延迟到函数即将返回的时候执行
	fmt.Println("end")
}
func main() {
	deferdemo()
}
