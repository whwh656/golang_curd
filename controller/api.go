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
func ListUser(c *gin.Context) {
	users := []model.User{}
	//gorm 获取数据库中的所有用户信息
    if err := model.DB.Find(&users); err != nil {
		fmt.Println("get all users fail!")
		c.JSON(500, gin.H{
			"errcode": 99991,
			"msg":    "get fail",
		})
	}
	//将数据库中获取的数据以json的格式返回给客户端
	c.JSON(http.StatusOK,gin.H{
		"errcode": 999999,
		"msg":     users,
	})
}

//CreateUser Post add a user  如何获取json数据存入mysql数据库
func CreateUser(c *gin.Context) {
	user := model.User{}
	// name := c.PostForm("name")
	// age  := c.PostForm("age")
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

//UpdateUser 更新数据
func UpdateUser(c *gin.Context){
	var user model.User
	c.BindJSON(&user)
	fmt.Println(user.Name)
	err := model.UpdateUser(&user,user.Age,user.Name)
	if err != nil {
		c.JSON(404, nil)
	} else {
		c.JSON(200, user)
	}
    
}

//删除数据 
func DelUser(c *gin.Context){
	var person model.User
	id := c.Param("uid")
	fmt.Println(id)
	err := model.DeleteUser(&person, id)
	if err != nil {
		c.JSON(400, gin.H{
			"errcode": 400,
			"msg":     "del Data Err",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"errcode": 999999,
			"msg":     "del Data Success",
		})
	}

	
}