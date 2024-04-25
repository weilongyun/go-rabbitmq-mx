package main

import (
	"fmt"
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
func newRabbitmqSimple(queueName string) *Rabbitmq {
	return newRabbitmq(queueName, "", "")
}

//rabbitmq官网：https://www.rabbitmq.com/tutorials
//发送普通消息，客户端直接发送消息到队列中，消费者直接从队列中消费
//先创建队列
func (r *Rabbitmq) sendSimpleMsg(msg string) error {
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
func main() {
	simpleMsg := "Hello World simple message01!"
	res := newRabbitmqSimple("hello_queue_simple01")
	err := res.sendSimpleMsg(simpleMsg)
	fmt.Println(err)
}
