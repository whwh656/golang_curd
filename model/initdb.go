package model

import (
	"curd/config"
	"fmt"

	"github.com/jinzhu/gorm"
)

//InitMySQL initdb
func InitMySQL() int {
	mysqlConfig := config.GetdbConf()
	db, err := gorm.Open("mysql", mysqlConfig.Mysql.Dsn)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer db.Close()
	return 1
}
