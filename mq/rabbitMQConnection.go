package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"judge_server/utils"
	"log"
)

//rabbitMQ连接工具类
func GetRabbitMQConnection() (conn *amqp.Connection, err error) {
	// 配置文件
	var c utils.Config
	c = utils.YamlConfig{}
	content, err := c.LoadYamlConfig("mq")
	if err != nil {
		return
	}
	// url
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/%s",
		// 账号
		content["code"]["id"].(string),
		// 密码
		content["code"]["pwd"].(string),
		// url
		content["code"]["url"].(string),
		// 端口号
		content["code"]["port"].(string),
		// 虚拟主机
		content["code"]["vhost"].(string),
	)
	// 获取连接
	conn, err = amqp.Dial(url)
	if err != nil {
		log.Println(err)
	}
	// 返回连接
	return
}
