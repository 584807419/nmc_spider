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
		fmt.Println("等半个小时再采")
		time.Sleep(1800 * time.Second)
	}

}
