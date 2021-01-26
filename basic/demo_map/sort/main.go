package main

import (
	"fmt"
	"sort"
)

func main() {
	
	map1 := make(map[int]int, 10)
	map1[10]= 100
	map1[1] = 13
	map1[4] = 56
	map1[2] = 78
	map1[8] = 90

	// map 本身是無序的
	for k,v := range map1 {
		fmt.Printf("k:%v, v:%v\n",k,v)
	}
	// fmt.Println(map1)

	// 按照map的key順序進行排序輸出
	// 1. 先將map的key放入到切片中
	// 2. 對切片進行排序
	// 3. 遍歷切片，然後按照jey輸出map值

	var keys []int
	for k,_ := range map1 {
		keys = append(keys,k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Printf("map[%v]=%v\n",v,map1[v])
	}
}