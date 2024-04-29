package main

import "go-rabbitmq-mx/rabbitmq/mode"

//消费pubsub消息
func main() {
	res := mode.NewRabbitmqPubSub("hello_exchange_pubsub")
	err := res.ConsumerPubSubMsg()
	if err != nil {
		panic(err)
	}

}
