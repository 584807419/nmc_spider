package db

import (
	"fmt"
	"nmc_spider/log_manage"
)

var logger = log_manage.FSLogger

// 获取数据
func GetAllLocation() []Location {
	sqlStr := "select * from location where valid = 1"
	var location []Location
	// err := DB.Select(&location, sqlStr, 0)
	err := DB.Select(&location, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	return location
	// %v 按默认格式输出
	// %+v 在%v的基础上额外输出字段名
	// %#v 在%+v的基础上额外输出类型名
	// fmt.Printf("多行查询:%+v\n", location)
}

// 插入数据
func InsertRow(sqlStr, uuid string) int64 {
	ret, err := DB.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		logger.Errorf("%v get lastinsert ID failed, err:%v", uuid, err)
		return 0
	}
	return theID
}

// 获取单行数据
func GetData(sqlStr, uuid string) EverydayData {
	var everyday_data EverydayData
	err := DB.Get(&everyday_data, sqlStr)
	if err != nil {
		logger.Errorf("%vquery failed, err:%v", uuid, err)
	}
	return everyday_data

}
