package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性与JSON序列化
//如果一个Go语言包中定义的标识符是首字母大写的，那么就是对外可见的。
//如果一个结构体中的字段名首字母是大写的，那么该字段就是对外可见的。
type student struct {
	Id   int
	Name string
}

//student构造函数
func newStudent(id int, name string) student {
	return student{Id: id, Name: name}
}

type class struct {
	Title    string
	Students []student
}

func main() {
	//创建一个班级变量c1
	c1 := class{Title: "火箭班",
		Students: make([]student, 0, 20)}
	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("Stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
		//造10个学生
	}
	fmt.Printf("%#v\n", c1)
	//json序列化
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Printf("json marshal failed,err", err)
		return
	}
	fmt.Printf("%s\n", data)
	//JSON反序列化
	jsonStr := `{"Title":"火箭班","Students":[{"Id":0,"Name":"Stu00"},{"Id":1,"Name":"易建联"},{"Id":2,"Name":"Stu02"},{"Id":3,"Name":"Stu03"}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Printf("json unmarshal failed,err")
		return
	}
	fmt.Printf("%#v\n", c2)
}
