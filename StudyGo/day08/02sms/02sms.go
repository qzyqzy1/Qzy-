package main

import (
	"fmt"
	"os"
)

/*
函数版学生管理系统
写一系统能够查看\新增\删除学生
*/
var (
	allStudent map[int64]*student
)

type student struct {
	id   int64
	name string
}

//newStudent 是student的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		name: name,
		id:   id}
}
func showALLStudent() {
	//把所有学生都打印出来
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名：%s\n", k, v.name)
	}

}

func addStudent() {
	//	 向allstudent中添加一个学生
	//1.创建一个新学生
	//1.1获取用户输入

	var (
		id   int64
		name string
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名")
	fmt.Scanln(&name)
	//1.2造学生(调用student的构造函数)
	newStu := newStudent(id, name)
	allStudent[id] = newStu

	//2.追加到allStudent这个map中
}
func deleteStudent() {
	var (
		deleteID int64
	)
	//1.请用户输入要删除的学生的序号
	fmt.Println("请输入要删除的学生学号：")
	fmt.Scanln(&deleteID)

	//2.去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteID)

}

func main() {
	allStudent = make(map[int64]*student, 50) //初始化map,开辟内存空间
	//1.打印菜单
	for {
		fmt.Println("---------------欢迎光临学生管理系统-------------")
		fmt.Println(`
			1.查看所有学生
			2.新增学生
			3.删除学生
			4.退出
		`)

		//2.等待用户选择
		fmt.Print("请输入你要干啥：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		//3.执行对应函数
		switch choice {
		case 1:
			showALLStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) //退出
		default:
			fmt.Println("滚~")

		}
	}
}
