package model

import (
	"curd/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//InitMySQL initdb
func InitMySQL() (*gorm.DB, error) {
	mysqlConfig := config.GetdbConf()
	db, err := gorm.Open("mysql", mysqlConfig.Mysql.Dsn)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return 0
	// }else{
	if err == nil {
		db.AutoMigrate(&User{})
	}
	defer db.Close()
	return nil, err
}
