package main

import "fmt"


// type Person struct{
// 	Name string
// 	Age int
// 	Scores [5] float64
// 	ptr *int
// 	slice []int
// 	map1 map[string]string
// }

// type Monster struct{
// 	Name string
// 	Age int
// }

type Person struct{
	Name string
	Age int
}
func main() {
	

	// var p1 Person
	// fmt.Println(p1)

	// if p1.ptr == nil {
	// 	fmt.Println("ok1")
	// }
	// if p1.slice == nil {
	// 	fmt.Println("ok2")
	// }
	// if p1.map1 == nil {
	// 	fmt.Println("ok3")
	// }

	// // 使用slice再次說明一定要用make
	// p1.slice = make([]int, 2)
	// p1.slice[0] = 100
	// fmt.Println(p1)

	// // 使用map一定要用make
	// p1.map1 = make(map[string]string)
	// p1.map1["key1"] = "tom"
	// fmt.Println(p1)

	// var monster1 Monster
	// monster1.Name = "牛魔王"
	// monster1.Age = 500

	// monster2  := &monster1
	// monster2.Name = "fuck"
	// fmt.Println("monster1 = ",monster1)
	// fmt.Println("monster2 = ",*monster2)

	var p1 Person
	p1.Age = 10
	p1.Name ="tim"

	p2 := &p1

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)
	p2.Name = "tom~"

	fmt.Printf("p2.Name=%v p1.Name=%v\n",p2.Name,p1.Name)
	fmt.Printf("p2.Name=%v p1.Name=%v\n",(*p2).Name,p1.Name)
	fmt.Printf("p1的地址=%p\n", &p1)
	fmt.Printf("p2指向的地址=%p p2的值=%p", p2,&p2)
}