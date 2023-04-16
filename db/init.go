package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB = initDB() // 常量

// 建立连接
func initDB() (DB *sqlx.DB) {
	dsn := "root:Zk199245*99@tcp(192.168.1.16:3306)/www_nmc_cn?charset=utf8mb4&parseTime=False&loc=Local&tls=false"
	// 也可以使用MustConnect连接不成功就panic
	DB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	return DB
}
