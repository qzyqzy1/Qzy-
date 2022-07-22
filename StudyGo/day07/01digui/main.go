package main

import "fmt"

//递归
// 递归适合处理那种问题相同\问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件
//5! = 1*2*3*4*5
//计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
func main() {
	ret := f(5)
	fmt.Println(ret)

}
