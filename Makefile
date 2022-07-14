build-deployer:
	docker build --build-arg=main=cmd/deployer/main.go -t nats-demo-deployer:1.0.0 .

build-publisher:
	docker build --build-arg=main=cmd/publisher/main.go -t nats-demo-publisher:1.0.0 .

run-publisher:
	docker run --net=host -v ~/.kube:/root/.kube  nats-demo-publisher:1.0.0

run-deployer:
	docker run --net=host -v ~/.kube:/root/.kube  nats-demo-deployer:1.0.0

run-nats-server:
	docker run --name=nats -d -p 8222:8222 --rm  -p 4222:4222 -ti nats:latest -m 8222

stop-and-remove-nats-server:
	docker rm -f nats