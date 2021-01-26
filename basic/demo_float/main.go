package main

import(
	"fmt"
)
func main() {
	var price float32 = 89.12
	fmt.Println("price: ",price)
	var num1 float32 = -0.0000089
	var num2 float64 = -7809656.09
	fmt.Println("num1 = ",num1,"num2 = ",num2)
	

	// 可能會造成精度損失
	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3 = ",num3,"num4 = ",num4)
}