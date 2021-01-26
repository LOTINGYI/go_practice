package main

import "fmt"

func main() {
	var slice []float64 = make([]float64, 5,10)

	slice[0] = 50
	slice[3] = 87
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	var slice2 []int = []int {1,3,5}
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}