package main

import (
	"fmt"
	"unsafe"
)
func main() {
	// var i int =1
	// fmt.Println("i: ",i)

	// // test int8
	// var j int8 = -129
	// fmt.Println("j: ",j)


	// // test uint8
	// var k uint8 = -1
	// fmt.Println("k: ",k)

	// var c byte = 256
	// fmt.Println("c: ",c)

	var n1 = 100
	// 如何查看某個變量數據類型
	fmt.Printf("n1的類型是:%T \n",n1)
	// 如何在程序查看某個變量占用字節大小和數據類型
	var n2 int64 = 10
	fmt.Printf("n1的類型是:%T n2占用的字節數:%d",n2,unsafe.Sizeof(n2))

	// 保小不保大原則
	// var age byte = 25

}