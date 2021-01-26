package main


import "fmt"

type Student struct{
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("學生名=%v 年齡=%v 成績=%v\n",stu.Name,stu.Age,stu.Score)
}

func (stu * Student) SetScore(score int){
	stu.Score = score
}

type Pupil struct{
	Student
}
func (p *Pupil) testing()  {
	fmt.Println("小學生在考試中")
}
type Graduate struct{
	Student
}
func (p *Graduate) testing()  {
	fmt.Println("大學生在考試中")
}


func main() {
	// 當我們對struct嵌入匿名struct
	pupil := &Pupil{}
	pupil.Student.Name = "tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()

}