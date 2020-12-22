package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
)

func main() {
	//定义nsq生产者
	var producer *nsq.Producer
	//初始化生产者
	producer, err := nsq.NewProducer("47.103.9.218:4150", nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	err = producer.Ping()
	if err != nil {
		//关闭生产者
		producer.Stop()
		producer = nil
	}

	topic := "test_topic"

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("message:%d", i)

		err = producer.Publish(topic,[]byte(message))
		if err != nil {
			fmt.Printf("producer publish err:%v",err)
		}
		fmt.Println(message)
	}
	fmt.Println("producer Publish success")
}
