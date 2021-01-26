package model

import "errors"

var (
	ERROR_USER_NOTEXISTS = errors.New("用戶不存在")
	ERROR_USER_EXISTS    = errors.New("用戶已存在")
	ERROR_USER_PASSWORD  = errors.New("密碼不正確")
)
