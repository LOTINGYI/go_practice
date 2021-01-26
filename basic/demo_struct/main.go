package main


import "fmt"

type Cat struct{
	Name string
	Age int
	Color string
	Hobby string
}

func main() {
	
	var cat1 Cat 
	fmt.Printf("cat1的地址是: %p\n",&cat1)
	fmt.Printf("cat1的Name地址是: %p\n",&cat1.Name)
	cat1.Name = "小白"
	cat1.Age = 13
	cat1.Color = "白色"
	cat1.Hobby = "吃魚"
	fmt.Println("cat1 = ",cat1)
}