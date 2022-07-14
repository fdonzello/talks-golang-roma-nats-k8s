package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"go-roma.io/nats-demo/internal/k8s"
	"go-roma.io/nats-demo/internal/messaging"
)

func main() {
	nc := messaging.MustConnect()

	fmt.Println("waiting for messages")

	nc.Subscribe("deployments.k8s.demo.create", func(msg *nats.Msg) {
		fmt.Printf("deploy: %v\n", string(msg.Data))

		deploy, err := k8s.Deploy()
		if err != nil {
			log.Printf("error: %v\n", err)
			nc.Publish("notifications.k8s.demo.failed", []byte(err.Error()))
			return
		}

		nc.Publish("notifications.k8s.demo.created", []byte(deploy.Name))
	})

	for {

	}
}
