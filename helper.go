package nsqhelper

import (
	"encoding/json"
	"errors"

	nsq "github.com/nsqio/go-nsq"
)

// NSQHandler ...
type NSQHandler struct {
	Addr string
	Cfg  *nsq.Config
}

// Publish 发布消息
func (n *NSQHandler) Publish(topic string, data interface{}) error {
	if topic == "" {
		return errors.New("nsq topic cannot empty for publish")
	}

	p, err := nsq.NewProducer(n.Addr, n.Cfg)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.Publish(topic, jsonData)
}

// Consumer nsq 消费者
func (n *NSQHandler) Consumer(topic string, channel string, handle func(m *nsq.Message) error) error {
	if topic == "" || channel == "" {
		return errors.New("topic or channel cannot empty for nsq consumer")
	}

	c, err := nsq.NewConsumer(topic, channel, n.Cfg)
	if err != nil {
		return err
	}

	c.AddHandler(nsq.HandlerFunc(handle))
	if err = c.ConnectToNSQD(n.Addr); err != nil {
		return err
	}

	return nil
}
