package main


import "fmt"



type Student struct{
	name string
	gender string
	age int
	id int
	score float64
}

func (student * Student) say() string {
	infoStr := fmt.Sprintf("student的訊息 name=[%v] gender=[%v] age=[%v] id=[%v] scrore=[%v]",
	student.name,student.gender,student.age,student.id,student.score)
	return infoStr
}

func main() {
	var stu = Student{
		name:"tom",
		gender: "male",
		age:18,
		id: 10000,
		score: 99.98,
	}

	str := stu.say()

	fmt.Println(str)
}