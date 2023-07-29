package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB = initDB() // 饿汉式单例

// 建立连接池
func initDB() (DB *sqlx.DB) {
	fmt.Printf("建立连接池")
	// 内网IP单独用户保证安全
	dsn := "nmc_spider:nmc_spider@tcp(127.0.0.1:32768)/www_nmc_cn?charset=utf8mb4&parseTime=False&loc=Local&tls=false"
	// 也可以使用MustConnect连接不成功就panic
	DB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("建立连接池失败, err:%v\n", err)
		return
	}
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	return DB
}

// SetMaxOpenConns()方法允许您设置池中“打开”连接(使用中+空闲连接)数量的上限。默认情况下，打开的连接数是无限的。
// SetMaxIdleConns()方法的作用是：设置池中空闲连接数的上限。缺省情况下，最大空闲连接数为2。
// SetConnMaxLifetime()方法用于设置ConnMaxLifetime的极限值，表示一个连接保持可用的最长时间。默认连接的存活时间没有限制，永久可用。
// 如果设置ConnMaxLifetime的值为1小时，意味着所有的连接在创建后，经过一个小时就会被标记为失效连接，标志后就不可复用。但需要注意：
// 这并不能保证一个连接将在池中存在一整个小时；有可能某个连接由于某种原因变得不可用，并在此之前自动关闭。
// 连接在创建后一个多小时内仍然可以被使用—只是在这个时间之后它不能被重用。
// 这不是一个空闲超时。连接将在创建后一小时过期，而不是在空闲后一小时过期。
// Go每秒运行一次后台清理操作，从池中删除过期的连接。
// 设置最大生存时间为1小时
// 设置为0，表示没有最大生存期，并且连接会被重用
// forever (这是默认配置).
// DB.SetConnMaxLifetime(time.Hour)
// SetConnMaxIdleTime()方法在Go 1.15版本引入对ConnMaxIdleTime进行配置。其效果和ConnMaxLifeTime类似，但这里设置的是：在被标记为失效之前一个连接最长空闲时间。例如，如果我们将ConnMaxIdleTime设置为1小时，那么自上次使用以后在池中空闲了1小时的任何连接都将被标记为过期并被后台清理操作删除。
