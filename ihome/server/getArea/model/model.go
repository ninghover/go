package model

import (
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func GetArea() ([]Area, error) {
	var areas []Area
	conn := RedisPool.Get()
	defer conn.Close()
	areaByte, _ := redis.Bytes(conn.Do("get", "area_info"))
	if len(areaByte) == 0 {
		if err := GlobalDB.Find(&areas).Error; err != nil {
			fmt.Println(areas, err)
			return areas, nil
		}
		areaJson, _ := json.Marshal(areas)
		conn.Do("setex", "area_info", 5*60, areaJson)
	} else {
		json.Unmarshal(areaByte, &areas)
	}
	return areas, nil
}
