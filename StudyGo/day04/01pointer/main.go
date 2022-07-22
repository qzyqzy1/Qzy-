package main

//指针
import "fmt"

func main() {
	//1. &:取地址

	n := 18
	p := &n
	fmt.Println(&n)
	fmt.Printf("%T\n", p)
	//	2.*：根据地址取值
	m := *p
	fmt.Println(m) //int

	//
	var a = new(int) //nill pointer panic报错
	*a = 100
	fmt.Println(*a)
	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)
	//new函数申请一个内存地址

}
