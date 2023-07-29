package main

import (
	"nmc_spider/data_operation"
	"nmc_spider/http_requests"
	"sync"
)

func main() {
	//data_operation.GetProvinceData() // 获取信息并且根据信息建表
	var wg sync.WaitGroup
	wg.Add(7) //计数器 +3

	// 向无缓冲通道中持续发送
	go data_operation.GetData(&wg)

	// 三个worker每隔5s去无缓冲通道中获取一次然后发请求
	go http_requests.HttpGetWorker(&wg)
	go http_requests.HttpGetWorker(&wg)
	go http_requests.HttpGetWorker(&wg)

	go data_operation.ParsingDataWorker(&wg)
	go data_operation.ParsingDataWorker(&wg)
	go data_operation.ParsingDataWorker(&wg)
	wg.Wait() // 阻塞，直到 WaitGroup 的计数器的值为 0 才会解除阻塞状态，爬虫需要一直运行，不会调用 wg.Done 将计数器 -1
}
