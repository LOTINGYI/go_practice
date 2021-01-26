package processes

import (
	"encoding/json"
	"fmt"
	"go_code/proj/big_chat/common/message"
	"go_code/proj/big_chat/server/model"
	"go_code/proj/big_chat/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json Unmarshal 失敗 err= ", err)
		return
	}
	// 先聲明resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "註冊發生未知錯誤"
		}
	} else {
		registerResMes.Code = 200
		// fmt.Println(user, "登入成功")
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json Marshal 失敗 err= ", err)
		return
	}

	// 將data賦值給resMes
	resMes.Data = string(data)

	// 準備發送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json Marshal 失敗 err= ", err)
		return
	}

	// 發送，將其封裝到writePkg
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

// 編寫ServerProcessMess
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	// 核心
	// 從mess取出mess.Data並直接反序列化LoginMes
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json Unmarshal 失敗 err= ", err)
		return
	}

	// 先聲明resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PASSWORD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服務器內部錯誤"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登入成功")
	}
	// if (loginMes.UserId == 100) && (loginMes.UserPwd == "123456") {
	// 	// 合法
	// 	loginResMes.Code = 200

	// } else {
	// 	// 不合法
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "該用戶不存在，請註冊在使用"
	// }

	// 將loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json Marshal 失敗 err= ", err)
		return
	}

	// 將data賦值給resMes
	resMes.Data = string(data)

	// 準備發送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json Marshal 失敗 err= ", err)
		return
	}

	// 發送，將其封裝到writePkg
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
