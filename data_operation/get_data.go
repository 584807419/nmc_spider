package data_operation

import (
	"nmc_spider/db"
	"nmc_spider/http_requests"
	"nmc_spider/log_manage"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var logger = log_manage.FSLogger
var HttpGet = http_requests.HttpGet

func GetData() {
	location := db.GetAllLocation()
	for _, value := range location {
		u4 := uuid.New()
		uuidv4 := u4.String()
		logger.Infof("%v %v-%v-%v", uuidv4, value.Stationid, value.Province, value.City)
		time_stamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
		url := "http://www.nmc.cn/rest/weather?stationid=" + value.Stationid + "&_=" + time_stamp
		resp_data := HttpGet(url, uuidv4)
		SaveData(resp_data, uuidv4, value.Stationid)
		time.Sleep(4 * time.Second)
	}
}

func GetProvinceData() {
	province := db.GetAllProvince()
	logger.Infof("%v", "开始逐个处理表中的市区县信息")
	for _, value := range province {
		u4 := uuid.New()
		uuidv4 := u4.String()
		logger.Infof("%v %v-%v", uuidv4, value.Name, value.Abbr)
		time_stamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
		url := "http://www.nmc.cn/rest/province/" + value.Abbr + "?_=" + time_stamp
		logger.Infof("%v %v", uuidv4, "获取市区县信息")
		resp_data := HttpGet(url, uuidv4)
		SaveProvinceCityData(resp_data, uuidv4)
		logger.Infof("%v %v %v", uuidv4, value.Name, "市区县信息处理完成-sleep")
		time.Sleep(5 * time.Second)
	}
	logger.Infof("%v", "表中所有的市区县信息全部处理完成")
}
