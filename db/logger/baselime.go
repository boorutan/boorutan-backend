package logger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var Ctx = NewClient(Option{
	ApiKey:  "b6ba54a209eaef8d560e61ce00455f8c50f5803d",
	Service: "booru",
})

type Logger struct {
	ApiKey  string
	Service string
}

type Option struct {
	ApiKey  string
	Service string
}

func NewClient(option Option) Logger {
	return Logger{ApiKey: option.ApiKey, Service: option.Service}
}

func (l Logger) Request(body any) (*http.Response, error) {
	url := fmt.Sprintf("https://events.baselime.io/v1/logs")
	marshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(string(marshal))

	req, _ := http.NewRequest("POST", url, r)
	req.Header.Add("x-api-key", l.ApiKey)
	req.Header.Add("x-service", l.Service)
	client := new(http.Client)
	resp, err := client.Do(req)
	return resp, err
}

func (l Logger) pushLog(data any) {
	go func() {
		_, _ = l.Request(data)
	}()
}

func (l Logger) SendEvent(data any) {
	l.pushLog(data)
}
