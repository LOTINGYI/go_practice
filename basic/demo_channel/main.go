package main


import (
	"fmt"
	"time"
)


func writeData(intChan chan int)  {
	for i := 1; i <= 50; i++ {
		intChan<-i
		fmt.Printf("writeData 寫入數據=%v\n",i)
		// time.Sleep(time.Second)
	}

	close(intChan)
}

func readData(intChan chan int,exitChan chan bool)  {
	for {
		v,ok := <-intChan
		if !ok {
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 讀到數據=%v\n",v)
	}
	exitChan <- true
	close(exitChan)
}
func main() {
	intChan := make(chan int, 1)
	exitChan:=make(chan bool, 1)
	
	go writeData(intChan)
	go readData(intChan,exitChan)

	// time.Sleep(time.Second)
	for {
		_,ok := <-exitChan
		if !ok {
			break
		}
	}
}