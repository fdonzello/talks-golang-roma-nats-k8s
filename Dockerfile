FROM golang:alpine AS builder

WORKDIR /demo

ARG main

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o app ${main}

FROM alpine

COPY --from=builder /demo/app .

CMD ["./app"]