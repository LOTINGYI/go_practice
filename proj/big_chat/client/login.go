package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/proj/big_chat/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {
	// 連接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial err= ", err)
		return
	}
	defer conn.Close()

	// 準備通過conn發送消息給服務器
	var mess message.Message
	mess.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 將loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal err= ", err)
		return
	}

	mess.Data = string(data)
	// 再將mes序列化
	data, err = json.Marshal(mess)
	if err != nil {
		fmt.Println("json marshal err= ", err)
		return
	}

	// data就是要發送的消息
	// 先把data長度發送給Server
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write 失敗", err)
		return
	}
	// fmt.Printf("客戶端發送消息長度=%d 內容=%s", len(data), string(data))

	// 發送消息本身
	_, err = conn.Write(data)
	if n != 4 || err != nil {
		fmt.Println("conn write 失敗", err)
		return
	}

	// 需要處理服務器端返回處理
	mes, err := readPkg(conn)

	if err != nil {
		fmt.Println("readPkg err= ", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		fmt.Println("登入成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}

	return
}
