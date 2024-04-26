# go-rabbitmq-mx
go-rabbitmq各种工作模式进行发消息以及商品的秒杀功能
主要技术栈：golang、mysql、redis、rabbitmq、iris框架
rabbitmq-docker安装：https://blog.csdn.net/Relievedz/article/details/131081440

#注意：在消费者代码中，这样是不会发生死锁的，你要是放在main包下就会发生deadlock
###
var forever chan struct{}
	go func() {
		for d := range msg {
			log.Printf("Received a message: %s\n", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//阻塞
	<-forever
###