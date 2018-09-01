package main

import (
	"os"

	// "github.com/inloop/go-transport-queue/model"
	// "github.com/inloop/go-transport-queue/transports"

	"github.com/inloop/nsq-transport/cmd"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "nsq-transport"
	app.Usage = "..."
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "topic,t",
			EnvVar: "NSQ_TOPIC",
			Usage:  "NSQ topic name to consume",
		},
		cli.StringFlag{
			Name:   "channel,c",
			EnvVar: "NSQ_CHANNEL",
			Usage:  "Consumer channel name",
			Value:  "nsq-transport",
		},
		cli.StringFlag{
			Name:   "lookupd,l",
			EnvVar: "NSQ_LOOKUPD",
			Usage:  "Address for nsq lookup",
		},
		cli.StringFlag{
			Name:   "nsqds",
			EnvVar: "NSQD_ADDRESSES",
			Usage:  "Comma separated list of addresses for nsqd",
		},
		cli.StringFlag{
			Name:   "filter",
			EnvVar: "NSQ_FILTER",
			Usage:  "Golang template executed on parsed message(if JSON). If returns \"true\", message will be processed (ex. `{{eq .type \"CREATED\"}}`)",
		},
	}

	app.Commands = []cli.Command{
		cmd.SlackCommand(),
		cmd.SMTPCommand(),
	}

	app.Run(os.Args)
}
