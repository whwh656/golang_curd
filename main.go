package main

// 1 如何读取yaml中的配置文件先运行起来
// 2 如何把api封装在api的文件夹下
// 3 包内的main.go的调试
// 4 如何打印日志

import (
	"curd/api"
	"curd/config"
	"fmt"

	"curd/model"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	if err := model.InitMySQL(); err != 1 {
		fmt.Println("mysql init error")
	} else {
		r := gin.Default()
		route := api.Registerapi(r)
		port := config.GetConf()
		route.Run(port.Server.Port)
	}
}
