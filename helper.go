package nsqhelper

import (
	"encoding/json"
	"errors"

	nsq "github.com/nsqio/go-nsq"
)

// NSQHandler ...
type NSQHandler struct {
	Producer *nsq.Producer
}

// Publish 发布消息
func (n *NSQHandler) Publish(topic string, data interface{}) error {
	if topic == "" {
		return errors.New("nsq topic cannot empty for publish")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return n.Producer.Publish(topic, jsonData)
}
