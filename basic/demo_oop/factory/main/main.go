package main

import (
	"fmt"
	"go_code/basic/demo_oop/factory/model"
)
func main() {
	// var stu = model.Student {
	// 	Name:"tom",
	// 	Score: 78.9,
	// }
	var stu = model.NewStudent("tom!",98.8)
	fmt.Println(*stu)
	fmt.Println("name= ",stu.Name," scores= ",stu.GetScore())
}