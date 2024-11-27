package nsq

import (
	"github.com/nsqio/go-nsq"
)

func NewConsumer(topic, channel, lookupdAddress string, handler nsq.Handler) (*nsq.Consumer, error) {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, err
	}

	consumer.AddHandler(handler)

	err = consumer.ConnectToNSQLookupd(lookupdAddress)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
