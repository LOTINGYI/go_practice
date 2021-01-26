package main

import "fmt"

func main() {
	var intArr [5]int = [...]int{1,22,33,66,99}

	slice:=intArr[1:3]
	fmt.Println("intArr= ",slice)
	fmt.Println("slice 元素是= ",slice)
	fmt.Println("slice 元素個數是= ",len(slice))
	fmt.Println("slice 的容量是= ",cap(slice))

	fmt.Printf("intArr的地址是=%p\n", &intArr)
	fmt.Printf("intArr[1]的地址是=%p\n", &intArr[1])
	fmt.Printf("slice[0]的地址是=%p slice[0]=%v\n",&slice[0],slice[0])
	fmt.Printf("slice的地址是=%p",&slice)

	fmt.Println("修改之前")
	fmt.Println("intArr= ",intArr)
	fmt.Println("slice= ",slice)
	slice[1] = 34
	fmt.Println("修改之後")
	fmt.Println("intArr= ",intArr)
	fmt.Println("slice= ",slice)
}