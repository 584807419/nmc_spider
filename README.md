# nmc_spider

# 天气数据爬虫

## 1.生产
任务分发的Goroutine->无缓冲Channel
## 2.消费
数据抓取的Goroutine->100缓冲Channel
## 3.保存
100缓冲Channel->解析json、html数据并保存到MySql的Goroutine



## 控制频率
1. 控制消费者Goroutine数量
2. 控制Channel缓冲数值
3. 消费者读Channel用了定时器Ticker

## linux打包
`SET CGO_ENABLED=0`

`SET GOARCH=amd64`

`SET GOOS=linux`

`go build nmc_spider`

## 数据库
`docker pull mariadb`

`docker run -d -P --name 名字 -e MYSQL_ROOT_PASSWORD=密码 mariadb`
   
-P 参数，Docker 自动分配一个未被使用的端口，docker ps -a 看一下端口号

挂载数据卷，以挂载本地 /data/Db/mariadb 为例

`docker run -d --name mariadb_volume -P -v /data/Db/mariadb:/var/lib/mysql mariadb`