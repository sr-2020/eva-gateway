FROM golang:1.12 as builder

WORKDIR /go/src/github.com/sr2020/gateway

COPY ./src .

RUN GO111MODULE=on go get ./...

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/gateway .


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /go/bin/gateway .

EXPOSE 80

CMD ["./gateway"]
