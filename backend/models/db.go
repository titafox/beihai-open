package models

import (
	"coltstail-server/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var Eloquent *gorm.DB

func init() {
	var err error
	config := conf.Conf()
	log.Printf("Connecting to database %s:%d...", config.DB_HOST, config.DB_PORT)
	Eloquent, err = gorm.Open("mysql", config.DB_CONN_STR)
	if err != nil {
		Eloquent.Close()
		log.Panicln(err)
	}
	// 设置连接池，空闲连接
	Eloquent.DB().SetMaxIdleConns(500)
	Eloquent.SingularTable(true)
	// 打开链接
	Eloquent.DB().SetMaxOpenConns(2000)

	//defer Eloquent.Close()
	Eloquent.LogMode(true)
}
