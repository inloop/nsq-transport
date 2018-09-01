package cmd

import (
	"net/mail"
	"net/url"
	"strconv"
	"strings"

	"github.com/inloop/nsq-transport/handlers"
	"github.com/urfave/cli"
	gomail "gopkg.in/gomail.v2"
)

// SMTPCommand ...
func SMTPCommand() cli.Command {
	return cli.Command{
		Name: "smtp",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "url",
				EnvVar: "SMTP_URL",
				Usage:  "Connection url for SMTP server",
			},
			cli.StringFlag{
				Name:   "sender",
				EnvVar: "SMTP_SENDER",
				Usage:  "Default sender",
			},
			cli.StringFlag{
				Name:   "from",
				EnvVar: "SMTP_FROM",
				Usage:  "Message sender (templated string)",
			},
			cli.StringFlag{
				Name:   "to",
				EnvVar: "SMTP_TO",
				Usage:  "Message receiver (templated string)",
			},
			cli.StringFlag{
				Name:   "subject",
				EnvVar: "SMTP_SUBJECT",
				Usage:  "Message subject (templated string)",
			},
			cli.StringFlag{
				Name:   "body-text",
				EnvVar: "SMTP_BODY_TEXT",
				Usage:  "Message text/plain body (templated string)",
			},
			cli.StringFlag{
				Name:   "body-html",
				EnvVar: "SMTP_BODY_HTML",
				Usage:  "Message text/html body (templated string)",
			},
		},
		Action: func(c *cli.Context) error {

			URL := c.String("url")
			sender := c.String("sender")

			dialer := getDialer(URL, sender)

			from := c.String("from")
			if from == "" {
				from = sender
			}
			to := c.String("to")
			subject := c.String("subject")
			bodyText := c.String("body-text")
			bodyHTML := c.String("body-html")

			return handlers.StartMessageConsumer(c.Parent(), func(m handlers.NSQTransporterMessage) error {

				address, err := mail.ParseAddress(m.GetString(from))
				if err != nil {
					return err
				}

				message := gomail.NewMessage()
				message.SetAddressHeader("From", address.Address, address.Name)
				message.SetHeader("To", strings.Split(m.GetString(to), ",")...)
				message.SetHeader("Subject", m.GetString(subject))
				message.SetBody("text/plain", m.GetString(bodyText))
				message.SetBody("text/html", m.GetString(bodyHTML))
				// m.Attach("/home/Alex/lolcat.jpg")
				if err := dialer.DialAndSend(message); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}

				return nil
			})

		},
	}
}

func getDialer(urlString, sender string) *gomail.Dialer {

	URL, _ := url.Parse(urlString)

	if URL == nil {
		panic("SMTP url not provided")
	}
	if URL.User == nil {
		panic("user credentials not provided")
	}

	host := strings.Split(URL.Host, ":")[0]
	username := URL.User.Username()
	password := ""
	if pass, exists := URL.User.Password(); exists == true {
		password = pass
	}
	port := 25
	if portValue, err := strconv.ParseInt(URL.Port(), 10, 32); err == nil {
		port = int(portValue)
	}

	return gomail.NewDialer(host, port, username, password)
}
