package main

import (
	"fmt"
	"os"
)

//学员信息管理系统
//1.添加学员信息
//2.编辑学员信息
//3.展示所有学员信息
//打印
func showMenu() {
	//把所有学生都打印出来
	fmt.Println("---------------欢迎光临学生管理系统-------------")
	fmt.Println(`
			1.新增学生
			2.修改学生信息
			3.展示所有学生
			4.退出
	`)
}

//获取用户输入的学员信息。
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学员信息:")
	fmt.Print("请输入学员的学号")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员的姓名")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的班级")
	fmt.Scanf("%s\n", &class)          //就能拿到用户输入的学员所有信息
	stu := newStudent(id, name, class) //构造函数，造一个学生。
	return stu
}

func main() {
	sm := newStudentMar()
	for { //1.打印系统菜单
		showMenu()
		//2.等待用户选择要执行的选项
		var input int
		fmt.Println("请输入你要输入的序号：")
		fmt.Scanf("%d\n", &input)
		fmt.Println("用户输入的是：", input)
		//3.执行用户选择的动作
		switch input {
		case 1:
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			sm.showStudent()
		case 4:
			os.Exit(1) //退出
		default:
			fmt.Println("滚~")

		}
	}

}
