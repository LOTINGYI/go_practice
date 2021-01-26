package main

import("fmt")

var (
	// 此為全局變量
	Fun1 = func(n1 int,n2 int)int{
		return n1 * n2
	}
)

func main() {
	// 在定義匿名函數時直接調用，這種方式匿名函數只能調用一次

	res := func (n1 int,n2 int) int {
		return n1+n2	
	}(10,20)
	fmt.Println(res)

	// 將匿名函數 func (n1 int,n2 int) int賦給a變量
	// 則a的數據類型就是函數類型
	a := func(n1 int, n2 int) int {
		return n1-n2	
	}
	res2 := a(10,30)
	fmt.Println(res2)
	res3 := a(90,30)
	fmt.Println(res3)

	res4 := Fun1(4,9)
	fmt.Println(res4)
}