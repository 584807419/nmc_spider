package db

import (
	"fmt"
	"nmc_spider/log_manage"
	"sync/atomic"
)

var logger = log_manage.FSLogger

func GetAllLocation() []Location {
	// sqlStr := "select * from location where id = 1872 and valid = 1"
	sqlStr := "select * from location where valid = 1"
	var location []Location
	// err := DB.Select(&location, sqlStr, 0)
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("数据库调用并发度%d\n", atomic.LoadInt32(&concurrent))
	err := DB.Select(&location, sqlStr)
	atomic.AddInt32(&concurrent, -1)
	if err != nil {
		fmt.Printf("GetAllLocation query failed, err:%v\n", err)
	}
	return location
	// %v 按默认格式输出
	// %+v 在%v的基础上额外输出字段名
	// %#v 在%+v的基础上额外输出类型名
	// fmt.Printf("多行查询:%+v\n", location)
}

func GetAllProvince() []Province {
	logger.Infof("%v", "获取省、自治区、直辖市信息")
	sqlStr := "select * from province where valid = 1"
	var province []Province
	// err := DB.Select(&location, sqlStr, 0)
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("数据库调用并发度%d\n", atomic.LoadInt32(&concurrent))
	err := DB.Select(&province, sqlStr)
	atomic.AddInt32(&concurrent, -1)
	if err != nil {
		fmt.Printf("GetAllProvince query failed, err:%v\n", err)
	}
	return province
}

func InsertRow(sqlStr, uuid string) int64 {
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("数据库调用并发度%d\n", atomic.LoadInt32(&concurrent))
	ret, err := DB.Exec(sqlStr)
	atomic.AddInt32(&concurrent, -1)
	if err != nil {
		logger.Errorf("%v insert failed, err:%v", uuid, err)
		return 0
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		logger.Errorf("%v get lastinsert ID failed, err:%v", uuid, err)
		return 0
	}
	return theID
}

func GetData(sqlStr, uuid string) (EverydayData, error) {
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("数据库调用并发度%d\n", atomic.LoadInt32(&concurrent))
	var everyday_data EverydayData
	err := DB.Get(&everyday_data, sqlStr)
	atomic.AddInt32(&concurrent, -1)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			logger.Infof("%v %v %v", uuid, "没查到数据", sqlStr)
		} else {
			logger.Errorf("%v", uuid, err.Error())
		}

	}
	return everyday_data, err
}

var concurrent int32

func GetRData(sqlStr, uuid string) (EveryTimeData, error) {
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("数据库调用并发度%d\n", atomic.LoadInt32(&concurrent))
	var everyr_data EveryTimeData
	err := DB.Get(&everyr_data, sqlStr)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			logger.Infof("%v %v %v", uuid, "没查到数据", sqlStr)
		} else {
			logger.Errorf("%v", uuid, err.Error())
		}

	}
	atomic.AddInt32(&concurrent, -1)
	return everyr_data, err
}

func GetMultiData(sqlStr, uuid string) Location {
	var location_data Location
	err := DB.Select(&location_data, sqlStr)
	if err != nil {
		logger.Errorf("%v query failed, err:%v", uuid, err)
		logger.Errorf("%v %v", uuid, sqlStr)
	}
	return location_data
}

func ExecSql(sqlStr, uuid string) (int64, error) {
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("数据库调用并发度%d\n", atomic.LoadInt32(&concurrent))
	ret, err := DB.Exec(sqlStr)
	atomic.AddInt32(&concurrent, -1)
	if err != nil {
		logger.Errorf("%v ExecSql:%v err:%v", uuid, sqlStr, err)
		return 0, err
	} else {
		n, _ := ret.RowsAffected() // 操作影响的行数
		if err != nil {
			logger.Errorf("%v get RowsAffected failed, sql:%v err:%v", uuid, sqlStr, err)
		}
		return n, err
	}
}

func GetLocationRec(sqlStr, uuid string) (Location, error) {
	var location_data Location
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("数据库调用并发度%d\n", atomic.LoadInt32(&concurrent))
	err := DB.Get(&location_data, sqlStr)
	atomic.AddInt32(&concurrent, -1)
	return location_data, err
}

func GetLocationData(sqlStr, uuid string) Location {
	var location_data Location
	atomic.AddInt32(&concurrent, 1)
	fmt.Printf("数据库调用并发度%d\n", atomic.LoadInt32(&concurrent))
	_ = DB.Get(&location_data, sqlStr)
	atomic.AddInt32(&concurrent, -1)
	return location_data
}
