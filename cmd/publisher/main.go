package main

import "go-roma.io/nats-demo/internal/messaging"

func main() {
	nc := messaging.MustConnect()

	nc.Publish("deployments.k8s.demo.create", []byte("XXXXX"))

	nc.Flush()
}
