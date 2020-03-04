package model

import (
	"github.com/jinzhu/gorm"
)

// User user对象 如何迁移结构体到表中 属性首字母大写不然外不访问不到`json:",omitempty"`
type User struct {
	gorm.Model  
	Name string `json:"-"`
	Age  int    
}

var user User
