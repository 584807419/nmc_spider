package data_operation

import (
	"nmc_spider/db"
	"nmc_spider/http_requests"
	"nmc_spider/log_manage"
	"nmc_spider/message_queue"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
)

var logger = log_manage.FSLogger
var HttpGet = http_requests.HttpGet

func GetData(wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("error%v", err)
		}
	}()
	for {
		location := db.GetAllLocation()
		for _, value := range location {
			u4 := uuid.New()
			uuidv4 := u4.String()
			logger.Infof("%v %v %v %v", uuidv4, value.Stationid, value.Province, value.City)
			time_stamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
			// 返回json的接口url
			urlHashMap := map[string]string{
				"url":       "http://www.nmc.cn/rest/weather?stationid=" + value.Stationid + "&_=" + time_stamp,
				"uuid":      uuidv4,
				"stationid": value.Stationid,
			}
			// logger.Infof("%v %v", uuidv4, "往无缓冲通道中发送接口url")
			message_queue.TempUrlChan <- urlHashMap
			// resp_data := HttpGet(url, uuidv4)
			// SaveData(resp_data, uuidv4, value.Stationid)
			// time.Sleep(3 * time.Second)

			// 增添html页面url 2023-06-08
			detailUrlHashMap := map[string]string{
				"html_url":  "http://www.nmc.cn" + value.Url,
				"uuid":      uuidv4,
				"stationid": value.Stationid,
			}
			// 无缓冲 没接收的话会阻塞
			// logger.Infof("%v %v", uuidv4, "往无缓冲通道中发送html页面url")
			message_queue.TempUrlChan <- detailUrlHashMap
		}
	}
	// wg.Done()
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
		time.Sleep(5 * time.Second)
	}
	logger.Infof("%v", "表中所有的市区县信息全部处理完成")
}
