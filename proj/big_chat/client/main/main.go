package main

import (
	"fmt"
	"go_code/proj/big_chat/client/processes"
	"os"
)

var userId int
var password string
var userName string

func main() {

	// 接收用戶選擇
	var key int
	// 判斷是否繼續使用
	// var loop = true

	for {
		fmt.Println("----------------------歡迎登入多人聊天系統-----------------------")
		fmt.Println("\t\t\t 1 登入聊天")
		fmt.Println("\t\t\t 2 註冊用戶")
		fmt.Println("\t\t\t 3 退出系統")
		fmt.Println("\t\t\t 請選擇(1~3): ")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登入聊天室")
			fmt.Println("請輸入id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("請輸入用戶密碼")
			fmt.Scanf("%s\n", &password)

			up := &processes.UserProcess{}
			up.Login(userId, password)
			// loop = false
		case 2:
			fmt.Println("註冊用戶")
			fmt.Println("請輸入id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("請輸入用戶密碼")
			fmt.Scanf("%s\n", &password)
			fmt.Println("請輸入用戶名字")
			fmt.Scanf("%s\n", &userName)
			up := &processes.UserProcess{}
			up.Register(userId, password, userName)
			// loop = false
		case 3:
			fmt.Println("退出系統")
			os.Exit(0)
		default:
			fmt.Println("你的輸入有誤請重新輸入~")
		}
	}

	// if key == 1 {
	// 	fmt.Println("請輸入id")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Println("請輸入用戶密碼")
	// 	fmt.Scanf("%s\n", &password)
	// 	// 先把登入的函數，寫到另外一個文件
	// 	// 這裡會重新調用
	// 	login(userId, password)
	// 	// if err != nil {
	// 	// 	fmt.Println("登入失敗")
	// 	// } else {
	// 	// 	fmt.Println("登入成功")
	// 	// }
	// } else if key == 2 {
	// 	fmt.Println("進行用戶註冊")
	// }
}
