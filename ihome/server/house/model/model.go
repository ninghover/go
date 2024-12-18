package model

import (
	"fmt"
	house "house/proto"
	"strconv"
	"time"
)

func GetUserHouse(name string) ([]*house.AHouse, error) { // 切片里面存的是指针 是因为pb需要的是指针
	var user User
	var houses []*House
	// 方式1 	Preload
	// err := GlobalDB.Where("name=?", name).Preload("Houses").First(&user).Error
	// if err != nil {
	// 	fmt.Println("查询失败, ", err)
	// 	return nil, err
	// }
	// houses = user.Houses

	// 方式2 Join   最终从houses表中查，应该Join users表
	e := GlobalDB.Joins("JOIN users ON houses.user_id = users.id").Where("users.name = ?", name).Find(&houses).Error

	if e != nil {
		fmt.Println("查询失败, ", e)
		return nil, e
	}

	var house_res []*house.AHouse
	for _, v := range houses {
		h := house.AHouse{}
		h.Address = v.Address
		loc, _ := time.LoadLocation("Asia/Shanghai")
		h.Ctime = v.CreatedAt.In(loc).Format("2006-01-02 15:04:05")
		h.HouseId = int32(v.ID)
		h.ImgUrl = v.Index_image_url
		h.OrderCount = int32(v.Order_count)
		h.Price = int32(v.Price)
		h.RoomCount = int32(v.Room_count)
		h.Title = v.Title
		h.UserAvatar = user.Avatar_url

		var area Area
		GlobalDB.Where("id=?", v.AreaId).Find(&area)
		h.AreaName = area.Name
		house_res = append(house_res, &h)
	}

	return house_res, nil
}

func AddHouseInfo(req *house.PostHouseReq) (int, error) {
	var house House
	var user User
	if err := GlobalDB.Where("name=?", req.UserName).First(&user).Error; err != nil {
		fmt.Println("查询当前用户失败, ", err)
		return 0, err
	}
	house.UserId = uint(user.ID)
	areaId, _ := strconv.Atoi(req.AreaId)
	house.AreaId = uint(areaId)
	house.Title = req.Title
	house.Address = req.Address
	house.Room_count, _ = strconv.Atoi(req.RoomCount)
	house.Acreage, _ = strconv.Atoi(req.Acreage)
	house.Price, _ = strconv.Atoi(req.Price)
	house.Unit = req.Unit
	house.Capacity, _ = strconv.Atoi(req.Capacity)
	house.Beds = req.Beds
	house.Deposit, _ = strconv.Atoi(req.Deposit)
	house.Min_days, _ = strconv.Atoi(req.MinDays)
	house.Max_days, _ = strconv.Atoi(req.MaxDays)

	var facility []*Facility
	for _, v := range req.Facility {
		id, _ := strconv.Atoi(v)
		fac := Facility{
			Id: id,
		}
		if err := GlobalDB.First(&fac).Error; err != nil {
			fmt.Println("家具查询错误, ", err)
			return 0, err
		}
		facility = append(facility, &fac)
	}
	house.Facilities = facility

	if err := GlobalDB.Create(&house).Error; err != nil {
		fmt.Println("插入房屋信息失败, ", err)
		return 0, err
	}
	return int(house.ID), nil
}
