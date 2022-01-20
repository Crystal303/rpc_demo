package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/docker/docker/pkg/pubsub"
)

func main() {
	p := pubsub.NewPublisher(100*time.Millisecond, 10)

	golang := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			return strings.HasPrefix(key, "golang:")
		}
		return false
	})

	docker := p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			return strings.HasPrefix(key, "docker:")
		}
		return false
	})

	go p.Publish("hi")
	go p.Publish("golang: https://golang.org")
	go p.Publish("docker: https://www.docker.com/")
	time.Sleep(time.Second)

	go func() {
		fmt.Println("golang topic:", <-golang)
	}()
	go func() {
		fmt.Println("docker topic:", <-docker)
	}()

	time.Sleep(time.Second)
}
