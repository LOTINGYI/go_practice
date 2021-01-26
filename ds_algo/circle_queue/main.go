package main


import (
	"fmt"
	"errors"
	"os"
)


type CircleQueue struct{
	maxSize int
	array [5]int
	head int
	tail int
}

func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	// 分析出this.tail在Queue尾部，但是不包含最後元素
	this.array[this.tail] = val
	// this.tail++ // error
	this.tail = (this.tail+1)%this.maxSize
	return
}

func (this *CircleQueue) Pop() (val int,err error) {
	if this.IsEmpty(){
		return 0, errors.New("queue empty")
	}
	val = this.array[this.head]
	// this.head++
	this.head = (this.head+1)%this.maxSize
	return
}

func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func (this *CircleQueue) List()  {
	// if this.IsEmpty() {
	// 	return
	// }

	size := this.Size()
	if size == 0 {
		fmt.Println("queue為空")
	}
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t",tempHead,this.array[tempHead])
		tempHead = (tempHead+1)%this.maxSize
	}
	fmt.Println()
}

func main() {

	queue := &CircleQueue{
		maxSize: 5,
		head: 0,
		tail: 0,
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
				err := queue.Push(val)
				if err!= nil {
					fmt.Println(err.Error())
				}else{
					fmt.Println("加入對列ok")
				}
			case "get":
				val,err := queue.Pop()
				if err != nil {
					fmt.Println(err.Error())
				}else{
					fmt.Println("從對列取出 ",val)
				}
			case "show":
				queue.List()
			case "exit":
				os.Exit(0)
		}
	}
}


