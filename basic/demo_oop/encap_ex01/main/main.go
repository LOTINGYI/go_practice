package main

import (
	"fmt"
	"go_code/basic/demo_oop/encap_ex01/model"
)


func main() {
	p:= model.NewAccount("11111cd","999998",40)
	if p != nil {
		fmt.Println("創建成功= ",*p)
	}else{
		fmt.Println("創建失敗")
	}
}