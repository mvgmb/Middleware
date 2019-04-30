FROM golang:latest

WORKDIR /go/src/github.com/mvgmb/Middleware
COPY ./app/naming_service .
COPY ./rpc ./rpc

RUN go get -v ./... && \
    go build -v ./...

CMD ["/go/bin/Middleware"]
