package mq

import (
	"errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"time"
	. "tools/util"
)

var (
	conn *amqp.Connection
)

const (
	connectFailed          = "Failed to connect to RabbitMQ"
	openChannelFailed      = "Failed to open a channel"
	declareQueueFailed     = "Failed to declare a queue"
	bindQueueFailed        = "Failed to bind a queue"
	declareExchangeFailed  = "Failed to declare an exchange"
	registerConsumerFailed = "Failed to register a consumer"
	publishMessageFailed   = "Failed to publish a message"
	setQoSFailed           = "Failed to set Qos"
)

func Connect(userName, password, address string) {
	var err error

	url := fmt.Sprintf("amqp://%v:%v@%v/", userName, password, address)

	conn, err = amqp.Dial(url)

	if err != nil {
		log.Fatalf("%s: %s; url: %s", connectFailed, err, url)
	}
}

// 普通模式
// 支持持久化
func Send(QueueName string, request []byte) error {
	ch, err := conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		QueueName,
		true,
		false,
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
			Headers:         amqp.Table{},    // 开启
			DeliveryMode:    amqp.Persistent, // 消息 durable
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            request,
		})

	if err != nil {
		log.Error(publishMessageFailed)
	}

	return err
}

func Receive(QueueName string) (<-chan amqp.Delivery, error) {
	ch, err := conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		QueueName,
		true,
		false,
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
	ch, err := conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
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
			Headers:         amqp.Table{},    // 开启
			DeliveryMode:    amqp.Persistent, // 消息 durable
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            request,
		})

	if err != nil {
		log.Error(publishMessageFailed)
		return err
	}

	return nil
}

func Worker(queueName string) (<-chan amqp.Delivery, error) {
	ch, err := conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
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

	return msgs, err
}

// fanout，广播模式
// 开启持久化
func Publish(exchangeName string, request []byte) error {
	ch, err := conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		false,
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
			Headers:         amqp.Table{},    // 开启
			DeliveryMode:    amqp.Persistent, // 消息 durable
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            request,
		})

	if err != nil {
		log.Error(publishMessageFailed)
	}

	return nil
}

func Subscribe(exchangeName string) (<-chan amqp.Delivery, error) {
	ch, err := conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		false,
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
		false,
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
	ch, err := conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return err
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
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

func RouteConsume(exchangeName, key string) (<-chan amqp.Delivery, error) {
	ch, err := conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
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
		false,
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

// RPC模式
func RPCClient(queueName string, request []byte) (reply []byte, err error) {
	ch, err := conn.Channel()

	if err != nil {
		log.Error(openChannelFailed)
		return nil, err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
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
	ch, err := conn.Channel()

	FatalOnError(err, openChannelFailed)

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	FatalOnError(err, declareQueueFailed)

	err = ch.Qos(
		1,
		0,
		false,
	)

	FatalOnError(err, setQoSFailed)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	FatalOnError(err, registerConsumerFailed)

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
