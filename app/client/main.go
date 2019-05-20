package main

import (
	"fmt"
	"log"

	"github.com/mvgmb/Middleware/rpc/client"
)

var proxy *client.Proxy

func init() {
	var err error
	proxy, err = client.NewProxy()
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	var price int
	var err error

	for i := 0; i < 1; i++ {
		price, err = proxy.MoviePrice("Titanic")
		if err != nil {
			log.Println("Error:", err.Error())
		}
	}

	fmt.Println(price)
}
