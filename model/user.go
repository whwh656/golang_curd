package model

import (
	"github.com/jinzhu/gorm"
)

// User user对象
type User struct {
	gorm.Model
	Name string
	age  int
}
