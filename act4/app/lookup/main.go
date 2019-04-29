package main

import (
	"log"
	"math"
	"net"

	"github.com/golang/protobuf/proto"
	pb "github.com/mvgmb/Middleware/act4/proto"
	"github.com/mvgmb/Middleware/act4/util"
)

func main() {
	marshaller, err := util.NewMarshaller()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening at localhost:1337")

	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Print(err)
	}

	for {
		netConn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		buffer := make([]byte, math.MaxInt16)

		n, err := netConn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		buffer = buffer[:n]

		req := pb.Message{}
		marshaller.Unmarshal(&buffer, &req)

		var message proto.Message

		switch req.TypeName {
		case "Lookup":
			result, err := util.Lookup(string(req.MessageData))
			if err != nil {
				message = util.ErrNotFound
				break
			}
			message = util.NewMessage([]byte(result.String()), "AOR", "OK", 200)
		case "Bind":
			util.Bind(util.StringToAOR(string(req.MessageData)))
			message = util.NewMessage([]byte(""), "", "OK", 200)
		default:
			message = util.ErrBadRequest
		}

		bytes, err := marshaller.Marshal(&message)
		if err != nil {
			log.Fatal(err)
		}

		_, err = netConn.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}

	}
}
