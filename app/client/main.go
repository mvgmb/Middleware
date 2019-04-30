package main

import (
	"fmt"
	"github.com/mvgmb/Middleware/rpc/client"
	"log"
	"time"
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

	t := time.Now()

	for i := 0; i < 1000; i++ {
		_, err := MoviePrice("Titanic")
		if err != nil {
			log.Println("Error:", err.Error())
		}
	}

	fmt.Println(time.Since(t))
}

// MoviePrice returns the requested movie price
func MoviePrice(movieName string) (int, error) {
	price, err := proxy.MoviePrice(movieName)
	if err != nil {
		return -1, err
	}
	return price, nil
}
