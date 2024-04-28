package main

import "rabbitmq/mode"

//消费简单消息
func main() {
	res := mode.NewRabbitmqRouting("hello_exchange_routing", "routing_key1")
	err := res.ConsumerRoutingMsg()
	if err != nil {
		panic(err)
	}

}
