package handlers

import (
	"fmt"
	"strings"
	"sync"

	"github.com/urfave/cli"
)

// StartMessageConsumer start NSQConsumer with NSQTransport for given context
func StartMessageConsumer(c *cli.Context, handler NSQTransporterHandler) error {
	topic := c.String("topic")
	channel := c.String("channel")
	lookupd := c.String("lookupd")
	nsqds := c.String("nsqds")
	filter := c.String("filter")

	consumer, err := NewNSQConsumer(topic, channel)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	transport, err := NewNSQTransporter(handler, filter)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}
	consumer.AddHandler(transport)

	if lookupd != "" {
		err = consumer.ConnectToLookupd(lookupd)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
	}
	if nsqds != "" {
		fmt.Println(strings.Split(nsqds, ","))
		err = consumer.ConnectToNSQDs(strings.Split(nsqds, ","))
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

	return nil
}
