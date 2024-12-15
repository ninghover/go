package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/gomodule/redigo/redis"
)

func UserRegister(mobile, passeord string) error {
	md5 := md5.New()
	md5.Write([]byte(passeord))
	hash := hex.EncodeToString(md5.Sum(nil))
	return GlobalDB.Create(&User{Mobile: mobile, Password_hash: hash, Name: mobile}).Error
}

// SaveImgCode写在getCaptcha服务中的

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

func GetSms(phone string) string {
	conn := RedisPool.Get()
	defer conn.Close()
	reply, _ := redis.String(conn.Do("get", phone+"_ssm_code"))
	return reply
}

func Login(phone, password string) (string, error) {
	md5 := md5.New()
	md5.Write([]byte(password)) // 将字符串转换为字节切片
	hash := hex.EncodeToString(md5.Sum(nil))
	var user User
	err := GlobalDB.Where("mobile=? and password_hash=?", phone, hash).First(&user).Error
	if err != nil {
		return "", err // 没查到
	}
	return user.Name, nil
}
