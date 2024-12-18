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

func GetUserInfo(name string) (User, error) {
	user := User{}
	err := GlobalDB.Where("name=?", name).First(&user).Error
	return user, err
}

func UpdateUserName(oldName, newName string) error {
	// Model  只是声明要操作的表模型，但它不会添加任何 WHERE 查询条件
	// gorm 默认只有主键才可以默认带条件查询
	// GlobalDB.Model(&User{Id:2}).Update("name", newName).Error	// 即：不用写Where了
	return GlobalDB.Model(&User{}).Where("name=?", oldName).Update("name", newName).Error
}

func UserAuthPost(name, realName, idCard string) error {
	// 可以先调外部接口做一个身份信息校验
	// 如果身份信息真实才继续下面的
	return GlobalDB.Model(&User{}).Where("name=?", name).Updates(map[string]interface{}{
		"real_name": realName,
		"id_card":   idCard,
	}).Error
}
