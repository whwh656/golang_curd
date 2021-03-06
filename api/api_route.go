package api

import (
	"curd/controller"
	"github.com/gin-gonic/gin"
	"fmt"
)

const (
	TOKEN = "123456"
)

//gin认证中间件的使用
func TokenAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context){
		//请求头中获取token参数
		token := c.Request.Header.Get("token")
		//token := c.Request.FormValue("token")
		fmt.Printf("%T",token)
		fmt.Println("sas",len(token))
		fmt.Println("aaaa",len("123456"))


		fmt.Println(token)
		if token ==""{
			c.JSON(401,"api token required")
			c.Abort()
			return
		}
		if token != TOKEN {
			c.JSON(401,"Invalid api token")
			c.Abort()
			return
		}
		// c.Next作用？
		c.Next()  
	}
}

//Registerapi 外部调用必须首字母大写
func Registerapi(r *gin.Engine) *gin.Engine {
	api := r.Group("api/v1")
	//api.use中间的作用范围是下面所有的接口
	// api.Use(TokenAuthMiddleware())
	// api.GET("/ping", controller.PingAPI)
	// api.POST("/users", controller.CreateUser)
	//方法2:作用于部分接口
	api.GET("/ping",TokenAuthMiddleware(),controller.PingAPI)
	api.POST("/users", controller.CreateUser)
	api.GET("/users",controller.ListUser)
	api.PUT("/users/",controller.UpdateUser)
	api.DELETE("/user/:uid",controller.DelUser)
    return r
}
