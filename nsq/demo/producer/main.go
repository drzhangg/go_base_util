package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"go_base_util/private"
)

func main() {
	//定义nsq生产者
	var producer *nsq.Producer
	//初始化生产者
	producer, err := nsq.NewProducer(private.ADDR+":4150", nsq.NewConfig())
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

		if producer != nil && message != "" { //不能发布空串，会导致error
			err = producer.Publish(topic, []byte(message))
			if err != nil {
				fmt.Printf("producer publish err:%v", err)
			}
			fmt.Println(message)
		}
	}
	fmt.Println("producer Publish success")
}
