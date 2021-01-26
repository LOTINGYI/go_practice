package main

import (
	"fmt"
	"go_code/basic/demo_oop/encap/model"
)


func main() {
	p:= model.NewPerson("Smith")
	fmt.Println(*p)
	p.SetAge(20)
	p.SetSal(20000)
	fmt.Println(*p)
}