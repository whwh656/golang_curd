package main

// 1 如何读取yaml中的配置文件先运行起来
// 2 如何把api封装在api的文件夹下
// 3 包内的main.go的调试
// 4 如何打印日志

import (
	"curd/api"
	"curd/config"
	"curd/model"
	"fmt"
    "github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
		if _, err := model.InitMySQL(); err != nil {
			fmt.Println("mysql init error")
		} 
	    r := gin.Default()
		route := api.Registerapi(r)
		conf := config.GetConf()
		route.Run(conf.Server.Port)
}
