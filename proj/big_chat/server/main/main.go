package main

import (
	"fmt"
	"go_code/proj/big_chat/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {
	// 這裡需要延時關閉
	defer conn.Close()

	p := &Processor{
		Conn: conn,
	}
	err := p.process2()
	if err != nil {
		fmt.Println("客戶端服務器端協程錯誤=err", err)
		return
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func init() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func main() {

	fmt.Println("新架構服務器在8889監聽...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net Listen err= ", err)
		return
	}
	for {
		fmt.Println("等待客戶端連接服務器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen Accept err= ", err)
		}

		// 一但連接成功，則成功啟動一個協程和客戶端保持通訊
		go process(conn)
	}
}
