package main


import (
	"fmt"
	_ "time"
	"sync"
	_ "runtime"
)

var (
	result = make(map[int]int,20)
	lock sync.Mutex	// 全局互斥鎖
)

func test(n int)  {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	result[n] = res	// concurrent map writes: race condition
	lock.Unlock()
}
func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	// time.Sleep(time.Second * 10)

	lock.Lock()
	for i,v := range result {
		fmt.Printf("map[%d]=%d\n",i,v)
	}
	lock.Unlock()
}