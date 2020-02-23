package controller

import (
	"fmt"

	"curd/model"

	"github.com/gin-gonic/gin"
)

// PingAPI pingapi return pong
func PingAPI(c *gin.Context) {
	// c.JSON：返回JSON格式的数据
	c.JSON(200, gin.H{
		"message": "Pong!",
	})

}

//ListUser get users
func ListUser() {

}

//CreateUser Post add a user  如何获取json数据存入mysql数据库
func CreateUser(c *gin.Context) {
	user := model.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"errcode":     400,
			"description": "Post Data Err",
		})
	} else {
		fmt.Println(user)
		//存入数据库
	}
}
