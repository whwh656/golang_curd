package api

import (
	"curd/controller"
    "github.com/gin-gonic/gin"
)

//Registerapi 外部调用必须首字母大写
func Registerapi(r *gin.Engine) *gin.Engine {
	api := r.Group("api/v1")
	api.GET("/ping", controller.PingAPI)
	api.POST("/users", controller.CreateUser)
	return r
}
