package db

import "fmt"

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
func InsertRow(sqlStr string) {
	ret, err := DB.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 获取单行数据
func GetData(sqlStr string) EverydayData {
	var everyday_data EverydayData
	err := DB.Get(&everyday_data, sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
	}
	return everyday_data

}
