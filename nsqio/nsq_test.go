package nsqio

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"testing"
	"time"
)

var producer *nsq.Producer

func TestSend(t *testing.T) {
	strIP1 := "127.0.0.1:4150"

	InitProducer(strIP1)

	running := true

	//读取控制台输入
	//reader := bufio.NewReader(os.Stdin)
	for running {
		//data, _, _ := reader.ReadLine()
		//command := string(data)
		//if command == "stop" {
		//	running = false
		//}

		for err := Publish("test", "333"); err != nil; err = Publish("test", "333") {
			InitProducer(strIP1)
		}
	}
	//关闭
	producer.Stop()

}

// 初始化生产者
func InitProducer(str string) {
	var err error
	fmt.Println("address: ", str)
	producer, err = nsq.NewProducer(str, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
}

//发布消息
func Publish(topic string, message string) error {
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		err = producer.Publish(topic, []byte(message)) // 发布消息
		return err
	}
	return fmt.Errorf("producer is nil:%v", err)
}

// 消费者
type ConsumerT struct{}

func TestReceive(t *testing.T) {
	InitConsumer("test", "test-channel", "127.0.0.1:4161")
	for {
		time.Sleep(time.Second * 10)
	}
}

//处理消息
func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

//初始化消费者
func InitConsumer(topic string, channel string, address string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second          //设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, 0)        //屏蔽系统日志
	c.AddHandler(&ConsumerT{}) // 添加消费者接口

	//建立NSQLookupd连接
	if err := c.ConnectToNSQLookupd(address); err != nil {
		panic(err)
	}

	//建立多个nsqd连接
	// if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
	//  panic(err)
	// }

	// 建立一个nsqd连接
	if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		panic(err)
	}
}
