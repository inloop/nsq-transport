package handlers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"strings"

	nsq "github.com/bitly/go-nsq"
)

// NSQTransporterMessage ...
type NSQTransporterMessage struct {
	Message        string
	DecodedMessage *interface{}
}

// GetString get string templated string
func (m *NSQTransporterMessage) GetString(tmpl string) string {
	if m.DecodedMessage == nil {
		return tmpl
	}

	var tpl bytes.Buffer
	_tmpl := template.Must(template.New(tmpl).Delims("[[", "]]").Parse(tmpl))
	_tmpl.Execute(&tpl, m.DecodedMessage)
	str := tpl.String()
	return strings.Replace(str, "\\n", "\n", -1)
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
