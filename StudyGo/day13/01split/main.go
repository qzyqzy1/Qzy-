package split

import "strings"

//split.go
//Split 将s按照sep进行分割，返回一个字符串的切片
//Split(”我爱你“,"爱") => ["我","你"]
func Split(s, sep string) (ret []string) {
	index := strings.Index(s, sep)
	for index > -1 {
		ret = append(ret, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
