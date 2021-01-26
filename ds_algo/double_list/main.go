package main


import (
	"fmt"
)

type HeroNode struct{
	no int
	name string
	nickName string
	pre *HeroNode
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
	newHeroNode.pre = temp
}

// 給雙向鏈表中插入節點 
func insertHeroNode2(head *HeroNode, newHeroNode *HeroNode)  {
	temp := head
	flag:= true
	for{
		if temp.next == nil {
			break
		}else if temp.next.no >= newHeroNode.no{
			break
		}else if temp.next.no == newHeroNode.no{
			flag  = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("對不起，已經存在no= ",newHeroNode.no)
	}else{
		newHeroNode.pre = temp
		newHeroNode.next = temp.next
		if temp.next !=nil {
			temp.next.pre = newHeroNode	// 存在一個bug，如果是最後一個節點
		}
		temp.next = newHeroNode
	}
}

// 逆續輸出
func ListHeroNode(head *HeroNode)  {
	temp:= head
	if temp.next == nil {
		fmt.Println("鏈表為空")
		return
	}
	for{
		if temp.next == nil{
			break
		}
		temp = temp.next
	}
	
	for{
		fmt.Printf("[%d,%s,%s]==>",
		temp.no,temp.name,temp.nickName)
		temp = temp.pre
		if temp.pre == nil {
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
		if temp.next != nil{
			temp.next.pre = temp
		}
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

	hero4 :=&HeroNode{
		no:4,
		name:"非凡",
		nickName:"文在否",
	}

	hero3 :=&HeroNode{
		no:3,
		name:"莫凡",
		nickName:"文在銀",
	}

	hero5 :=&HeroNode{
		no:5,
		name:"周元",
		nickName:"元尊",
	}

	insertHeroNode(head,hero1)
	insertHeroNode(head,hero2)
	insertHeroNode(head,hero4)
	insertHeroNode2(head,hero3)
	insertHeroNode(head,hero5)
	DelHeroNode(head,4)
	ListHeroNode(head)

}