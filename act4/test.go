package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mvgmb/Middleware/act4/client"
	"github.com/mvgmb/Middleware/act4/server"
	"github.com/mvgmb/Middleware/act4/util"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Are you a client (1) or server (2): ")
	test, _ := reader.ReadString('\n')

	switch strings.TrimRight(test, "\n") {
	case "1":
		clientTest()
	case "2":
		serverTest()
	default:
		fmt.Println("Invalid input")
	}
}

func clientTest() {
	options := util.Options{
		Host:     "localhost",
		Port:     8080,
		Protocol: "tcp",
	}
	requestHandler, err := client.NewRequestHandler(options)
	if err != nil {
		log.Fatal(err)
	}

	err = requestHandler.Send([]byte("HelloWorld!"))
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := requestHandler.Receive()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}

func serverTest() {
	options := util.Options{
		Host:     "localhost",
		Port:     8080,
		Protocol: "tcp",
	}
	requestHandler, err := server.NewRequestHandler(options)
	if err != nil {
		log.Fatal(err)
	}

	err = requestHandler.Accept()
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := requestHandler.Receive()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))

	err = requestHandler.Send([]byte("Ok!"))
	if err != nil {
		log.Fatal(err)
	}

	err = requestHandler.Close()
	if err != nil {
		log.Fatal(err)
	}
}
