package main

import (
	"fmt"
	"go_code/proj/big_chat/common/message"
	"go_code/proj/big_chat/server/processes"
	"go_code/proj/big_chat/server/utils"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) process2() (err error) {
	// 循環讀取客戶端發送訊息
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		// 讀取數據包，直接封裝成readPkg
		mess, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客戶端退出，服務器也退出")
				return err
			} else {
				fmt.Println("readPkg err= ", err)
				return err
			}

		}

		err = this.serverProcessMes(&mess)
		if err != nil {
			return err
		}
	}
}

// 根據客戶端發送消息總類不同，決定調用哪個函數
func (this *Processor) serverProcessMes(mess *message.Message) (err error) {
	switch mess.Type {
	case message.LoginMesType:
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mess)
	case message.RegisterMesType:
		// 處理註冊
		up := &processes.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mess)
	default:
		fmt.Println("消息類型不存在，無法處理")
	}
	return
}
