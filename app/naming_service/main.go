package main

import (
	"log"

	"github.com/mvgmb/Middleware/rpc/naming"
	"github.com/mvgmb/Middleware/rpc/util"
)

var (
	invoker *naming.Invoker
	options = util.Options{
		Host:     "localhost",
		Port:     1337,
		Protocol: "tcp",
	}
)

func init() {
	var err error
	invoker, err = naming.NewInvoker(&options)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	invoker.Invoke()
}
