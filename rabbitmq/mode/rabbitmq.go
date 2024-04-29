package mode

import (
	"github.com/streadway/amqp"
	"log"
)

// amqp://用户名:密码@ip:port/virtualhost虚拟机
//virtualhost虚拟机可以隔离开发环境和生产环境的，比较方便
const mqUrl = "amqp://weilongyun:weilongyun@localhost:5673/test1_host"

type Rabbitmq struct {
	conn     *amqp.Connection //链接
	channel  *amqp.Channel
	Excahnge string
	Queue    string
	Key      string
	MqUrl    string
}

func newRabbitmq(queue, excahnge, key string) *Rabbitmq {
	r := &Rabbitmq{
		Excahnge: excahnge,
		Queue:    queue,
		Key:      key,
		MqUrl:    mqUrl,
	}
	var err error
	r.conn, err = amqp.Dial(r.MqUrl)
	r.failOnErr(err, "创建链接失败")
	r.channel, err = r.conn.Channel()
	r.failOnErr(err, "创建channel失败")
	return r
}

//断开cahnnel和connection
func (r *Rabbitmq) DestoryConn() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *Rabbitmq) failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		//panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

//创建rabbitmq简单模式
func NewRabbitmqSimple(queueName string) *Rabbitmq {
	return newRabbitmq(queueName, "", "")
}

//创建发布订阅模式(发消息通过交换机发送到多个队列中，不同消费者可以通过不同队列消费同一个消息)
//这种模式下不需要指定队列名称，只需要指定交换机名称
func NewRabbitmqPubSub(exchangeName string) *Rabbitmq {
	return newRabbitmq("", exchangeName, "")
}

//初始化Routin模式
func NewRabbitmqRouting(exchangeName, routing_key string) *Rabbitmq {
	return newRabbitmq("", exchangeName, routing_key)
}

//rabbitmq官网：https://www.rabbitmq.com/tutorials
//发送普通消息，客户端直接发送消息到队列中，消费者直接从队列中消费
//先创建队列
func (r *Rabbitmq) SendSimpleMsg(msg string) error {
	q, err := r.channel.QueueDeclare(
		r.Queue,
		false, //持久化
		false, //自动删除
		false, //是否具有排他性，true就代表只有自己可以看到这个队列，其他用户看不到
		false, //发消息到服务器是否阻塞
		nil,   //额外参数
	)
	if err != nil {
		r.failOnErr(err, "创建普通队列失败")
		return err
	}
	r.failOnErr(err, "创建队列失败")
	err = r.channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		r.failOnErr(err, "发布普通消息失败")
		return err
	}
	return nil

}

//消费普通消息
func (r *Rabbitmq) ConsumerSimpleMsg() error {
	//消费消息需要指定是哪个队列
	q, err := r.channel.QueueDeclare(
		r.Queue,
		false, //持久化
		false, //自动删除
		false, //是否具有排他性，true就代表只有自己可以看到这个队列，其他用户看不到
		false, //发消息到服务器是否阻塞
		nil,   //额外参数
	)
	if err != nil {
		r.failOnErr(err, "创建普通队列失败")
		return err
	}
	msg, err := r.channel.Consume(
		q.Name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		r.failOnErr(err, "消费消息失败")
	}

	var forever chan struct{}
	go func() {
		for d := range msg {
			log.Printf("Received a message: %s\n", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages from simple queue. To exit press CTRL+C")
	//阻塞
	<-forever
	return nil
}

//发送订阅消息
func (r *Rabbitmq) SendPubSubMsg(msg string) error {
	//创建交换机
	err := r.channel.ExchangeDeclare(
		r.Excahnge, // name
		"fanout",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments)
	)
	if err != nil {
		r.failOnErr(err, "创建订阅模式队列失败")
		return err
	}
	//发送消息
	err = r.channel.Publish(
		r.Excahnge,
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		r.failOnErr(err, "发送订阅消息失败")
		return err
	}
	return nil

}

//消费发布订阅消息
func (r *Rabbitmq) ConsumerPubSubMsg() error {
	//创建交换机
	err := r.channel.ExchangeDeclare(
		r.Excahnge, // name
		"fanout",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments)
	)
	if err != nil {
		r.failOnErr(err, "创建订阅模式队列失败")
		return err
	}
	//创建队列
	q, err := r.channel.QueueDeclare(
		"",    // name 发布订阅模式，队列一定要为空
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments)
	)
	r.failOnErr(err, "发送订阅队列QueueDeclare失败")
	//绑定队列到交换机上
	err = r.channel.QueueBind(
		q.Name,     // queue name
		"",         // routing key 这里一定为空
		r.Excahnge, // exchange
		false,
		nil,
	)
	r.failOnErr(err, "队列绑定交换机错误")
	//消费消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args)
	)
	r.failOnErr(err, "消息订阅模式消费错误")
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf(" receive message %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs from pubsub. To exit press CTRL+C")
	<-forever
	return nil
}

//根据routing key发送routing消息
func (r *Rabbitmq) SendRoutingMsg(msg string) error {
	//创建交换机
	//创建交换机
	err := r.channel.ExchangeDeclare(
		r.Excahnge, // name
		"direct",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments)
	)
	if err != nil {
		r.failOnErr(err, "创建routing模式队列失败")
		return err
	}
	//发送消息
	err = r.channel.Publish(
		r.Excahnge,
		r.Key, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		r.failOnErr(err, "发送routing工作模式下的消息失败")
		return err
	}
	return nil
}

//根据routing key消费消息
func (r *Rabbitmq) ConsumerRoutingMsg() error {
	//创建交换机
	err := r.channel.ExchangeDeclare(
		r.Excahnge, // name
		"direct",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments)
	)
	if err != nil {
		r.failOnErr(err, "创建routing模式队列失败")
		return err
	}
	//创建队列
	q, err := r.channel.QueueDeclare(
		"",    // name 发布订阅模式，队列一定要为空
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments)
	)
	r.failOnErr(err, "发送routing队列QueueDeclare失败")
	//绑定队列到交换机上
	err = r.channel.QueueBind(
		q.Name,     // queue name
		r.Key,      // routing key 这里一定不能为空
		r.Excahnge, // exchange
		false,
		nil,
	)
	r.failOnErr(err, "队列绑定交换机错误routing模式")
	//推模式
	//消费消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args)
	)
	r.failOnErr(err, "routing模式消费错误")
	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf(" receive routing message %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs from routing. To exit press CTRL+C")
	<-forever
	return nil
}
