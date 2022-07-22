package main

import "fmt"

func main() {
	var a1 map[string]int
	fmt.Println(a1 == nil)        //还没有初始化（没有在内存中开辟空间）
	a1 = make(map[string]int, 10) //要估算好该map容量，避免在程序运行期间在动态扩容
	a1["qzy"] = 18
	a1["QZY11"] = 20
	fmt.Println(a1)
	//约定成俗用ok接受返回布尔值
	value, ok := a1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}

	//map的遍历
	for k, v := range a1 {
		fmt.Println(k, v)
	}
	//只遍历key
	for k := range a1 {
		fmt.Println(k)
	}
	//只遍历v
	for _, v := range a1 {
		fmt.Println(v)
	}

}
