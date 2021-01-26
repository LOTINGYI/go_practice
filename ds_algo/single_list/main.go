package main


import (
	"fmt"
)

type HeroNode struct{
	no int
	name string
	nickName string
	next *HeroNode
}

func insertHeroNode(head *HeroNode, newHeroNode *HeroNode)  {
	temp := head
	for{
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
}
func ListHeroNode(head *HeroNode)  {
	temp:= head
	if temp.next == nil {
		fmt.Println("鏈表為空")
		return
	}
	for{
		fmt.Printf("[%d,%s,%s]==>",
		temp.next.no,temp.next.name,temp.next.nickName)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func DelHeroNode(head *HeroNode,id int)  {
	temp := head
	flag :=false

	for{
		if temp.next == nil {
			break
		}else if temp.next.no == id{
			flag = true
			break
		}
		temp = temp.next
	}

	if flag {
		temp.next = temp.next.next
	}else{
		fmt.Println("sorry刪除id不存在")
	}
}
func main()  {
	head := &HeroNode{}

	hero1 :=&HeroNode{
		no:1,
		name:"宋江",
		nickName:"及時雨",
	}

	hero2 :=&HeroNode{
		no:2,
		name:"盧俊毅",
		nickName:"預期林",
	}

	hero3 :=&HeroNode{
		no:3,
		name:"非凡",
		nickName:"文在否",
	}

	insertHeroNode(head,hero1)
	insertHeroNode(head,hero2)
	insertHeroNode(head,hero3)
	DelHeroNode(head,0)
	ListHeroNode(head)

}