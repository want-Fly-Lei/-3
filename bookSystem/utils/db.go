package utils

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	err error
	Db *gorm.DB
)

//用于初始化数据库连接
func init() {
	var driverName string = "mysql"
	var host string = "localhost"
	var port string = "3306"
	var database string = "booksystem"
	var username string = "root"
	var charset string = "utf8"
	var password string = "12345"
	var args string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,password,
		host,port,
		database,charset,
	)

	Db, err = gorm.Open(driverName,args)

	if err != nil {
		panic("failed to connect,err:" + err.Error())
	}
}
