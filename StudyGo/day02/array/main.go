package main

import "fmt"

//数组
// 存放元素的容器
// 必须指定存放的元素的类型和容量（长度）
func main() {
	var a1 [3]bool //[true false true]
	var a2 [4]bool //[true true false false]
	// a1 a2不能比较，长度与容量是数组类型的一部分，类型不同，不能比较

	fmt.Printf("a1:%T a2:%T\n ", a1, a2)

	//数组的初始化,
	// 如果不初始化默认元素都是零值
	fmt.Println(a1, a2)
	// 初始化方式1
	a1 = [3]bool{false, true, true}
	fmt.Println(a1)
	//初始化方式2，根据初始值自动推断数组的长度是多少
	a100 := [...]int{12, 3, 213, 13, 21, 31, 31, 3, 215, 34, 645, 6, 547, 78, 79, 79, 9, -8, 6, 432}
	fmt.Println(a100)
	//初始化方式3,根据索引初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	//数组的遍历
	//1.根据索引遍历
	citys := [...]string{"北京", "上海", "深圳"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//2.for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}
	//多维数组
	//[[1,2,][3,4][5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)
	//多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1)
	fmt.Println(b2)

}
