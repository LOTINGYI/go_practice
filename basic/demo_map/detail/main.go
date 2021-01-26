package main

import "fmt"


func modify(map1 map[int]int){
	map1[10] = 900
}


type Stu struct{
	Name string
	Age int
	Address string
}
func main() {
	// map1 := make(map[int]int,2)

	// map1[1] = 90
	// map1[2] = 88
	// map1[10] = 1
	// map1[20] = 2

	// fmt.Println(map1)
	// modify(map1)
	// fmt.Println(map1)

	// students := make(map[string]Stu, 10)
	// stu1 := Stu{"tom",18,"Taipei"}
	// stu2 := Stu{"julia",28,"Taipei"}
	// students["no1"] = stu1
	// students["no2"] = stu2
	// fmt.Println(students)


	// for k,v := range students {
	// 	fmt.Printf("學生的編號%v\n",k)
	// 	fmt.Printf("學生的名字%v\n",v.Name)
	// 	fmt.Printf("學生的年齡%v\n",v.Age)
	// 	fmt.Printf("學生的地址%v\n",v.Address)
	// 	fmt.Println()
	// }

	users := make(map[string]map[string]string, 10)

	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "99999"
	users["smith"]["nickName"] = "小花"
	modifyUser(users,"tom")
	modifyUser(users,"mary")
	modifyUser(users,"smith")
	fmt.Println(users)
}

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "88888"
	}else{
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "88888"
		users[name]["nickName"] = "暱稱"+name
	}
}