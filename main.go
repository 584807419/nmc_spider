package main

import (
	"nmc_spider/data_operation"
	"time"
)

func main() {
	data_operation.GetProvinceData()
	for {
		data_operation.GetData()
		time.Sleep(4 * time.Second)
	}
}
