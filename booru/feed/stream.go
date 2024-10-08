package feed

import (
	"applemango/boorutan/backend/booru"
	"applemango/boorutan/backend/db/logger"
	"applemango/boorutan/backend/handler"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

// https://github.com/gin-gonic/examples/blob/master/server-sent-event/main.go

type PostLogger struct {
	Post    booru.Post `json:"post"`
	Message string     `json:"message,omitempty"`
}

func RegisterBooruStreamer(app *gin.Engine) {
	stream := NewServer()
	go func() {
		for {
			b := handler.GetBooruFromString("danbooru")
			_ = generateFeedCore(b, 1, func(p booru.Post) {
				ps, err := p.ToString()
				if err != nil {
					return
				}
				stream.Message <- ps
				go func() {
					logger.Ctx.SendEvent(PostLogger{
						Post:    p,
						Message: "New Post",
					})
					/*err = SendWebhook(p)
					if err != nil {
						return
					}*/
				}()
			})
			time.Sleep(time.Second * 30)
		}
	}()

	app.GET("/post/stream", HeadersMiddleware(), stream.serveHTTP(), func(c *gin.Context) {
		v, ok := c.Get("clientChan")
		if !ok {
			return
		}
		clientChan, ok := v.(ClientChan)
		if !ok {
			return
		}
		c.Stream(func(w io.Writer) bool {
			if msg, ok := <-clientChan; ok {
				c.SSEvent("message", msg)
				return true
			}
			return false
		})
	})

}

type Event struct {
	Message       chan string
	NewClients    chan chan string
	ClosedClients chan chan string
	TotalClients  map[chan string]bool
}

type ClientChan chan string

func NewServer() (event *Event) {
	event = &Event{
		Message:       make(chan string),
		NewClients:    make(chan chan string),
		ClosedClients: make(chan chan string),
		TotalClients:  make(map[chan string]bool),
	}
	go event.listen()
	return
}

func (stream *Event) listen() {
	for {
		select {
		case client := <-stream.NewClients:
			stream.TotalClients[client] = true
		case client := <-stream.ClosedClients:
			delete(stream.TotalClients, client)
			close(client)
		case eventMsg := <-stream.Message:
			for clientMessageChan := range stream.TotalClients {
				clientMessageChan <- eventMsg
			}
		}
	}
}

func (stream *Event) serveHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientChan := make(ClientChan)
		stream.NewClients <- clientChan
		defer func() {
			stream.ClosedClients <- clientChan
		}()
		c.Set("clientChan", clientChan)
		c.Next()
	}
}

func HeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Next()
	}
}
