ASYNC K8S DEPLOY WITH NATS
==========================

This repo contains the code developed during a talk at the Golang Roma Meetup.

It shows how to use Nats to deploy asynchronously on Kubernetes.

Requirements
------------

In order to successfully run the demo commands, you must have:

- a Kubernetes cluster running locally with the .kube folder in your home directory (Kubernetes shipped with Docker Desktop should be enough)
- Docker installed and running

Run the demo
------------

Run the Nats server:

    make run-nats-server

Build the deployer and publisher images:

    make build-deployer
    make build-publisher

Start the deployer that will consume all deployment messages:

    make run-deployer

In another terminal, Run the publisher to start a deploy:

    make run-publisher

Stop the nats-server
---------------

    make stop-and-remove-nats-server
