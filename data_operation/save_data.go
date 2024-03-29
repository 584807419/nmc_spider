package data_operation

import (
	"encoding/json"
	"fmt"
	"nmc_spider/db"
	"strconv"
	"strings"
	"time"
)

func saveRtableData(respData map[string]interface{}, uuid, stationid string) {
	// logger.Debugf("%v-%v", uuid, "实时")
	realData, ok := respData["real"].(map[string]interface{})
	if ok {
		realWeatherPublishTime := realData["publish_time"].(string)
		temp_t, _ := time.ParseInLocation("2006-01-02 15:04", realWeatherPublishTime, time.Local)
		temp_t_date := temp_t.Format("2006-01-02")
		temp_t_time := temp_t.Format("15:04")
		realWeatherData := realData["weather"].(map[string]interface{})

		yearStr := strconv.FormatInt(int64(time.Now().Year()), 10)
		rtableName := stationid + "r" + "_" + yearStr

		//  type-switch 断言来判断某个 interface 变量中实际存储的变量interface{}类型
		var temperature any
		switch realWeatherData["temperature"].(type) {
		case float64:
			temperature = fmt.Sprintf("%.1f", realWeatherData["temperature"].(float64))
		case string:
			temperature = realWeatherData["temperature"].(string)
		}

		// 湿度
		var humidity any
		switch realWeatherData["humidity"].(type) {
		case float64:
			humidity = fmt.Sprintf("%.1f", realWeatherData["humidity"].(float64))
		case string:
			humidity = realWeatherData["humidity"].(string)
		}

		var rain any
		switch realWeatherData["rain"].(type) {
		case float64:
			rain = fmt.Sprintf("%.1f", realWeatherData["rain"].(float64))
		case string:
			rain = realWeatherData["rain"].(string)
		}
		realWeatherDataIcomfortValue, ok := realWeatherData["icomfort"]
		var icomfort string
		if ok {
			icomfort = fmt.Sprintf("%.0f", realWeatherDataIcomfortValue.(float64))
		}

		info := realWeatherData["info"].(string)
		var feelst string
		realWeatherDataFeelstValue, ok := realWeatherData["feelst"]
		if ok {
			feelst = fmt.Sprintf("%.1f", realWeatherDataFeelstValue.(float64))
		}

		realWindData := realData["wind"].(map[string]interface{})
		wind_direct := realWindData["direct"].(string)
		if wind_direct == "9999" {
			wind_direct = "无"
		}
		wind_power := ""
		windPowerValue, ok := realWindData["power"]
		if ok {
			wind_power = windPowerValue.(string)
		}
		var wind_speed any
		switch realWindData["speed"].(type) {
		case float64:
			wind_speed = fmt.Sprintf("%.1f", realWindData["speed"].(float64))
		case string:
			wind_speed = realWindData["speed"].(string)
		}

		warn_str := ""
		realDataValue, ok := realData["warn"]
		if ok {
			realWarnData := realDataValue.(map[string]interface{})
			signaltype := realWarnData["signaltype"].(string)
			signallevel := realWarnData["signallevel"].(string)
			issuecontent := realWarnData["issuecontent"].(string)
			warn_str = signaltype + ":" + signallevel + "\n" + issuecontent
			if signaltype == "9999" {
				warn_str = ""
			}
		}

		aqi := ""
		aq := ""
		airData, ok := respData["air"].(map[string]interface{})
		if ok {
			aqifl64, ok := airData["aqi"].(float64)
			if ok {
				aqi = fmt.Sprintf("%.0f", aqifl64)
			}
			aqfl64, ok := airData["aq"].(float64)
			if ok {
				aq = fmt.Sprintf("%.0f", aqfl64)
			}
			if aq == "9999" {
				aq = ""
			}
		}

		rtableNameSqlStr := fmt.Sprintf("insert into %v (date, time, temperature,humidity,rain,icomfort,info,feelst,wind_direct,wind_power,wind_speed,warn,aqi,aq) values ('%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v')", rtableName, temp_t_date, temp_t_time, temperature, humidity, rain, icomfort, info, feelst, wind_direct, wind_power, wind_speed, warn_str, aqi, aq)
		if !strings.Contains(rtableNameSqlStr, "9999") {
			getOneData := fmt.Sprintf("select * from %v where date = '%v' and time = '%v' order by id desc limit 1", rtableName, temp_t_date, temp_t_time)
			everyr_data, err := db.GetRData(getOneData, uuid)
			if err != nil {
				_pk := db.InsertRow(rtableNameSqlStr, uuid)
				logger.Infof("%v %v %v%v", uuid, rtableName, "插入实时数据 pk:", _pk)
			} else {
				if everyr_data.Temperature == temperature {
					logger.Infof("%v %v %v %v", uuid, everyr_data.Date, everyr_data.Time, "无新数据")
				} else {
					logger.Infof("%v %v %v %v", uuid, everyr_data.Date, everyr_data.Time, "有新数据")
					// updatesql := fmt.Sprintf("UPDATE %v SET `date` = '%v', `day_info` = '%v', `day_temperature` = '%v', `day_direct` = '%v', `day_power` = '%v', `night_info` = '%v', `night_temperature` = '%v', `night_direct` = '%v', `night_power` = '%v' WHERE `id` = %v;", tableName, temp_t_date, dayInfo_weather_info, dayInfo_weather_temperature, dayInfo_wind_direct, dayInfo_wind_power, nightInfo_weather_info, nightInfo_weather_temperature, nightInfo_wind_direct, nightInfo_wind_power, everyday_data.Id)
					updatesql := fmt.Sprintf("update %v set date='%v', time='%v', temperature='%v',humidity='%v',rain='%v',icomfort='%v',info='%v',feelst='%v',wind_direct='%v',wind_power='%v',wind_speed='%v',warn='%v',aqi='%v',aq='%v' where id = '%v'", rtableName, temp_t_date, temp_t_time, temperature, humidity, rain, icomfort, info, feelst, wind_direct, wind_power, wind_speed, warn_str, aqi, aq, everyr_data.Id)
					rowCount, err := db.ExecSql(updatesql, uuid)
					if err != nil {
						logger.Errorf("%v %v %v %v", uuid, "更新失败", updatesql, err)
					} else {
						logger.Infof("%v %v %v %v", uuid, rtableName, "更新成功,row:", rowCount)
					}
				}
			}

		} else {
			logger.Debugf("%v-%v", uuid, "没插入-发现9999")
		}
	} else {
		logger.Debugf("%v-%v", uuid, "无实时数据")
	}

}

