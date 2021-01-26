package message

type User struct {
	// 為了序列化和反序列化成功，我們必須保證
	// 用戶信息的json字串和key和struct自斷對應的tag必須一致
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
