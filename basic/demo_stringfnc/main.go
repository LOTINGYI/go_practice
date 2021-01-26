package main

import (
	"fmt"
	"strconv"
	"strings"
)


func main() {
	str :="hello北京"

	fmt.Println("str len= ",len(str))

	// 字符串遍歷，處理中文問題
	str2 := []rune(str)
	for i:=0 ;i<len(str2);i++{
		fmt.Printf("字符=%c\n",str2[i])
	}

	// 字符串轉整數 
	n,err := strconv.Atoi("123")
	if err!=nil{
		fmt.Println("轉換錯誤",err)
	}else{
		fmt.Println("轉換結果是: ",n)
	}

	// 整數轉字符串
	str = strconv.Itoa(123456)
	fmt.Println("轉換結果是: ",str)

	// 字符串轉[]byte
	var bytes= []byte("hello go")
	fmt.Println("轉換結果是: ",bytes)

	// []byte轉字符串
	str = string([]byte{97,98,99})
	fmt.Println("轉換結果是: ",str)

	// 10進制轉2、8、16進制: str = strconv.FormatInt(123,2)返回對應字符串
	str = strconv.FormatInt(123,2)
	fmt.Println("轉換結果是: ",str)

	b:=strings.Contains("seafood","fcdsa")
	fmt.Println("轉換結果是: ",b)

	num:=strings.Count("cheese","esee")
	fmt.Println("轉換結果是: ",num)
}