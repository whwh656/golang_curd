package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// WrapConfInfo portinfomation
type WrapConfInfo struct {
	Server ConfInfo `yaml:"Server"`
}

// ConfInfo port
type ConfInfo struct {
	Port        string `yaml:"Port"`
	Environment string `yaml:"Environment"`
}

//WrapMysql mysqlconf
type WrapMysql struct {
	Mysql MysqlDB `yaml:"Mysql"`
}

//MysqlDB dbconfig
type MysqlDB struct {
	Dsn          string `yaml:"Dsn"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxLifeConns int    `yaml:"max_life_conns"`
}

// GetConf 获取yaml中的port字符串
func GetConf() *WrapConfInfo {
	var confInfo WrapConfInfo
	yamlFile, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		fmt.Println("error")
	}
	yaml.Unmarshal(yamlFile, &confInfo)
	return &confInfo
}

// GetdbConf 获取yaml中的dsn字符串
func GetdbConf() *WrapMysql {
	var mysqlconf WrapMysql
	yamlFile, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		fmt.Println("error")
	}
	yaml.Unmarshal(yamlFile, &mysqlconf)
	return &mysqlconf
}
