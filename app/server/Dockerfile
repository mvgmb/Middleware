FROM golang:latest

WORKDIR /go/src/github.com/mvgmb/Middleware
COPY ./app/server .
COPY ./rpc ./rpc

RUN go get -v ./... && \
    go build -v ./...

CMD ["/go/bin/Middleware"]
