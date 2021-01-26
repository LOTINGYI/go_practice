package main

import "fmt"

func test01(arr [3]int)  {
	arr[0] = 88
}

func test02(arr *[3]int)  {
	(*arr)[0] = 88
}
func main() {
	// 定義
	var hens [6]float64
	// 賦值
	hens[0] = 3.0
	hens[1] = 5.0
	hens[2] = 1.0
	hens[3] = 3.4
	hens[4] = 2.0
	hens[5] = 50.0

	totalweight :=0.0
	for i:=0; i<len(hens);i++{
		totalweight += hens[i]
	}

	avgWeight := totalweight/float64(len(hens))
	fmt.Printf("total weight: %v, average weight: %v\n",totalweight,avgWeight)

	var intArr [3]int	// int占8字節
	fmt.Printf("intArr的地址%p intArr[0]的地址%p intArr[1]的地址%p intArr[2]的地址%p\n",
	&intArr,&intArr[0],&intArr[1],&intArr[2])

	var numArr01 [3]int = [3]int {1,2,3}
	var numArr02 = [3] int{1,2,3}
	var numArr03 = [...]int{1,2,3,4,5,6}
	var numArr04 = [...]int{1:800,0:900,2:999}
	fmt.Println(numArr01,numArr02,numArr03,numArr04)
	fmt.Println()
	heroes :=[...]string {"timlo","martin","alex"}
	for i,v := range heroes{
		fmt.Printf("i=%v v=%v\n",i,v)
	}

	arr:= [3]int{11,22,33}
	test01(arr)
	fmt.Println(arr)
	test02(&arr)
	fmt.Println(arr)
}