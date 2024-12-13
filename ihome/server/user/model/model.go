package model

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisPool redis.Pool

func InitRedisPool() {
	RedisPool = redis.Pool{
		MaxIdle:     10,
		MaxActive:   50,
		IdleTimeout: 5 * 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func GetImgCode(uuid string) string {
	conn := RedisPool.Get()
	defer conn.Close()
	reply, _ := redis.String(conn.Do("get", uuid))
	return reply
}

func SaveSms(phone, code string) error {
	conn := RedisPool.Get()
	defer conn.Close()
	_, err := conn.Do("setex", phone+"_ssm_code", 5*60, code)
	return err
}
