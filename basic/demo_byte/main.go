package main

import(
	"fmt"
)
func main() {
	var c1 byte = 'a'
	var c2 byte = 'b'
	var c3 byte = '0'

	// 當我們輸出byte值就是輸出了對應的字符ASCII碼值
	fmt.Println("c1 = ", c1)
	fmt.Println("c2 = ", c2)
	fmt.Println("c3 = ", c3)

	fmt.Printf("c1 = %c \nc2 = %c\n",c1,c2)

	// var c4 byte = '北' // overflow
	var c4 int = '北'
	fmt.Println("c4 = ",c4)
	fmt.Printf("c4 = %c \n",c4)

	// 字符類型是可以進行運算的，相當於一個整數
	var n1 = 10 + 'a'
	fmt.Println("n1 = ",n1)	// 10 + 97
	fmt.Printf("n1 = %c \n",n1) // k
}