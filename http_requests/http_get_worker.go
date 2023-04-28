package http_requests

import (
	"io"
	"net/http"
	"nmc_spider/message_queue"
	"sync"
)

func HttpGetWorker(wg *sync.WaitGroup) {
	for {
		urlHashMap, ok := <-message_queue.TempUrlChan
		logger.Infof("HttpGetWorker %v", "从无缓冲通道中获取")
		if ok {
			url := urlHashMap["url"]
			uuid := urlHashMap["uuid"]
			stationid := urlHashMap["stationid"]
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
			resp, err := HttpClient.Do(req)
			if err != nil {
				logger.Errorf("%v HttpGet1 %v", uuid, err)
			} else {
				body, err := io.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil {
					logger.Errorf("%v HttpGet2 %v", uuid, err)
				}
				respDataHashMap := map[string]interface{}{
					"uuid":      uuid,
					"resp_body": body,
					"stationid": stationid,
				}
				logger.Infof("%v HttpGetWorker %v", uuid, "往100缓冲通道发送")
				message_queue.TempRespDataChan <- respDataHashMap
			}
		}
	}
	// wg.Done()
}
