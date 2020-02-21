package controller

import "github.com/gin-gonic/gin"

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

//CreateUser Post add a user
func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong!",
	})
}
