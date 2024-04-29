package main

import (
	"fmt"
	"go-rabbitmq-mx/rabbitmq/mode"
	"log"
)

//消费简单消息
func main() {
	Msg := "Hello World pubsub message!"
	res := mode.NewRabbitmqPubSub("hello_exchange_pubsub")
	for i := 0; i < 50; i++ {
		err := res.SendPubSubMsg(Msg + fmt.Sprintf("%d", i))
		if err != nil {
			log.Printf(" sendSimpleMsg mode error", err)
			panic(err)
		}
	}
	fmt.Println("发送pubsub成功")
}
