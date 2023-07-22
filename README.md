# nmc_spider

# 天气数据爬虫

## 1.生产
任务分发的Goroutine->无缓冲Channel
## 2.消费
数据抓取的Goroutine->100缓冲Channel
## 3.保存
100缓冲Channel->解析json、html数据并保存到MySql的Goroutine



控制消费者Goroutine或者Channel缓冲数值即可控制频率