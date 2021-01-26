package model

import (
	"encoding/json"
	"fmt"
	"go_code/proj/big_chat/common/message"

	"github.com/garyburd/redigo/redis"
)

type UserDao struct {
	pool *redis.Pool
}

// 我們在服務器啟動後，就初始化userDao
// 利用全局變量，在需要和redis操作時，直接使用即可
var (
	MyUserDao *UserDao
)

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGET", "users", id))
	if err != nil {
		if err == redis.ErrNil { // 表示users hash中，沒有找到對應id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json Unmarshal err= ", err)
		return
	}
	return
}

func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PASSWORD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	// 說明還沒有此用戶
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSET", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存用戶失敗 err= ", err)
		return
	}
	return
}
