package data_operation

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"nmc_spider/db"
	"strconv"
	"time"
)

func SendReq(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Host", "www.nmc.cn")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Language", "en-US,zh-CN;q=0.8,zh;q=0.7,zh-TW;q=0.5,zh-HK;q=0.3,en;q=0.2")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Referer", "http://www.nmc.cn/")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func GetData() {
	location := db.GetAllLocation()
	for _, value := range location {
		// fmt.Println(value.Id)
		// fmt.Println(value.Stationid)
		// fmt.Println(value.Country)
		// fmt.Println(value.Province)
		// fmt.Println(value.City)
		// fmt.Println(value.Valid)
		time_stamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
		url := "http://www.nmc.cn/rest/weather?stationid=" + value.Stationid + "&_=" + time_stamp
		resp_data := SendReq(url)
		SaveData(resp_data)
	}
}