func savetableData(respData map[string]interface{}, uuid, stationid string) {
	// logger.Debugf("%v-%v", uuid, "预报")
	yearStr := strconv.FormatInt(int64(time.Now().Year()), 10)
	tableName := stationid + "_" + yearStr
	respDataValue := respData["predict"]
	switch respDataValue.(type) {
	case map[string]interface{}:
		predictData := respDataValue.(map[string]interface{})
		detailSlice := predictData["detail"].([]interface{})
		for _, HMapItem := range detailSlice {
			HMapDict := HMapItem.(map[string]interface{})
			dayInfo := HMapDict["day"].(map[string]interface{})
			temp_t_date := HMapDict["date"].(string)
			dayInfo_weather := dayInfo["weather"].(map[string]interface{})
			dayInfo_weather_info := dayInfo_weather["info"]
			dayInfo_weather_temperature := dayInfo_weather["temperature"]
			dayInfo_wind := dayInfo["wind"].(map[string]interface{})
			dayInfo_wind_direct := dayInfo_wind["direct"]
			dayInfo_wind_power := dayInfo_wind["power"]

			nightInfo := HMapDict["night"].(map[string]interface{})
			nightInfo_weather := nightInfo["weather"].(map[string]interface{})
			nightInfo_weather_info := nightInfo_weather["info"]
			nightInfo_weather_temperature := nightInfo_weather["temperature"]
			nightInfo_wind := nightInfo["wind"].(map[string]interface{})
			nightInfo_wind_direct := nightInfo_wind["direct"]
			nightInfo_wind_power := nightInfo_wind["power"]

			tableNameSqlStr := fmt.Sprintf("insert into %v (date, day_info,day_temperature,day_direct,day_power,night_info,night_temperature,night_direct,night_power) values ('%v','%v','%v','%v','%v','%v','%v','%v','%v')", tableName, temp_t_date, dayInfo_weather_info, dayInfo_weather_temperature, dayInfo_wind_direct, dayInfo_wind_power, nightInfo_weather_info, nightInfo_weather_temperature, nightInfo_wind_direct, nightInfo_wind_power)
			if !strings.Contains(tableNameSqlStr, "9999") {
				getOneData := fmt.Sprintf("select * from %v where date = '%v' order by id desc limit 1", tableName, temp_t_date)
				everyday_data, err := db.GetData(getOneData, uuid)
				if err != nil {
					_pk := db.InsertRow(tableNameSqlStr, uuid)
					logger.Infof("%v %v-%v%v", uuid, tableName, "插入成功 pk:", _pk)
				} else {
					if (everyday_data.Day_info == dayInfo_weather_info) && (everyday_data.Day_temperature == dayInfo_weather_temperature) && (everyday_data.Day_direct == dayInfo_wind_direct) && (everyday_data.Day_power == dayInfo_wind_power) && (everyday_data.Night_info == nightInfo_weather_info) && (everyday_data.Night_temperature == nightInfo_weather_temperature) && (everyday_data.Night_direct == nightInfo_wind_direct) && (everyday_data.Night_power == nightInfo_wind_power) {
						// logger.Debugf("%v %v %v", uuid, everyday_data.Date, "无新数据")
					} else if (everyday_data.Day_info != dayInfo_weather_info) || (everyday_data.Day_temperature != dayInfo_weather_temperature) || (everyday_data.Day_direct != dayInfo_wind_direct) || (everyday_data.Day_power != dayInfo_wind_power) || (everyday_data.Night_info != nightInfo_weather_info) || (everyday_data.Night_temperature != nightInfo_weather_temperature) || (everyday_data.Night_direct != nightInfo_wind_direct) || (everyday_data.Night_power != nightInfo_wind_power) {
						// logger.Debugf("%v %v %v", uuid, everyday_data.Date, "有新数据")
						updatesql := fmt.Sprintf("UPDATE %v SET `date` = '%v', `day_info` = '%v', `day_temperature` = '%v', `day_direct` = '%v', `day_power` = '%v', `night_info` = '%v', `night_temperature` = '%v', `night_direct` = '%v', `night_power` = '%v' WHERE `id` = %v;", tableName, temp_t_date, dayInfo_weather_info, dayInfo_weather_temperature, dayInfo_wind_direct, dayInfo_wind_power, nightInfo_weather_info, nightInfo_weather_temperature, nightInfo_wind_direct, nightInfo_wind_power, everyday_data.Id)
						rowCount, err := db.ExecSql(updatesql, uuid)
						if err != nil {
							logger.Errorf("%v %v %v %v", uuid, "更新失败", updatesql, err)
						} else {
							logger.Infof("%v %v %v %v", uuid, tableName, "更新成功,row:", rowCount)
						}
					}
				}

			}
		}
	default:
		logger.Errorf("%v %v %v %v", uuid, stationid, respDataValue, "没数据----ERROR----")
	}

}

