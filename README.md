# nsq-transport

NSQ consumer for transporting messages to various targets.

Currently supported:

- [x] slack
- [x] smtp
- [ ] nsq

## Usage

Each application setup NSQ consumer for specific topic with given channel. You can also filter received messages.

You can run application directly:

```
# install binary (macos/linux)
GOBIN=/usr/local/bin/ go install github.com/inloop/nsq-transport

# run with SMTP transport
nsq-transport -t some-topic -c my-smtp-transporter --lookup hostname:4161 smtp --sender john.doe@example.com --body-text "Hello, new message was just created" --body-html "test [[.entity]]" --to blah@example.com --subject "transport test" --url smtp://user:pass@hostname
```

Or using docker:

```
docker run --rm inloopx/nsq-transport -t some-topic -c my-smtp-transporter --lookup hostname:4161 smtp --sender john.doe@example.com --body-text "Hello, new message was just created" --body-html "test [[.entity]]" --to blah@example.com --subject "transport test" --url smtp://user:pass@hostname
```

Each attribute can be provided by environment variable.

### NSQ Consumer config

```
COMMANDS:
     slack
     smtp
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --topic value, -t value          NSQ topic name to consume [$NSQ_TOPIC]
   --channel value, -c value        Consumer channel name (default: "nsq-transport") [$NSQ_CHANNEL]
   --lookupd value, -l value        Address for nsq lookup [$NSQ_LOOKUPD]
   --nsqds value                    Comma separated list of addresses for nsqd [$NSQD_ADDRESSES]
   --filter {{eq .type "CREATED"}}  Golang template executed on parsed message(if JSON). If returns "true", message will beprocessed (ex. {{eq .type "CREATED"}}) [$NSQ_FILTER]
```

### Slack

```
USAGE:
   nsq-transport slack [command options] [arguments...]

OPTIONS:
   --channel value   Slack channel to send messages to (templated string) [$SLACK_CHANNEL]
   --webhook value   Slack webhook url [$SLACK_URL]
   --text value      Slack message text (templated string) [$SLACK_TEXT]
   --username value  Username for sender of the message (templated string) [$SLACK_USERNAME]
```

### SMTP

```
USAGE:
   nsq-transport smtp [command options] [arguments...]

OPTIONS:
   --url value        Connection url for SMTP server [$SMTP_URL]
   --sender value     Default sender [$SMTP_SENDER]
   --from value       Message sender (templated string) [$SMTP_FROM]
   --to value         Message receiver (templated string) [$SMTP_TO]
   --subject value    Message subject (templated string) [$SMTP_SUBJECT]
   --body-text value  Message text/plain body (templated string) [$SMTP_BODY_TEXT]
   --body-html value  Message text/html body (templated string) [$SMTP_BODY_HTML]
```

# Templating

Each attribute marked as "templated string" is used as golang template. If event body contains JSON parsed data, you can access these information using golang templates. Be aware that `"[[","]]"` are used as delimiters.

Example:

```
Entity message: {hello:"world"}

Template: "Greetings from template [[.hello]]"
```
