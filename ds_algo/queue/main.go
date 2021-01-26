package main


import (
	"fmt"
	"errors"
	"os"
)


type Queue struct {
	maxSize int
	array [5]int
	front int
	rear int
}

func (this *Queue) AddQueue(val int) (err error)  {
	if this.rear == this.maxSize-1 {
		return errors.New("queue fail")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

func (this *Queue) ShowQueue()  {
	fmt.Println("對列當前的情況是: ")
	for i := this.front+1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t\n",i,this.array[i])
	}
}

func (this *Queue) GetQueue() (val int,err error) {
	if this.rear == this.front {
		return -1,errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val,err
}


func main()  {
	queue := &Queue{
		maxSize: 5,
		front:-1,
		rear:-1,
	}
	var key string
	var val int
	for{
		fmt.Println("1. 輸入add 表示添加數據")
		fmt.Println("2. 輸入get 表示獲取數據")
		fmt.Println("3. 輸入show 表示顯示數據")
		fmt.Println("4. 輸入exit 表示退出數據")

		fmt.Scanln(&key)
		switch key {
			case "add":
				fmt.Println("輸入你要入對列數")
				fmt.Scanln(&val)
				err := queue.AddQueue(val)
				if err!= nil {
					fmt.Println(err.Error())
				}else{
					fmt.Println("加入對列ok")
				}
			case "get":
				val,err := queue.GetQueue()
				if err != nil {
					fmt.Println(err.Error())
				}else{
					fmt.Println("從對列取出 ",val)
				}
			case "show":
				queue.ShowQueue()
			case "exit":
				os.Exit(0)
		}
	}
}