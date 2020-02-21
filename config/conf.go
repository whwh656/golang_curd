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

// GetConf 获取yaml中的port字符串
func GetConf() *WrapConfInfo {
	var confInfo WrapConfInfo
	yamlFile, err := ioutil.ReadFile("./config/app.yaml")
	if err != nil {
		fmt.Println("error")
	}
	yaml.Unmarshal(yamlFile, &confInfo)
	fmt.Println(confInfo.Server.Port)
	return &confInfo
}
