package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/proj/big_chat/common/message"
	"net"
)

func readPkg(conn net.Conn) (mess message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("讀取客戶端發送數據")
	n, err := conn.Read(buf[:4])
	if n != 4 || err != nil {
		// fmt.Println("conn Read err= ", err)
		// err = errors.New("read pkg header error")
		return
	}
	fmt.Println("總長: ", len(buf))
	fmt.Println("讀到buf= ", buf[:4])
	// 根據buf[:4]轉成一個uint32類型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	// 根據pkgLen讀取消息內容
	n, err = conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// fmt.Println("conn Read err= ", err)
		// err = errors.New("read pkg body error")
		return
	}

	// 把pkgLen反序列化成->message.Message
	err = json.Unmarshal(buf[:pkgLen], &mess)
	if err != nil {
		fmt.Println("json Unmarshal err= ", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write 失敗", err)
		return
	}

	// 發送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write 失敗", err)
		return
	}
	return
}
