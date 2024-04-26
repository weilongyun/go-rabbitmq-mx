package main

import (
	"fmt"
	"log"
	"rabbitmq/simple"
)

//消费简单消息
func main() {
	simpleMsg := "Hello World simple message01!"
	res := simple.NewRabbitmqSimple("hello_queue_simple01")
	err := res.SendSimpleMsg(simpleMsg)
	if err != nil {
		log.Printf(" sendSimpleMsg simple error", err)
		panic(err)
	}
	fmt.Println("发送成功")
}
