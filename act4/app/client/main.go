package main

import (
	"fmt"
	"github.com/mvgmb/Middleware/act4/client"
	"log"
)

var proxy *client.Proxy

func main() {
	var err error
	proxy, err = client.NewProxy()
	if err != nil {
		log.Fatal(err)
	}

	price, err := MoviePrice("Titanic")
	if err != nil {
		log.Println("Error:", err.Error())
	}
	fmt.Println(price)
}

// MoviePrice returns the requested movie price
func MoviePrice(movieName string) (int, error) {
	price, err := proxy.MoviePrice(movieName)
	if err != nil {
		return -1, err
	}
	return price, nil
}
