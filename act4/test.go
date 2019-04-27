package main

import (
	"bufio"
	"fmt"
	// "github.com/golang/protobuf/proto"
	"github.com/mvgmb/Middleware/act4/client"
	pb "github.com/mvgmb/Middleware/act4/proto"
	"github.com/mvgmb/Middleware/act4/server"
	"github.com/mvgmb/Middleware/act4/util"
	"log"
	"os"
	"strings"
)

func main() {
	// // Marshal test
	message := util.NewMovieMessage([]byte("matilda"), "movieName", "OK", 200)
	bytes, err := util.Marshal(&message)
	if err != nil {
		log.Fatal(err)
	}

	res := pb.MovieMessage{}

	err = util.Unmarshal(&bytes, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.String())

	//	RequestHandlers test
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
	requestor, err := client.NewRequestor(options)
	if err != nil {
		log.Fatal(err)
	}

	req := util.NewMovieMessage([]byte("matilda"), "movieName", "OK", 200)
	res := pb.MovieMessage{}

	err = requestor.Invoke(&req, &res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.String())
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

	res := pb.MovieMessage{}
	err = util.Unmarshal(&bytes, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.String())

	response := util.NewMovieMessage([]byte("12,99"), "movieName", "OK", 200)

	data, err := util.Marshal(&response)
	if err != nil {
		log.Fatal(err)
	}

	err = requestHandler.Send(data)
	if err != nil {
		log.Fatal(err)
	}

	err = requestHandler.Close()
	if err != nil {
		log.Fatal(err)
	}
}
