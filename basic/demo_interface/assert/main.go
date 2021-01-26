package main

import "fmt"

type Point struct{
	x int 
	y int
}


func main() {
	// var a interface{}
	// var point Point = Point{1,2}
	// a = point
	// var b Point
	// // b = a 	// =>error
	
	// b= a.(Point)	// 類型斷言
	// fmt.Println(b)
	var x interface{}
	var b2 float32 = 2.1
	x = b2

	if y,ok:= x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y的類型=%T值是=%v\n",y,y)
	}else{
		fmt.Println("convert fail")
	}

	fmt.Println("keep going")
}