package model

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func SaveImage(code, uuid string) error {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis Dial err: ", err)
		return err
	}
	defer conn.Close()

	_, err = conn.Do("setex", code, 5*60, uuid)
	return err
}
