package data_operation

import (
	"encoding/json"
	"fmt"
	"nmc_spider/db"
	"strconv"
	"strings"
	"time"
)

func saveRtableData(respData map[string]interface{}) {
	realData := respData["real"].(map[string]interface{})
	realWeatherPublishTime := realData["publish_time"].(string)
	temp_t, _ := time.ParseInLocation("2006-01-02 15:04", realWeatherPublishTime, time.Local)
	temp_t_date := temp_t.Format("2006-01-02")
	temp_t_time := temp_t.Format("15:04")
	realWeatherData := realData["weather"].(map[string]interface{})

	stationId := realData["station"].(map[string]interface{})["code"].(string)
	yearStr := strconv.FormatInt(int64(time.Now().Year()), 10)
	rtableName := stationId + "r" + "_" + yearStr

	temperaturefloat := realWeatherData["temperature"].(float64)
	temperature := fmt.Sprintf("%.1f", temperaturefloat)
	humidity := fmt.Sprintf("%.1f", realWeatherData["humidity"].(float64))
	rain := fmt.Sprintf("%.1f", realWeatherData["rain"].(float64))
	icomfort := fmt.Sprintf("%.0f", realWeatherData["icomfort"].(float64))
	info := realWeatherData["info"].(string)
	feelst := fmt.Sprintf("%.1f", realWeatherData["feelst"].(float64))

	realWindData := realData["wind"].(map[string]interface{})
	wind_direct := realWindData["direct"].(string)
	if wind_direct == "9999" {
		wind_direct = "无"
	}
	wind_power := realWindData["power"].(string)
	wind_speed := fmt.Sprintf("%.1f", realWindData["speed"].(float64))

	realWarnData := realData["warn"].(map[string]interface{})
	signaltype := realWarnData["signaltype"].(string)
	signallevel := realWarnData["signallevel"].(string)
	issuecontent := realWarnData["issuecontent"].(string)
	warn_str := signaltype + ":" + signallevel + "\n" + issuecontent
	if signaltype == "9999" {
		warn_str = ""
	}

	airData := respData["air"].(map[string]interface{})
	aqi := fmt.Sprintf("%.0f", airData["aqi"].(float64))
	aq := fmt.Sprintf("%.0f", airData["aq"].(float64))

	rtableNameSqlStr := fmt.Sprintf("insert into %v (date, time, temperature,humidity,rain,icomfort,info,feelst,wind_direct,wind_power,wind_speed,warn,aqi,aq) values ('%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v')", rtableName, temp_t_date, temp_t_time, temperature, humidity, rain, icomfort, info, feelst, wind_direct, wind_power, wind_speed, warn_str, aqi, aq)
	if !strings.Contains(rtableNameSqlStr, "9999") {
		db.InsertRow(rtableNameSqlStr)
	}
}

func savetableData(respData map[string]interface{}) {
	realData := respData["real"].(map[string]interface{})

	stationId := realData["station"].(map[string]interface{})["code"].(string)
	yearStr := strconv.FormatInt(int64(time.Now().Year()), 10)
	tableName := stationId + "_" + yearStr

	predictData := respData["predict"].(map[string]interface{})
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
			everyday_data := db.GetData(getOneData)
			if (everyday_data.Day_info == dayInfo_weather_info) && (everyday_data.Day_temperature == dayInfo_weather_temperature) && (everyday_data.Night_info == nightInfo_weather_info) && (everyday_data.Night_temperature == nightInfo_weather_temperature) {
				fmt.Println("无新数据")
			} else {
				db.InsertRow(tableNameSqlStr)
			}
		}
	}
}

func SaveData(resp []byte) {
	var dataAttr map[string]interface{}
	err := json.Unmarshal(resp, &dataAttr)
	if err != nil {
		fmt.Println(err)
	}
	respData := dataAttr["data"].(map[string]interface{})
	saveRtableData(respData)
	savetableData(respData)
}
