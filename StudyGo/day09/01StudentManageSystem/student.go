package main

import "fmt"

type student struct {
	id    int //学号是唯一的
	name  string
	class string
}

//student的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		name:  name,
		id:    id,
		class: class}
}

//学员管理的类型
type studentMar struct {
	allStudent []*student
}

// studentMar的构造函数
func newStudentMar() *studentMar {
	return &studentMar{
		allStudent: make([]*student, 0, 100),
	}
}

//添加学生                        //接收的是一个新的学生对象。
func (s *studentMar) addStudent(newStu *student) {
	s.allStudent = append(s.allStudent, newStu)
}

//编辑学生
func (s *studentMar) modifyStudent(newStu *student) {
	for i, v := range s.allStudent {
		if newStu.id == v.id { //当学号相同时，就表示找到了要修改的学生
			s.allStudent[i] = newStu //根据切片的索引直接把新学生赋值进来。
			return
		}
	}
	//如果走到这里，则表示输入的学生找不到
	fmt.Println("学生信息有无，系统中没有学号是:%d的学生\n", newStu.id)
}

//展示学生
func (s *studentMar) showStudent() {
	for _, v := range s.allStudent {
		fmt.Printf("学号%d 姓名%s 班级%s\n", v.id, v.name, v.class)
	}
}
