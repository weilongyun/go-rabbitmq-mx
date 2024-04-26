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

##rabbitmq常见的交换机类型
1、fanout，也就是说说交换机类型是fanout，需要绑定队列，最后发的消息通过fanout交换机会发送到所有与该
交换机绑定的队列中
2、direct交换机，是需要一个routing key，通过routing key来告诉需要绑定到哪个队列中，此时是可以选择队列的
如果你定义的队列是同一个routing key，那么direct交换机和fanout交换机效果就一样了
3、topic交换机
topic交换机中的routing key是支持多个单词的，用点.分割，并且支持正则，非常灵活
#代表0个或多个单词
*代表一个单词

