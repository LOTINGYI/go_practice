package main

import (
	"fmt"
	util "go_code/basic/demo_package/utils"
)

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte='+'

	result := util.Cal(n1,n2,operator)
	fmt.Println("result: ",result)
}