package main

import "fmt"

func main() {
	// age := 19

	// if age > 18 {
	// 	fmt.Println("澳门赌场上线辣")
	// } else {
	// 	fmt.Println("该写暑假作业辣")
	// }

	if age := 19; age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习")
	}

}
