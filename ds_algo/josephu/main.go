package main


import (
	"fmt"
)

type Boy struct{
	No int
	Next *Boy
}

/*
	num 小孩個數
	*Boy: 小孩對列的頭指針
*/
func AddBoy(num int) *Boy{
	first := &Boy{}
	cur := &Boy{}
	if num < 1 {
		fmt.Println("num的值不對")
		return first
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			No:i,
		}
		// 構成尋換列表需要一個輔助指針
		// 處理第一個小孩，因為較特殊
		if i==1 {
			first = boy	// 不要亂動
			cur = boy 
			cur.Next = first
		}else{
			cur.Next = boy
			cur = boy 
			cur.Next = first
		}
	}

	return first
}

func ShowBoy(first *Boy)  {
	
	if first.Next == nil {
		fmt.Println("鏈表為空")
		return
	}
	cur := first
	for{
		fmt.Printf("小孩編號=%d ->",cur.No)
		if cur.Next == first {
			break
		}
		cur = cur.Next
	}
}

func PlayGame(first *Boy, startNo int, countNum int)  {
	
	// 空鏈表單獨處理
	if first.Next == nil {
		fmt.Println("空鏈表")
		return
	}
	// 判斷startNo至少要小於小孩總數
	// 輔助指針幫助刪除小孩
	tail := first
	// 讓tail指向隊尾
	for{
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	// 開始數countNum
	for{
		// 開始數
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩編號%d出列->\n",first.No)
		first = first.Next
		tail.Next = first
		if first == tail{
			fmt.Printf("小孩編號%d出列->\n",first.No)
			break
		}
	}
}

func main(){
	first := AddBoy(5)

	ShowBoy(first)
	fmt.Println()
	fmt.Println()
	PlayGame(first,2,3)
}