package main

import "fmt"

//你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
//分配规则如下：
//a. 名字中每包含'e'或'E'分1枚金币
//b. 名字中每包含'i'或'I'分2枚金币
//c. 名字中每包含'o'或'O'分3枚金币
//d: 名字中每包含'u'或'U'分4枚金币
//写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
//程序结构如下，请实现 ‘dispatchCoin’ 函数
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin(left int) int {
	//	1.依次拿到每个人的名字
	for _, name := range users {
		for _, c := range name {
			switch c {
			case 'e', 'E': //满足这个条件分1枚金币
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	//	2.拿到一个人名，根据分金币的规则去分金币。
	//	2.1每个人分的金币数应该保存到distribution中
	//	2.2还要记录下剩余的金币数
	left = coins
	return left
}
func main() {
	left := dispatchCoin(50)
	fmt.Println("剩下：", left)
}