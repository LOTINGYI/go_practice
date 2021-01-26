package main

import (
	"fmt"
)


func main() {
	// var arr [5]int = [...]int {10,20,30,40,50}

	// slice := arr[1:4]
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Printf("slice[%v]=%v\n",i,slice[i])
	// }

	// fmt.Println()
	// for i,v := range slice {
	// 	fmt.Printf("i=%v v=%v\n",i,v)
	// }
	// fmt.Println(slice)

	// slice2 := slice[1:2]
	// fmt.Println(slice2)

	// slice2[0] = 100
	// fmt.Println(slice2)
	// fmt.Println(slice)

	// var slice3 []int = []int {100,200,300}
	// // 通過append直接給slice3追加具體元素
	// slice3 = append(slice3,400,500,600,700)
	// fmt.Println("slice3",slice3)
	// // 通過append將切片slice3追加slice3
	// slice4 := append(slice3,slice3...)
	// fmt.Println("slice4",slice4)

	// fmt.Println(cap(slice4))

	var slice5 []int = []int{1,2,3,4,5}
	var slice6 = make([]int,10)
	copy(slice6,slice5)
	fmt.Println("slice5= ",slice5)
	fmt.Println("slice6= ",slice6)
}