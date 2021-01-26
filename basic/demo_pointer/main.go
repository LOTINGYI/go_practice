package main
import "fmt"

func main()  {
	// var num int = 9

	// fmt.Printf("number address: %v\n",&num)

	// var ptr *int 
	// ptr = &num
	// *ptr = 10
	// fmt.Printf("number: %v",num)
	
	// var a int = 300
	// var ptr *int = a

	// var b int = 300
	// var ptr *float32 = &b 

	var c int =300
	var d int = 400
	var ptr *int = &c
	*ptr = 100
	ptr = &d 
	*ptr = 200
	fmt.Printf("a=%d,b=%d,ptr=%d",c,d,*ptr)
}