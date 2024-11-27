package nsq

import (
	"encoding/json"
	"log"

	"github.com/nsqio/go-nsq"
)

type Publisher struct {
	producer *nsq.Producer
}

type Message struct {
	Topic   string      `json:"topic"`
	Payload interface{} `json:"payload"`
}

func NewPublisher(nsqdAddress string) (*Publisher, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(nsqdAddress, config)
	if err != nil {
		return nil, err
	}

	return &Publisher{producer: producer}, nil
}

func (p *Publisher) Publish(topic string, message interface{}) error {
	nsqMessage := Message{
		Topic:   topic,
		Payload: message,
	}

	messageBytes, err := json.Marshal(nsqMessage)
	if err != nil {
		return err
	}

	err = p.producer.Publish(topic, messageBytes)
	if err != nil {
		log.Println("Error publishing message:", err)
		return err
	}
	return nil
}

func (p *Publisher) Stop() {
	p.producer.Stop()
}
