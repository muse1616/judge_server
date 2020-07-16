package main

import (
	"encoding/json"
	"judge_server/model"
	"judge_server/mq"
	"judge_server/utils"
	"log"
)

func main() {
	// 获取rabbitMQ连接
	conn, err := mq.GetRabbitMQConnection()
	utils.FailOnError(err, "Get RabbitMQ Connection Failed")
	// 关闭连接
	defer conn.Close()

	// 通道
	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 队列声明
	q, err := ch.QueueDeclare("code_submit", true, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare a queue")

	// 注册消费者
	// 关闭自动确认 autAck:false 同时一次只给一个消息
	err = ch.Qos(1, 0, false)
	utils.FailOnError(err, "Failed to set qos")

	// 内容消费
	content, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	utils.FailOnError(err, "Failed to register a consumer")

	//通道 异步读取 开启通道是为了阻塞
	forever := make(chan []byte)

	go func() {
		for d := range content {
			var c model.CodeModel
			err = json.Unmarshal(d.Body, &c)
			utils.FailOnError(err, "Code struct error")

			// 此处开启子进程处理代码
			log.Printf("%+v\n", c)

			//确认代码处理
			_ = d.Ack(false)

			//后续操作
		}
	}()

	log.Printf(" [*] Waiting for code.")
	<-forever
}
