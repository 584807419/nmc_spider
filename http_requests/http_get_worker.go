package http_requests

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"nmc_spider/message_queue"
	"sync"
	"sync/atomic"
	"time"
)

var concurrent int32

func HttpGetWorker(wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("HttpGetWorker-error%v", err)
		}
	}()
	ticker := time.NewTicker(time.Second * 5) // Ticker是周期性定时器，即周期性的触发一个事件
	for range ticker.C {                      // 定时器：每隔1秒发一次请求，生产者发过来的多快咱这也要控制好速度
		urlHashMap, ok := <-message_queue.TempUrlChan
		// logger.Infof("HttpGetWorker %v", "从无缓冲通道中获取")
		uuid := urlHashMap["uuid"]
		stationid := urlHashMap["stationid"]
		if ok {
			url, okjson := urlHashMap["url"]
			// 请求json接口
			if okjson {
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
				atomic.AddInt32(&concurrent, 1)
				fmt.Printf("并发请求数%d\n", atomic.LoadInt32(&concurrent))
				resp, err := HttpClient.Do(req)
				atomic.AddInt32(&concurrent, -1)
				if err != nil {
					logger.Errorf("%v HttpGet1 %v", uuid, err)
					continue
				} else {
					body, err := io.ReadAll(resp.Body)
					if err != nil {
						logger.Errorf("%v HttpGet2 %v", uuid, err)
					}
					respDataHashMap := map[string]interface{}{
						"uuid":      uuid,
						"resp_body": body,
						"stationid": stationid,
					}
					// logger.Infof("%v HttpGetWorker %v", uuid, "往100缓冲通道发送json数据")
					message_queue.TempRespDataChan <- respDataHashMap

				}
				err = resp.Body.Close()
				if err != nil {
					logger.Errorf("%v close http request %v", uuid, err)
				}
			}
			// 请求html页面
			html_url, okhtml := urlHashMap["html_url"]
			if okhtml {
				logger.Infof("%v HttpGetWorker %v", uuid, "拿到html链接")
				req, _ := http.NewRequest("GET", html_url, nil)
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
				atomic.AddInt32(&concurrent, 1)
				fmt.Printf("并发请求数%d\n", atomic.LoadInt32(&concurrent))
				resp, err := HttpClient.Do(req)
				atomic.AddInt32(&concurrent, -1)
				if err != nil {
					logger.Errorf("%v HttpGet1 %v", uuid, err)
					continue
				} else {
					// //获取响应体
					// bodyReader := bufio.NewReader(resp.Body)
					// //使用determiEncoding函数对获取的信息进行解析
					// e := determiEncoding(bodyReader)
					// utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
					// //读取并打印获取的信息
					// body, err := io.ReadAll(utf8Reader)
					// if err != nil {
					// 	panic(err)
					// }
					// fmt.Printf("%s", body)
					// print(body)
					// print(string(body))
					// resp.Body.Close()
					if resp.Header.Get("Content-Encoding") == "gzip" {
						// 响应是用gzip压缩，需要gzip解压，解决乱码问题
						reader, _ := gzip.NewReader(resp.Body)
						body, err := io.ReadAll(reader)
						if err != nil {
							logger.Errorf("%v HttpGet2 %v", uuid, err)
						}

						respDataHashMap := map[string]interface{}{
							"uuid":           uuid,
							"resp_html_body": body,
							"stationid":      stationid,
						}
						// logger.Infof("%v HttpGetWorker %v", uuid, "往100缓冲通道发送html数据")
						message_queue.TempRespDataChan <- respDataHashMap
					}

				}
				err = resp.Body.Close()
				if err != nil {
					logger.Errorf("%v Close %v", uuid, err)
				}
			}

		}
	}
	wg.Done()
}
