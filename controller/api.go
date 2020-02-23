package controller

import (
	"curd/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingAPI pingapi return pong
func PingAPI(c *gin.Context) {
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
			"errcode": 400,
			"msg":     "Post Data Err",
		})
	} else {
		if err := user.AddUser(); err != nil {
			fmt.Println("add user fail!")
			c.JSON(500, gin.H{
				"errcode": 99991,
				"msg":     "insert fail",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"errcode": 999999,
			"msg":     "success",
		})
	}
}
