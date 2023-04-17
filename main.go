package main

import (
	"fmt"
	"nmc_spider/data_operation"
	"time"
)

func main() {
	for {
		fmt.Println("开始采集")
		data_operation.GetData()
		time.Sleep(300 * time.Second)
	}

}
