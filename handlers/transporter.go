package handlers

import (
	"bytes"
	"encoding/json"
	"html/template"

	nsq "github.com/bitly/go-nsq"
)

// NSQTransporterMessage ...
type NSQTransporterMessage struct {
	Message        string
	DecodedMessage *interface{}
}

// GetText get string templated string
func (m *NSQTransporterMessage) GetString(tmpl string) string {
	var tpl bytes.Buffer
	_tmpl := template.Must(template.New(tmpl).Parse(tmpl))
	_tmpl.Delims("[[", "]]")
	_tmpl.Execute(&tpl, m.DecodedMessage)
	return tpl.String()
}

// NSQTransporterHandler ...
type NSQTransporterHandler func(m NSQTransporterMessage) error

// NSQTransporter ...
type NSQTransporter struct {
	handler NSQTransporterHandler
	filter  string
}

// NewNSQTransporter ...
func NewNSQTransporter(handler NSQTransporterHandler, filter string) (NSQTransporter, error) {
	return NSQTransporter{handler, filter}, nil
}

// HandleMessage handle NSQ message
func (t NSQTransporter) HandleMessage(m *nsq.Message) error {

	var body interface{}
	json.Unmarshal(m.Body, &body)

	message := NSQTransporterMessage{Message: string(m.Body), DecodedMessage: &body}

	defer m.Finish()

	if t.filter != "" && message.GetString(t.filter) != "true" {
		return nil
	}

	return t.handler(message)

}
