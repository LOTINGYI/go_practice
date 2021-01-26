package main


import (
	"fmt"
	"time"
)


func sayHello()  {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}

func test()  {
	// 可以使用錯誤處理機制 defer + recover
	defer func() {
		// 捕獲test拋出的panic
		if err:=recover();err!= nil {
			fmt.Println("test() 發生異常")
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"	// error
}
func main() {
	
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok= ",i)
	}
}