package main

import (
	"fmt"
	"go-rabbitmq-mx/rabbitmq/mode"
	"log"
)

//消费简单消息
func main() {
	simpleMsg := "Hello World mode message!"
	res := mode.NewRabbitmqSimple("hello_queue_simple01")
	for i := 0; i < 50; i++ {
		err := res.SendSimpleMsg(simpleMsg + fmt.Sprintf("%d", i))
		if err != nil {
			log.Printf(" sendSimpleMsg mode error", err)
			panic(err)
		}
	}
	fmt.Println("发送成功")
}
