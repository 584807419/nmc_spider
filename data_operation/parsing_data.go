package data_operation

import (
	"encoding/json"
	"nmc_spider/message_queue"
	"strings"
	"sync"
	"time"

	"github.com/antchfx/htmlquery"
)

// 处理json数据
func parsingJsonData(resp_body []byte, uuid, stationid string) {
	var dataAttr map[string]interface{}
	err := json.Unmarshal(resp_body, &dataAttr)
	if err != nil {
		logger.Errorf("%v json解析出错 %v, 原数据:", uuid, err, resp_body)
	} else {
		//  解决 interface conversion: interface {} is nil, not map[string]interface {}
		respDataPre, ok := dataAttr["data"]
		if ok && (respDataPre != nil) {
			respData := respDataPre.(map[string]interface{})
			// 保存数据到表中
			saveRtableData(respData, uuid, stationid)
			savetableData(respData, uuid, stationid)
		} else {
			logger.Errorf("uuid:%v 无数据 resp_body:%v", uuid, resp_body)
		}
	}
}

func parsingHtmlData(resp_html_body []byte, uuid, stationid string) {
	// var buf = bytes.NewBuffer([]byte{})
	html_tree_root, err := htmlquery.Parse(strings.NewReader(string(resp_html_body)))
	// node, err := html.Parse(strings.NewReader(string(resp_html_body)))
	// html.Render(buf, node)
	// print(buf.String())
	// rule_json_ins.Text = buf.String()
	// nodenew, err = xmlpath.Parse(strings.NewReader(rule_json_ins.Text))
	// print(nodenew)

	if err != nil {
		logger.Errorf("%v html解析出错 %v, 原数据:", uuid, err, resp_html_body)
	} else {
		logger.Infof("解析html成功")
		all_node_slice := htmlquery.Find(html_tree_root, "//div[@id=\"day0\"]/div[@class=\"hour3 hbg\"]")
		for _, item := range all_node_slice {
			time_node := htmlquery.InnerText(htmlquery.Find(item, "//div/div[1]")[0])

			// node_class := htmlquery.SelectAttr(htmlquery.Find(item, "//div")[0], "class")
			// if node_class == {
			// 	time.Now().Day()
			// 	temp_t, _ := time.ParseInLocation("2006-01-02 15:04", realWeatherPublishTime, time.Local)
			// 	temp_t_date := temp_t.Format("2006-01-02")
			// 	temp_t_time := temp_t.Format("15:04")
			// }

			weather_img := htmlquery.SelectAttr(htmlquery.Find(item, "//div/div[2]/img")[0], "src")
			weather_info := "纳尼"
			switch weather_img {
			case "http://image.nmc.cn/assets/img/w/40x40/3/0.png":
				weather_info = "晴"
			case "http://image.nmc.cn/assets/img/w/40x40/3/1.png":
				weather_info = "多云"
			case "http://image.nmc.cn/assets/img/w/40x40/3/2.png":
				weather_info = "阴"
			case "http://image.nmc.cn/assets/img/w/40x40/3/3.png":
				weather_info = "阵雨"
			case "http://image.nmc.cn/assets/img/w/40x40/3/4.png":
				weather_info = "雷阵雨"

			case "http://image.nmc.cn/assets/img/w/40x40/3/7.png":
				weather_info = "小雨"
			case "http://image.nmc.cn/assets/img/w/40x40/3/8.png":
				weather_info = "中雨"
			case "http://image.nmc.cn/assets/img/w/40x40/3/10.png":
				weather_info = "大雨"

			case "http://image.nmc.cn/assets/img/w/40x40/3/18.png":
				weather_info = "雾"

			default:
				weather_info = "纳尼"
			}
			rain_node := htmlquery.InnerText(htmlquery.Find(item, "//div/div[3]")[0])
			temperature_node := htmlquery.InnerText(htmlquery.Find(item, "//div/div[4]")[0])
			wind_speed_node := htmlquery.InnerText(htmlquery.Find(item, "//div/div[5]")[0])
			wind_direct_node := htmlquery.InnerText(htmlquery.Find(item, "//div/div[6]")[0])
			humidity_node := htmlquery.InnerText(htmlquery.Find(item, "//div/div[8]")[0])

			weathermap := make(map[string]interface{})
			weathermap["info"] = weather_info
			weathermap["humidity"] = strings.TrimSpace(strings.Replace(humidity_node, "%", "", 1))
			weathermap["rain"] = strings.TrimSpace(strings.Replace(strings.Replace(rain_node, "-", "0.0", 1), "mm", "", 1))
			weathermap["temperature"] = strings.TrimSpace(strings.Replace(temperature_node, "℃", "", 1))

			windmap := make(map[string]interface{})
			windmap["direct"] = strings.TrimSpace(wind_direct_node)
			windmap["speed"] = strings.TrimSpace(strings.Replace(wind_speed_node, "m/s", "", 1))

			realmap := make(map[string]interface{})
			realmap["publish_time"] = time.Now().Format("2006-01-02 ") + strings.TrimSpace(time_node)
			realmap["weather"] = weathermap
			realmap["wind"] = windmap

			respData := make(map[string]interface{})
			respData["real"] = realmap
			// logger.Infof("解析数据完成")
			saveRtableData(respData, uuid, stationid)
		}
	}

}

// 从队列中接收消息进行处理
func ParsingDataWorker(wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("ParsingDataWorker-error%v", err)
		}
	}()
	for {
		select {
		case respData, ok := <-message_queue.TempRespDataChan:
			if ok {
				logger.Infof("SaveDataWorker %v", "从100缓冲通道接收")
				uuid := respData["uuid"].(string)
				stationid := respData["stationid"].(string)

				resp_body, okjson := respData["resp_body"].([]byte)
				if okjson {
					// 解析json数据
					parsingJsonData(resp_body, uuid, stationid)
				}
				resp_html_body, okhtml := respData["resp_html_body"].([]byte)
				if okhtml {
					// 解析html数据
					parsingHtmlData(resp_html_body, uuid, stationid)
				}
			}

		}
		// default:
		// logger.Infof("从100缓冲通道 没拿到消息")
	}
}

// wg.Done()
