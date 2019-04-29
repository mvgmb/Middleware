package main

import (
	"github.com/mvgmb/Middleware/rpc/server"
	"github.com/mvgmb/Middleware/rpc/util"
	"log"
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

// MoviePrice returns the requested movie price
func MoviePrice(movieName string) (int, error) {
	price, err := invoker.Proxy.Movie.Price(movieName)
	if err != nil {
		return -1, err
	}
	return price, nil
}
