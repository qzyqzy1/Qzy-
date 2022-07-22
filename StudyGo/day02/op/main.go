package main

import "fmt"

func main() {
	var age = 19
	if age < 18 || age > 60 {
		fmt.Println("苦逼上班的")
	} else {
		fmt.Println("不用上班！")
	}

}
