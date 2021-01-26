package main

import (
	"fmt"
	"reflect"
)




func reflectTest01(b interface{})  {
	
	rType := reflect.TypeOf(b)
	fmt.Println("rType: ",rType)	
	fmt.Println("rType name:",rType.Size())
	rVal := reflect.ValueOf(b)
	n2:=2+rVal.Int()
	fmt.Println("n2= ",n2)
	fmt.Printf("rVal=%T \n",rVal)
	iV := rVal.Interface()
	num2:= iV.(int)
	fmt.Println(num2)
}

type Student struct{
	Name string
	Age int
}

func reflectTest02(b interface{})  {
	rType := reflect.TypeOf(b)
	fmt.Println("rType: ",rType)	
	rVal := reflect.ValueOf(b)
	
	fmt.Printf("kind=%v kind=%v\n",rVal.Kind(),rType.Kind())
	iV := rVal.Interface()

	fmt.Printf("iv=%v iv type=%T\n",iV,iV)
	stu,ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n",stu.Name)
	}
}
func main() {
	
	// var number int = 100
	// reflectTest01(number)

	stu := Student {
		Name:"tim",
		Age: 20,
	}
	reflectTest02(stu)
}