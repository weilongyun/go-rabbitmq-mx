package main

import (
	"fmt"
	"log"
	"rabbitmq/mode"
)

//消费简单消息
func main() {
	Msg1 := "Hello World  routing message1!"
	Msg2 := "Hello World  routing message2!"
	res1 := mode.NewRabbitmqRouting("hello_exchange_routing", "routing_key1")
	res2 := mode.NewRabbitmqRouting("hello_exchange_routing", "routing_key2")
	for i := 0; i < 50; i++ {
		err := res1.SendRoutingMsg(Msg1 + fmt.Sprintf("%d", i))
		if err != nil {
			log.Printf(" sendRoutingMsg mode error", err)
			panic(err)
		}
		err = res2.SendRoutingMsg(Msg2 + fmt.Sprintf("%d", i))
		if err != nil {
			log.Printf(" sendRoutingMsg mode error", err)
			panic(err)
		}
	}
	fmt.Println("发送sendRouting消息成功")
}
