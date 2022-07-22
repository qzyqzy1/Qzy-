package main

//匿名函数
import "fmt"

var f1 = func(x, y int) {
	fmt.Println(x + y)
}

//函数内部没有办法声明带名字的函数，只能定义匿名函数，匿名函数一般运用在函数内部。
func main() {
	f1(10, 20)

	//如果只是调用一次的函数，还可以简写成立即执行函数
	//匿名函数只需要用一次，可以立即执行。
	func(x, y int) {
		fmt.Printf("hhlladasd")

	}(100, 200)
}
