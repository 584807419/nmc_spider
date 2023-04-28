package main

import (
	"nmc_spider/data_operation"
	"nmc_spider/http_requests"
	"sync"
)

func main() {
	// data_operation.GetProvinceData()
	var wg sync.WaitGroup
	wg.Add(3)
	go data_operation.GetData(&wg)
	go http_requests.HttpGetWorker(&wg)
	go data_operation.SaveDataWorker(&wg)
	wg.Wait()
}
