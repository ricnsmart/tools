package mq

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/ricnsmart/tools/util"
	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"time"
)

var Conn *amqp.Connection

const (
	connectRabbitMQFailed  = "Failed to connect to RabbitMQ"
	openChannelFailed      = "Failed to open a channel"
	declareQueueFailed     = "Failed to declare a queue"
	bindQueueFailed        = "Failed to bind a queue"
	declareExchangeFailed  = "Failed to declare an exchange"
	registerConsumerFailed = "Failed to register a consumer"
	publishMessageFailed   = "Failed to publish a message"
	setQoSFailed           = "Failed to set Qos"
	connectRabbitMQSucceed = "RabbitMQ connected!"
)

func Connect(userName, password, address string) {
	var err error

	url := fmt.Sprintf("amqp://%v:%v@%v/", userName, password, address)

	Conn, err = amqp.Dial(url)

	util.FatalOnError(err, connectRabbitMQFailed, url)

	log.Info(connectRabbitMQSucceed, " Address:", address)
}

// 普通模式
// 支持持久化
func Send(QueueName string, request []byte) error {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		QueueName,
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareQueueFailed)
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 消息 durable
			ContentType:  "text/plain",
			Body:         request,
		})

	if err != nil {
		log.Error(publishMessageFailed)
	}

	return err
}

func Receive(QueueName string) (<-chan amqp.Delivery, error) {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		QueueName,
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareQueueFailed)
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(registerConsumerFailed)
		return nil, err
	}

	return msgs, nil
}

// worker模式
func Producer(queueName string, request []byte) error {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareQueueFailed)
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 消息 durable
			ContentType:  "text/plain",
			Body:         request,
		})

	if err != nil {
		log.Error(publishMessageFailed)
		return err
	}

	return nil
}

// worker使用时，必须手动ack
func Worker(queueName string) (<-chan amqp.Delivery, error) {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareQueueFailed)
		return nil, err
	}

	// 每个worker消费完上一个消息之后，mq才会发给它送下一个消息
	// 如果worker繁忙，mq会转给其他worker
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	if err != nil {
		log.Error(setQoSFailed)
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false, // 需要手动ack，否则会阻塞
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(registerConsumerFailed)
		return nil, err
	}

	return msgs, err
}

// fanout，广播模式
// 开启持久化
func Publish(exchangeName string, request []byte) error {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareExchangeFailed)
		return err
	}

	err = ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			// 开启
			DeliveryMode: amqp.Persistent, // 消息 durable
			ContentType:  "text/plain",
			Body:         request,
		})

	if err != nil {
		log.Error(publishMessageFailed)
	}

	return nil
}

func Subscribe(exchangeName string) (<-chan amqp.Delivery, error) {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareExchangeFailed)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",
		true,
		true,
		true,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareQueueFailed)
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		"",
		exchangeName,
		false,
		nil)

	if err != nil {
		log.Error(bindQueueFailed)
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(registerConsumerFailed)
		return nil, err
	}

	return msgs, err
}

// 路由模式
// 关闭持久化
func RoutePublish(exchangeName, key string, request []byte) error {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareExchangeFailed)
		return err
	}

	err = ch.Publish(
		exchangeName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        request,
		})

	if err != nil {
		log.Error(publishMessageFailed)
	}

	return nil
}

// 一个线程用一个channel，多个go程共用一个channel
func RouteConsume(ch *amqp.Channel, exchangeName, key string) (<-chan amqp.Delivery, error) {
	err := ch.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareExchangeFailed)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		true,
		true,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareQueueFailed)
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		key,
		exchangeName,
		false,
		nil)

	if err != nil {
		log.Error(bindQueueFailed)
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(registerConsumerFailed)
		return nil, err
	}

	return msgs, err
}

// TOPIC模式
func TopicEmit(exchangeName, key string, request []byte) error {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareExchangeFailed)
		return err
	}

	err = ch.Publish(
		exchangeName,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        request,
		})

	if err != nil {
		log.Error(publishMessageFailed)
	}

	return nil
}

// keys 绑定多个路由
// key支持以下规则
// *（星号）：可以（只能）匹配一个单词
// #（井号）：可以匹配多个单词（或者零个）
func TopicReceive(exchangeName string, keys ...string) (<-chan amqp.Delivery, error) {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareExchangeFailed)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",
		false,
		true,
		true,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareQueueFailed)
		return nil, err
	}

	for _, key := range keys {
		err = ch.QueueBind(
			q.Name,
			key,
			exchangeName,
			false,
			nil)

		if err != nil {
			log.Error(bindQueueFailed)
			return nil, err
		}
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(registerConsumerFailed)
		return nil, err
	}

	return msgs, err
}

// RPC模式
func RPCClient(queueName string, request []byte) (reply []byte, err error) {
	ch, err := Conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",
		false,
		true,
		true,
		false,
		nil,
	)

	if err != nil {
		log.Error(declareQueueFailed)
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Error(registerConsumerFailed)
		return nil, err
	}

	corrId := uuid.NewV4().String()

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrId,
			ReplyTo:       q.Name,
			Body:          request,
		})

	if err != nil {
		log.Error(publishMessageFailed)
		return nil, err
	}

	select {
	case <-time.After(15 * time.Second):
		err = errors.New("请求超时")
	case msg := <-msgs:
		if corrId == msg.CorrelationId {
			reply = msg.Body
		}
	}

	return
}

func RPCServer(queueName string, f func([]byte) []byte) {
	ch, err := Conn.Channel()

	util.FatalOnError(err, openChannelFailed)

	q, err := ch.QueueDeclare(
		queueName,
		false,
		true,
		false,
		false,
		nil,
	)

	util.FatalOnError(err, declareQueueFailed)

	err = ch.Qos(
		1,
		0,
		false,
	)

	util.FatalOnError(err, setQoSFailed)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	util.FatalOnError(err, registerConsumerFailed)

	forever := make(chan bool)

	go func() {
		for d := range msgs {

			err = ch.Publish(
				"",
				d.ReplyTo,
				false,
				false,
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          f(d.Body),
				})

			if err != nil {
				// 推送失败可能导致client一直在等待，所以client需要做超时设置
				log.Error(publishMessageFailed)
				d.Ack(false)
				continue
			}

			d.Ack(false)
		}
	}()

	<-forever
}
