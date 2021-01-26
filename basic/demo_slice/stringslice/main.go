package main

import "fmt"
func main() {
	str:= "hello@gmail"

	// 使用切片獲取gmail
	slice := str[6:]
	fmt.Println(slice)

	// string是不可變的，會報錯
	// str[0] := 'z'  

	// 所以要先轉換成[]byte()或是[]rune()
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println(str)
	// 細節，轉成[]byte後可以處理英文數字，但不能處理漢字
	arr2 := []rune(str)
	arr2[0] = '北'
	str = string(arr2)
	fmt.Println(str)
}