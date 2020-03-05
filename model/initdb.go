package model

import (
	"curd/config"
    "fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB


//InitMySQL initdb
func InitMySQL() (*gorm.DB, error) {
	mysqlConfig := config.GetdbConf()
	db, err := gorm.Open("mysql", mysqlConfig.Mysql.Dsn)
	DB = db
	db.LogMode(true)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return 0
	// }else{
	if err == nil {
		db.AutoMigrate(&User{})
		return db, err
	}
	// defer db.Close()
	return nil, err
}

//AddUser add one user
func (user *User) AddUser() error {
	return DB.Create(user).Error
}

// ListUser list all table user 
// func (users *User) ListUser() error {
// 	return  DB.Find(&users).Error
// }

func DeleteUser(p *User, id string) (err error) {
	if err := DB.Where("id = ?",id).Delete(p).Error; err != nil {
		return err
	}
	return nil
}


//UpdateUser 
func UpdateUser(p *User, age int,name string) (err error) {
	DB.LogMode(true)
	fmt.Println(age,name)
	if err := DB.Debug().Model(&p).Where("name = ? ",name).Update("age",age).Error; err != nil {
		return err
	} 
	return nil
}

