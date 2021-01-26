package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/proj/big_chat/common/message"
	"net"
)

// 將這些方法關連到struct中
type Transfer struct {
	// 分析應該有哪些field
	Conn net.Conn
	Buf  [8096]byte // 使用緩衝

}

func (this *Transfer) ReadPkg() (mess message.Message, err error) {
	// buf := make([]byte, 8096)
	fmt.Println("讀取客戶端發送數據")
	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		// fmt.Println("conn Read err= ", err)
		// err = errors.New("read pkg header error")
		return
	}
	fmt.Println("總長: ", len(this.Buf))
	fmt.Println("讀到buf= ", this.Buf[:4])
	// 根據buf[:4]轉成一個uint32類型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	// 根據pkgLen讀取消息內容
	n, err = this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		// fmt.Println("conn Read err= ", err)
		// err = errors.New("read pkg body error")
		return
	}

	// 把pkgLen反序列化成->message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mess)
	if err != nil {
		fmt.Println("json Unmarshal err= ", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write 失敗", err)
		return
	}

	// 發送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn write 失敗", err)
		return
	}
	return
}
