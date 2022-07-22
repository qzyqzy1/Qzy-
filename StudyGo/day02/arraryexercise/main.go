package main

import "fmt"

//遍历数组 并累加
func main() {
	a1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a1 {
		sum = sum + v
	}
	fmt.Println(sum)
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}

		}
	}
}

//求出数组中和为指定值的下标，指定值为8
//定义两个for循环，外层从第一个开始遍历，内层循环从外层后面的那个开始找，直到和为8
