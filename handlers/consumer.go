package handlers

import (
	"github.com/bitly/go-nsq"
)

// NSQConsumer connects to NSQ and consume messages
type NSQConsumer struct {
	Topic    string
	Channel  string
	consumer *nsq.Consumer
}

// NewNSQConsumer create new NSQConsumer
func NewNSQConsumer(topic, channel string) (NSQConsumer, error) {
	var consumer NSQConsumer
	c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		return consumer, err
	}

	consumer = NSQConsumer{Topic: topic, Channel: channel, consumer: c}

	return consumer, nil
}

// ConnectToNSQDs connect to nsq deamons
func (c *NSQConsumer) ConnectToNSQDs(nsqds []string) error {
	return c.consumer.ConnectToNSQDs(nsqds)
}

// ConnectToLookupd connect to lookupd
func (c *NSQConsumer) ConnectToLookupd(lookupd string) error {
	return c.consumer.ConnectToNSQLookupd(lookupd)
}

// AddHandler add message handler
func (c *NSQConsumer) AddHandler(handler nsq.Handler) {
	c.consumer.AddHandler(handler)
}
