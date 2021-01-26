package main

import (
	"fmt"
	"time"
)
func main() {
	start := time.Now().Unix()
	for i := 1; i <= 200000; i++ {
		// flag := true
		for j := 2; j < i; j++ {

			if i%j ==0 {
				// flag = false
				break
			}
		}
		// if flag {
		// 	fmt.Printf("質數是:%v\n",i)
		// }
	}
	end:=time.Now().Unix()
	fmt.Println("耗時: ",end-start)
	
}