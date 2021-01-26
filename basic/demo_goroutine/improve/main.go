package main


import (
	"fmt"
	"time"
)

func putNum(intChan chan int)  {
	for i := 1; i <= 200000; i++ {
		intChan<-i
		// fmt.Printf("寫入數據%v\n",i)
	}

	close(intChan)
}

func findPrime(intChan chan int,resChan chan int,exitChan chan bool)  {

	flag := true
	for{
		v,ok := <-intChan
		if !ok {
			break
		}

		// fmt.Printf("讀取數據%v\n",v)
		for i := 2; i < v; i++ {
			if v%i ==0 {
				flag = false
				break
			}
		}
		if flag {
			resChan <- v
		}
	}
	
	exitChan <- true
}

func main() {
	intChan := make(chan int, 200000)
	resChan := make(chan int, 300000)
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()
	go putNum(intChan)
	for i := 0; i < 4; i++ {
		go findPrime(intChan,resChan,exitChan)
	}

	// time.Sleep(time.Second)

	go func() {
		for{
			for i := 0; i < 4; i++ {
				<-exitChan
			}
			end:=time.Now().Unix()
			fmt.Println("耗時: ",end-start)
			close(resChan)
		}
	}()
	// findPrime(intChan,resChan,exitChan)

	for {
		_, ok:= <-resChan
		if !ok {
			break
		}
		// fmt.Printf("質數是:%v\n",res)

	}
	fmt.Printf("main線程退出")
}