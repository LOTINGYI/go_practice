package processes

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/proj/big_chat/client/utils"
	"go_code/proj/big_chat/common/message"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	// 連接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial err= ", err)
		return
	}
	defer conn.Close()

	// 準備通過conn發送消息給服務器
	var mess message.Message
	mess.Type = message.RegisterMesType
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId // 因為我們是直接賦User struct給RegisterMes struct，所以需要先registerMes.User
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	// 將loginMes序列化
	data, err := json.Marshal(registerMes)
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

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("註冊發送訊息出錯 err= ", err)
		return
	}

	// 需要處理服務器端返回處理
	mes, err := tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg err= ", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)

	if registerResMes.Code == 200 {
		fmt.Println("註冊成功，請重新登入")
		// os.Exit(0)

	} else {
		fmt.Println(registerResMes.Error)
		// os.Exit(0)
	}

	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
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
	fmt.Printf("客戶端發送消息長度=%d 內容=%s", len(data), string(data))

	// 發送消息本身
	_, err = conn.Write(data)
	if n != 4 || err != nil {
		fmt.Println("conn write 失敗", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	// 需要處理服務器端返回處理
	mes, err := tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg err= ", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {
		fmt.Println("登入成功")

		/*
			這裡還需要在客戶端起一個協程
			該協程保持和服務器端的通訊，
			如果服務器有數據推送給客戶端
			則接收並顯示在客戶端
		*/
		go ProcessServerMes(conn)

		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

// 和服務器端保持通訊
func ProcessServerMes(conn net.Conn) {
	// 創建一個transfer實例
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客戶端正在等待不停讀取")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf Reading 失敗 err=", err)
		}
		fmt.Printf("mess=%v", mes)
	}
}
