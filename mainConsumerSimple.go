package main

import "rabbitmq/simple"

//消费简单消息
func main() {
	res := simple.NewRabbitmqSimple("hello_queue_simple01")
	err := res.ConsumerSimpleMsg()
	if err != nil {
		panic(err)
	}

}
