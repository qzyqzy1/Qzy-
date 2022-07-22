package main

import "fmt"

//回文判断，字符串从左往右读和从右往左读是一样的，那么就是回文
//上海自来水来自海上
//山西运煤车煤运西山
//黄山落叶松叶落山黄
func main() {
	//	解题思路：
	//
	ss := "上海自来水来自海上"
	//把字符串中的字符拿出来放到一个[]rune中
	s1 := make([]rune, 0, len(ss))
	for _, c := range ss {
		s1 = append(s1, c)
	}
	for i := 0; i < len(s1)/2; i++ {
		if s1[i] != s1[len(s1)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
