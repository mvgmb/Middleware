package main

import (
	"log"

	"github.com/mvgmb/Middleware/rpc/server"
	"github.com/mvgmb/Middleware/rpc/util"
)

var (
	invoker *server.Invoker
	options = util.Options{
		Host:     "localhost",
		Port:     0,
		Protocol: "tcp",
	}
)

func init() {
	var err error
	invoker, err = server.NewInvoker(&options)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	invoker.Invoke()
}
