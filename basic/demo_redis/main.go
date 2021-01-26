package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 連接到Redis
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "name", "Tim Lo")
	if err != nil {
		fmt.Println("conn.Do err= ", err)
	}
	r, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("conn.Do err= ", err)
	}
	fmt.Println("r= ", r)
}
