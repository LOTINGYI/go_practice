package main

import "fmt"


func sum(n1 int,n2 int) int {
	// 當執行defer時，暫時會將defer後面語句壓入到棧中
	// 當函數執行完畢(return)，再從defer依照LIFO方式出棧
	defer fmt.Println("ok1 n1= ",n1)	// 10
	defer fmt.Println("ok2 n2= ",n2)	// 20
	n1++	// 11
	n2++	// 21
	res:=n1+n2	// 32
	fmt.Println("ok3 res= ",res)
	return res
}
func main() {
	res := sum(10,20)
	fmt.Println("res= ",res)	// 32
}