package util

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/mvgmb/Middleware/act4/proto"
)

type Marshaller struct {
}

// NewMarshaller constructs a new Mashaller
func NewMarshaller() (*Marshaller, error) {
	return &Marshaller{}, nil
}

// Marshal serializes the message into bytes
func Marshal(message proto.Message) ([]byte, error) {
	return proto.Marshal(message)
}

// Unmarshal retrieves the serialized message
func Unmarshal(message []byte) (*pb.MovieMessage, error) {
	result := &pb.MovieMessage{}
	err := proto.Unmarshal(message, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
proto.Message is an interface, Marshal/Unmarshal will receive &pb.MovieMessage{ ... }
https://developers.google.com/protocol-buffers/docs/gotutorial
*/
