package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		fmt.Println("服務企在等待客戶端發送訊息 " + conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("readString err=", err)
			return
		}
		fmt.Print(string(buf[:n])) // 如果不戴上n，會把1024位的資料全部打印

	}
}

func main() {
	fmt.Println("Serve listening")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}

	defer listen.Close() // 延時關閉

	for {
		fmt.Println("等待client連接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err= ", err)
		} else {
			fmt.Printf("Accept() success= %v\n", conn)
			// fmt.Printf("Client IP= %T\n",conn.RemoteAddr().String())
			go process(conn)
		}
		// 這裡轉背起個goroutine，為客戶服務
	}

	// fmt.Printf("listen suc=%v\n",listen)
}
