package model

import (
	"github.com/jinzhu/gorm"
)

// User user对象 如何迁移结构体到表中 属性首字母大写不然外不访问不到
type User struct {
	gorm.Model
	Name string
	Age  int
}

var user User
