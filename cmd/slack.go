package cmd

import (
	"fmt"

	slack "github.com/ashwanthkumar/slack-go-webhook"
	"github.com/inloop/nsq-transport/handlers"
	"github.com/urfave/cli"
)

// SlackCommand ...
func SlackCommand() cli.Command {
	return cli.Command{
		Name: "slack",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "channel",
				EnvVar: "SLACK_CHANNEL",
				Usage:  "Slack channel to send messages to (templated string)",
			},
			cli.StringFlag{
				Name:   "webhook",
				EnvVar: "SLACK_URL",
				Usage:  "Slack webhook url",
			},
			cli.StringFlag{
				Name:   "text",
				EnvVar: "SLACK_TEXT",
				Usage:  "Slack message text (templated string)",
			},
			cli.StringFlag{
				Name:   "username",
				EnvVar: "SLACK_USERNAME",
				Usage:  "Username for sender of the message (templated string)",
			},
		},
		Action: func(c *cli.Context) error {

			channel := c.String("channel")
			text := c.String("text")
			username := c.String("username")
			webhookURL := c.String("webhook")

			return handlers.StartMessageConsumer(c.Parent(), func(m handlers.NSQTransporterMessage) error {
				payload := slack.Payload{
					Text:     m.GetString(text),
					Username: m.GetString(username),
					Channel:  m.GetString(channel),
				}
				errs := slack.Send(webhookURL, "", payload)
				if len(errs) > 0 {
					return fmt.Errorf("Error sending message to slack %s", errs[0].Error())
				}

				return nil
			})

		},
	}
}