func SaveProvinceCityData(resp []byte, uuid string) {
	logger.Infof("%v %v", uuid, "解析市区县信息")
	var ProvinceData []map[string]interface{}
	err := json.Unmarshal(resp, &ProvinceData)
	if err != nil {
		logger.Errorf("%v 解析市区县信息出错 %v", uuid, err)
	}
	// respData := dataAttr["data"].(map[string]interface{})
	for _, hash_map_dict := range ProvinceData {
		saveLocationDataAndTable(hash_map_dict, uuid)
	}
}

func saveLocationDataAndTable(hash_map_dict map[string]interface{}, uuid string) {
	station_id := hash_map_dict["code"].(string)
	province := hash_map_dict["province"].(string)
	city := hash_map_dict["city"].(string)
	pcurl := hash_map_dict["url"].(string)
	logger.Debugf("%v %v %v %v", uuid, station_id, province, city)

	get_location_sql := fmt.Sprintf("select * from location where stationid = '%v' order by id desc limit 1", station_id)
	location_data := db.GetLocationData(get_location_sql, uuid)
	if (location_data.Province == province) && (location_data.City == city) {
		logger.Debugf("%v-%v", uuid, "省市已存在")
	} else {
		insert_location_str := fmt.Sprintf("INSERT INTO location (stationid,country,province,city,url,valid) VALUES ('%v','中国', '%v', '%v', '%v', '%v');", station_id, province, city, pcurl, 1)
		_pk := db.InsertRow(insert_location_str, uuid)
		logger.Infof("%v %v-%v%v", uuid, "location", "insert success,pk:", _pk)
	}
	yearStr := strconv.FormatInt(int64(time.Now().Year()), 10)
	tableName := station_id + "_" + yearStr
	table_check_sql := fmt.Sprintf("select id from %v limit 1", tableName)
	_, err := db.GetLocationRec(table_check_sql, uuid)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			logger.Debugf("%v %v %v", uuid, tableName, "开始预报表创建")
			create_table_str := fmt.Sprintf("CREATE TABLE `%v` (`id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,`date` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日期',`day_info` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '白天天气',`day_temperature` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '白天温度',`day_direct` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '白天风向',`day_power` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '白天风力',`night_info` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '夜间天气',`night_temperature` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '夜间温度',`night_direct` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '夜间风向',`night_power` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '夜间风力',PRIMARY KEY (`id`) USING BTREE) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;", tableName)
			_, err := db.ExecSql(create_table_str, uuid)
			if err != nil {
				logger.Errorf("%v %v %v %v", uuid, tableName, "建预报表失败", err)
			} else {
				logger.Debugf("%v %v %v", uuid, tableName, "建预报表成功")
			}
		} else {
			logger.Debugf("%v %v %v", uuid, tableName, "预报表已存在")
		}

	} else {
		logger.Debugf("%v %v %v", uuid, tableName, "预报表已存在")
	}

	tableRName := station_id + "r_" + yearStr
	rtable_check_sql := fmt.Sprintf("select id from %v limit 1", tableRName)
	_, err = db.GetLocationRec(rtable_check_sql, uuid)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			logger.Debugf("%v %v %v", uuid, tableRName, "开始实时表创建")
			create_rtable_str := fmt.Sprintf("CREATE TABLE `%v`  (`id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,`date` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日期',`time` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '时间',`temperature` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '温度',`humidity` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '相对湿度',`rain` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '降水量mm',`icomfort` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '舒适度',`info` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '天气',`feelst` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '体感温度',`wind_direct` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '风向',`wind_power` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '风力',`wind_speed` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '风速',`warn` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '预警',`aqi` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '空气质量',`aq` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '空气质量',PRIMARY KEY (`id`) USING BTREE) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;", tableRName)
			_, err := db.ExecSql(create_rtable_str, uuid)
			if err != nil {
				logger.Errorf("%v %v %v %v", uuid, tableRName, "建实时表失败", err)
			} else {
				logger.Debugf("%v %v %v", uuid, tableRName, "建实时表成功")
			}
		} else {
			logger.Debugf("%v %v %v", uuid, tableRName, "实时表已存在")
		}

	} else {
		logger.Debugf("%v %v %v", uuid, tableName, "实时表已存在")
	}
}
