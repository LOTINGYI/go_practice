package main

import "fmt"

func main() {

	var a int =1 >>2
	var b int = -1 >>2
	var c int = 1<<2
	var d int = -1<<2
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)	

	fmt.Println(2&3)
	fmt.Println(2|3)
	fmt.Println(13&7)
	fmt.Println(5|4)
	fmt.Println(-2^2)	

}