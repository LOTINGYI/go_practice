package main

import "fmt"


type Point struct{
	x int
	y int
}

type Rect struct{
	leftUp, rightDown Point
}

func main() {


	r1:=Rect {Point{1,2},Point{3,4}}

	// r1有四個int，在內存中連續分布
	fmt.Printf("r1.leftUp.x 地址=%p\nr1.leftUp.y 地址=%p\nr1.rightDown.x 地址=%p\nr1.rightDown.y 地址=%p\n",
	&r1.leftUp.x,&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)
}