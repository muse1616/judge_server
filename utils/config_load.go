package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// 配置文件接口
type Config interface {
	// 读取配置文件
	LoadYamlConfig(configName string)(content map[string]map[string]interface{}, err error)
}

// yaml
type YamlConfig struct {
}

// 目录即配置文件后缀 静态变量
const path = "config/"
const suffix = ".yaml"

// 接口实现
func (yc YamlConfig) LoadYamlConfig(configName string)(content map[string]map[string]interface{}, err error) {
	content = make(map[string]map[string]interface{})
	fileName := path + configName + suffix
	// 读取文件
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	// 转为json
	err = yaml.Unmarshal(yamlFile, content)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
