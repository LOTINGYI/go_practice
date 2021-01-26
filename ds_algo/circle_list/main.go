package main


import (
	"fmt"
)

type CatNode struct{
	no int
	name string
	next *CatNode
}


func InsertCatNode(head *CatNode, newCatNode *CatNode){

	// 假如是第一個節點
	if head.next ==nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println( newCatNode,"加入到環形鏈表")
		return
	}

	temp:=head
	for{
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCircle(head *CatNode){
	fmt.Println("環形鏈表狀況如下")
	temp:=head
	if head.next == nil {
		fmt.Println("此為空鏈表")
		return
	}
	for{
		fmt.Printf("貓的訊息=[id=%d name=%v]->",temp.no,temp.name)
		if temp.next == head {
			break	
		}
		temp = temp.next
	}
}

func DelCatNode(head *CatNode,id int) *CatNode {
	temp := head
	helper := head
	// fmt.Println("world",temp.next,head)
	if temp.next == nil {
		fmt.Println("這是一個空的鏈表")
		return head
	}

	// 如果只有一個節點
	if temp.next == head {
		// fmt.Println("hello")
		temp.next = nil
		return head
	}

	// 將helper定位到鏈表最後
	for{
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	for{
		// 說明比較到最後一個[但尚未比較]
		if temp.next == head {	
			break
		}
		if temp.no == id {
			if temp == head {
				fmt.Println("主要處理刪除第一個")
				head = head.next
			}
			helper.next =temp.next
			fmt.Printf("貓貓=%d被幹掉\n",id)
			flag = false
			break
		}

		temp = temp.next	// 用來比較
		helper = helper.next	// 一旦找到要刪除節點
	}

	// 用來比較最後一個
	if flag {
		fmt.Println("主要處理刪除最後一個")
		if temp.no == id {
			helper.next =temp.next
			fmt.Printf("貓貓=%d被幹掉\n",id)
		}else{
			fmt.Println("沒有此貓")
		}
	}

	return head
}

func main(){
	head := &CatNode{}

	cat1 := &CatNode{
		no:1,
		name:"tom",
	}
	cat2 := &CatNode{
		no:2,
		name:"martin",
	}
	cat3 := &CatNode{
		no:3,
		name:"alex",
	}
	InsertCatNode(head,cat1)
	InsertCatNode(head,cat2)
	InsertCatNode(head,cat3)
	ListCircle(head)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	head = DelCatNode(head,2)
	ListCircle(head)
}