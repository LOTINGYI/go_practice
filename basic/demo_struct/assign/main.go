package main

import "fmt"

type Stu struct{
	Name string
	Age int
}

func main() {
	// 方式1
	var stu1 = Stu{"小名",19} // stu1 --> struct數據
	stu2:= Stu{"小黃",19}

	var stu3 = Stu{
		Name:"jacl",
		Age:20,
	}

	fmt.Println(stu1,stu2,stu3)

	// 方式2
	var stu5 = &Stu{"小王",20} // stu5 --> 地址 --> struct數據
	stu6 := &Stu{"小王",20}
	var stu7 = &Stu{
		Name:"小王",
		Age:20,
	}
	fmt.Println(stu5,stu6,stu7)
	fmt.Println(*stu5,*stu6,*stu7)
}