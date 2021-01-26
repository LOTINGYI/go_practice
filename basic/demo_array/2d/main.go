package main

import "fmt"



func main() {
	
	/*
	0 0 0 0 0 0
	0 0 1 0 0 0
	0 2 0 3 0 0
	0 0 0 0 0 0 
	*/

	// var arr [4][6]int
	// arr[1][2] = 1
	// arr[2][1] = 2
	// arr[2][3] = 3
	// fmt.Println(arr)

	// fmt.Printf("arr[0]的地址是%p\n", &arr[0])
	// fmt.Printf("arr[1]的地址是%p\n", &arr[1])
	// fmt.Printf("arr[2]的地址是%p\n", &arr[2])

	// fmt.Printf("arr[0][0]的地址是%p\n", &arr[0][0])
	// fmt.Printf("arr[1][0]的地址是%p\n", &arr[1][0])

	// var arr2 [2][3]int = [2][3]int{{1,2,3},{4,5,6}}
	// fmt.Println("arr3= ",arr2)

	var scores [3][5] float64

	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("請輸入第%d班第%d個學生的成績:\n",i+1,j+1)
			fmt.Scanln(&scores[i][j])
		}
	}
	fmt.Println(scores)

	totalSum :=0.0
	for i := 0; i < len(scores); i++ {
		sum:=0.0
		for j := 0; j < len(scores[i]); j++ {
			sum+= scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班的總分，平均分%v\n",i+1,sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班級總分:%v",totalSum)
}