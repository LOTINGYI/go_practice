package main

import (
	// "encoding/json"
	"fmt"
)


// type Point struct{
// 	x int
// 	y int
// }

// type Rect struct{
// 	leftUp, rightDown Point
// }

// type Rect2 struct{
// 	leftUp, rightDown *Point
// }

// type A struct{
// 	Num int
// }
// type B struct{
// 	Num int
// }

// type Monster struct{
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// 	Skill string `json:"skill"`
// }

// type integer int

// func (i integer) print()  {
// 	fmt.Println("i= ",i)
// }

// func (i *integer) change()  {
// 	*i = *i +1
// }

type Student struct{
	Name string
	Age int
}

func (stu *Student) String2()string  {
	str := fmt.Sprintf("Name=[%v] Age=[%v]",stu.Name,stu.Age)
	return str
}
func main() {

	stu := Student{
		Name:"tom",
		Age: 20,
	}

	fmt.Println(&stu)
	// var a A
	// var b B
	// a = A(b)
	// fmt.Println(a,b)

	// monster := Monster{"牛魔王",500,"芭蕉扇"}
	// // 序列化
	// jsonStr, err := json.Marshal(monster)
	// if err != nil {
	// 	fmt.Println("json處理格式錯誤")
	// }
	// fmt.Println("json Str",string(jsonStr))

// 	r1:=Rect {Point{1,2},Point{3,4}}

// 	// r1有四個int，在內存中連續分布
// 	fmt.Printf("r1.leftUp.x 地址=%p r1.leftUp.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p\n",
// 	&r1.leftUp.x,&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)

// 	r2:=Rect2 {&Point{10,20},&Point{30,40}}
// 	// r2有兩個 *Point，這兩個*Point類型的本身地址也是連續的，但是它們指向的地址不一定是連續的
// 	fmt.Printf("r2.leftUp 本身地址=%p  r2.rightDown 本身地址=%p \n",
// &r2.leftUp,&r2.rightDown)
// 	fmt.Printf("r2.leftUp 指向地址=%p  r2.rightDown 指向地址=%p \n",
// 	r2.leftUp,r2.rightDown)
		// var i integer = 30
		// i.print()
		// (&i).change()
		// fmt.Println("main() i= ",i)
}