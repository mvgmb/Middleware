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
		Port:     1337,
		Protocol: "tcp",
	}
	requestor, err := client.NewRequestor(&options)
	if err != nil {
		log.Fatal(err)
	}

	req := util.NewMessage([]byte("BestPrice"), "Lookup", "OK", 200)

	res, err := requestor.Invoke(&options, &req)
	if err != nil {
		log.Fatal(err)
	}

	message := res.(pb.Message)
	fmt.Println(message.String())
}

func serverTest() {
	options := util.Options{
		Host:     "localhost",
		Port:     0,
		Protocol: "tcp",
	}

	invoker, err := server.NewInvoker(&options)
	if err != nil {
		log.Fatal(err)
	}

	invoker.Invoke()
}
