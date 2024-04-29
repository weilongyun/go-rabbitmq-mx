package main

import "go-rabbitmq-mx/rabbitmq/mode"

//消费简单消息
func main() {
	res := mode.NewRabbitmqSimple("hello_queue_simple01")
	err := res.ConsumerSimpleMsg()
	if err != nil {
		panic(err)
	}

}
