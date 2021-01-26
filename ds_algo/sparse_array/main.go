package main


import "fmt"

type Node struct {
	row int
	col int
	val int
}
func main() {
	
	var chessMap [11][11]int 
	/* 1表示黑棋 2表示藍棋*/
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for _,v := range chessMap {
		for _,v2 := range v {
			fmt.Printf("%d ",v2)
		}
		fmt.Println()
	}

	// 轉成稀疏數組
	var sparseArr []Node

	node := Node{
		row:11,
		col:11,
		val:0,
	}
	sparseArr = append(sparseArr,node)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				node := Node{
					row:i,
					col:j,
					val:v2,
				}
				sparseArr = append(sparseArr,node)
			}
			
		}
	}

	fmt.Println()
	fmt.Println("Sparse Array")
	for i,node := range sparseArr {
		fmt.Printf("%d: %d %d %d\n",i,node.row,node.col,node.val)
	}

	fmt.Println()

	var chessMap2 [11][11]int

	for i, node := range sparseArr {
		if i!=0 {
			chessMap2[node.row][node.col] = node.val
		}
	}

	// 恢復數組
	for _,v := range chessMap2 {
		for _,v2 := range v {
			fmt.Printf("%d ",v2)
		}
		fmt.Println()
	}

}