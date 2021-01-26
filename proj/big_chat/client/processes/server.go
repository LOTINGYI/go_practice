package processes

import (
	"fmt"
	"os"
)

func ShowMenu() {
	fmt.Println("----------------------恭喜***登入成功-----------------------")
	fmt.Println("\t\t\t 1 在線用戶列表")
	fmt.Println("\t\t\t 2 發送消息")
	fmt.Println("\t\t\t 3 信息列表")
	fmt.Println("\t\t\t 4 退出列表")
	fmt.Println("\t\t\t 請選擇(1~4): ")

	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("顯示在線用戶列表")
	case 2:
		fmt.Println("發送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系統")
		os.Exit(0)
	default:
		fmt.Println("您輸入的選項不正確")
	}
}
