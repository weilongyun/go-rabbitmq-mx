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

#rabbitmq工作模式
1、简单simple模式:也就是说生产者发送消息直接发送到队列中，只有一个消费者消费
2、worker模式:可以有多个消费者消费队列中的消息，每个消息只能被消费一次(实际上和简单模式的区别就是有多个消费者
代码和简单模式是一模一样的)
3、订阅模式publish/subcribe:
和前面二种方式不同，简单模式和worke queue工作队列模式一个消息投递到队列中只能被一个消费者消费，订阅模式是把一个消息投递到
多个队列中，消费者从不同队列中可以消费同一个消息，并且发消息先发送到excahnge交换机中
4、routing模式
消费者通过routing key来指定发往哪个队列，也就是说一个消息可以被多个消费者消费，并且消息队列可以被生产者指定

总结：
简单模式需要指定队列名，交换机不需要指定，routing key指定队列名称
work工作模式和简单模式一样的，唯一不同的就是允许有多个消费者，但是不同消费者是不能消费同一个消息的
pubsub发布订阅模式需要指定交换机，交换机随机不需要指定，，routing key也不需要，为空
routing模式需要指定交换机和routing key，队列随机不需要指定